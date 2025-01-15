package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/streadway/amqp"
	"golang.org/x/net/context"
	"golang.org/x/time/rate"
	"io/ioutil"
	"log"
	"my_e_commerce/config"
	model2 "my_e_commerce/data/req"
	service "my_e_commerce/data/response"
	service2 "my_e_commerce/service"
	"my_e_commerce/utils"
	"strconv"
)

type SeckillHandler struct {
	seckillRecordService service2.SeckillRecordService
	orderService         service2.OrderService
	stockService         service2.StockService
	goodsService         service2.GoodsService
}

func NewSeckillHanlder(seckRecordService service2.SeckillRecordService,
	orderService service2.OrderService, stockService service2.StockService,
	goodsService service2.GoodsService) *SeckillHandler {
	return &SeckillHandler{
		seckillRecordService: seckRecordService,
		orderService:         orderService,
		stockService:         stockService,
		goodsService:         goodsService,
	}
}

// 秒杀接口
/*
	@Input seckillRecordReq 秒杀请求数据
	@Output seckillNum 秒杀号
*/

func (h *SeckillHandler) CreateSeckill(c *gin.Context) {
	// TODO 对整个controller过程加事务
	// TODO 临时写死，有前端了以后用前端的header获取
	//userID := 1
	// 商品id 购买数量
	// 先创建订单再创建秒杀订单
	var seckillReq model2.SeckillReq
	err := c.BindJSON(&seckillReq)
	if err != nil {
		log.Printf("bind json err: %+v\n", err)
		c.JSON(service.ERR_JSON_BIND, service.GetResponse(service.ERR_JSON_BIND, service.GetErrMsg(service.ERR_JSON_BIND), nil))
		return
	}
	// 校验参数
	if seckillReq.GoodsAmount <= 0 {
		c.JSON(service.ERR_INPUT_INVALID, service.GetResponse(service.ERR_INPUT_INVALID, service.GetErrMsg(service.ERR_INPUT_INVALID), nil))
	}
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")

	utils.SendMessageOnce(conn, "seckill", seckillReq)
	// TODO 创建秒杀单
	c.JSON(200, service.GetResponse(service.SUCCESS, service.GetErrMsg(service.SUCCESS), nil))
	return

}

// 接收函数（带限流）
func (h *SeckillHandler) ReceiveMessage(conn *amqp.Connection, queueName string, limit rate.Limit) {
	conn, _ = amqp.Dial("amqp://guest:guest@localhost:5672/")
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("mq err1 : %+v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no - wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("mq err2 : %+v", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto - ack
		false,  // exclusive
		false,  // no - local
		false,  // no - wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("mq err3 : %+v", err)
	}

	limiter := rate.NewLimiter(limit, 1)
	log.Printf("begin listen mq")
	for d := range msgs {
		if limiter.Allow() {
			var data model2.SeckillReq
			err = json.Unmarshal(d.Body, &data)
			if err != nil {
				log.Fatalf("Failed to unmarshal data : %+v", err)
			}
			log.Printf("Received: %+v", data)
			h.CreateSeckillByRabbitmq(data)
		} else {
			log.Println("Rate limit exceeded, skipping message")
		}
	}
}
func (h *SeckillHandler) CreateSeckillByRabbitmq(seckillReq model2.SeckillReq) {
	// TODO 先去查redis库存数据
	// 读取.lua 文件的内容
	userID := 1
	scriptBytes, err := ioutil.ReadFile("./lua/redis_substock.lua")
	if err != nil {
		fmt.Println("Error reading Lua script file:", err)
		return
	}
	script := string(scriptBytes)
	rdb, err := config.GetRedisConnection()
	if err != nil {
		fmt.Println("Error GetRedisConnection:", err)
		return
	}
	ctx := context.Background()

	result, err := rdb.Eval(ctx, script, []string{strconv.Itoa(int(seckillReq.GoodsID))}, seckillReq.GoodsAmount).Result()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("redis subStock result: %+v", result)
	// 数据库扣库存
	ok, err := h.stockService.SubStock(seckillReq.GoodsID, seckillReq.GoodsAmount)
	if err != nil {
		log.Fatalf("this is err , %+v\n", err)
		return
	}
	if !ok {
		log.Fatalf("%s\n", service.GetErrMsg(service.ERR_DESC_STOCK_FAILED))
	}
	// 创建订单
	var orderReq model2.OrderReq
	utils.CopyStruct(seckillReq, orderReq)
	if err != nil {
		fmt.Printf("copyStruct err : %+v\n", err)
	}
	// 查询商品信息
	goods, err := h.goodsService.GetGoods(seckillReq.GoodsID)
	if err != nil {
		log.Fatalf("%s\n", service.GetErrMsg(service.ERR_GET_GOODS_FAILED))
	}
	if goods == nil || len(goods) == 0 {
		log.Fatalf("%s\n", service.GetErrMsg(service.ERR_GET_GOODS_EMPTY_FAILED))
	}
	good := goods[0]
	fmt.Printf("good : %+v\n", good)
	orderReq.GoodsID = seckillReq.GoodsID
	orderReq.GoodsNum = good.GoodsNum
	orderReq.Price = good.Price.Mul(decimal.NewFromUint64(uint64(seckillReq.GoodsAmount)))
	orderReq.Status = 0
	orderReq.GoodsAmount = &seckillReq.GoodsAmount
	orderReq.Buyer = uint32(userID)
	orderReq.Seller = good.Seller
	fmt.Printf("orderReq : %+v\n", orderReq)
	_, err = h.orderService.CreateOrder(orderReq)
	if err != nil {
		log.Fatalln("%s\n", service.GetErrMsg(service.ERR_CREATE_ORDER_FAILED))
		return
	}
	log.Printf("create a seckillOrder SUCCESS!")
}

func MultiplyFloatsAsDecimals(price float32, amount uint32) float64 {
	// 将 float64 转换为 Decimal
	decA := decimal.NewFromFloat(float64(price))

	decB := decimal.NewFromFloat(float64(amount))
	// 进行乘法运算
	result := decA.Mul(decB)

	return result.InexactFloat64()
}
