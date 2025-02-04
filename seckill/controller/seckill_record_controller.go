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
	"strings"
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
	dsn := config.GetRabbitmqDSN()
	conn, _ := amqp.Dial(dsn)

	utils.SendMessageOnce(conn, "seckill", seckillReq)
	// TODO 创建秒杀单
	c.JSON(200, service.GetResponse(service.SUCCESS, service.GetErrMsg(service.SUCCESS), nil))
	return

}

// 接收函数（带限流）
func (h *SeckillHandler) ReceiveMessage(conn *amqp.Connection, queueName string, limit rate.Limit) {
	dsn := config.GetRabbitmqDSN()
	conn, _ = amqp.Dial(dsn)
	log.Printf("receiver channel: %s", dsn)
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("mq err1 : %+v", err)
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
		log.Printf("mq err2 : %+v", err)
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
		log.Printf("mq err3 : %+v", err)
	}

	limiter := rate.NewLimiter(limit, 1)
	log.Printf("begin listen mq")
	for d := range msgs {
		if limiter.Allow() {
			var data model2.SeckillReq
			err = json.Unmarshal(d.Body, &data)
			if err != nil {
				if r := recover(); r != nil {
					log.Printf("Panic occurred while processing message: %v", r)
				}
				log.Printf("Failed to unmarshal data : %+v", err)
			}
			log.Printf("Received: %+v", data)
			defer func() {
				d.Ack(false)
			}()
			h.CreateSeckillByRabbitmq(data)
		} else {
			log.Println("Rate limit exceeded, skipping message")
		}
	}
}

func (h *SeckillHandler) CreateSeckillByRabbitmq(seckillReq model2.SeckillReq) {
	// TODO 先去查redis库存数据
	// 读取.lua 文件的内容
	defer func() {
		if err := recover(); err != nil {
			log.Printf("ERROR recover: %s", err)
		}
	}()
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
	log.Println(seckillReq.GoodsID)
	result, err := rdb.Eval(ctx, script, []string{strconv.Itoa(int(seckillReq.GoodsID))}, seckillReq.GoodsAmount).Result()
	res, ok := result.(int64)
	if !ok {
		log.Printf("秒杀接口断言转换失败！")
		return
	}
	if res == 0 {
		log.Printf("扣库存失败， message:%+v", seckillReq)
		return
	}
	if err != nil {
		fmt.Println("Error lua:", err)
		return
	}
	db := config.GetDB()
	tx := db.Begin()
	if tx.Error != nil {
		log.Printf("Failed to begin transaction: %+v", tx.Error)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			// 回滚数据库
			tx.Rollback()
			log.Printf("2恢复redis")
			// 恢复redis库存
			redisAdd(seckillReq)
			log.Printf("err from db: %+v", r)
		} else {
			log.Printf("事务提交")
			if err := tx.Commit().Error; err != nil {
				if !strings.Contains(err.Error(), "transaction has already been committed or rolled back") {
					// 回滚数据库
					tx.Rollback()
					log.Printf("1恢复redis")
					// 恢复redis库存
					redisAdd(seckillReq)
				}
				log.Printf("Failed to commit transaction: %+v", err)
				return
			}
		}
	}()

	// 数据库扣库存
	ok, err = h.stockService.SubStock(tx, seckillReq.GoodsID, seckillReq.GoodsAmount)
	if err != nil {
		log.Printf("this is err , %+v\n", err)
		// 回滚数据库
		tx.Rollback()
		log.Printf("3恢复redis")
		// 恢复redis库存
		redisAdd(seckillReq)
		return
	}
	if !ok {
		log.Printf("%s\n", service.GetErrMsg(service.ERR_DESC_STOCK_FAILED))
		// 回滚数据库
		tx.Rollback()
		log.Printf("4恢复redis")
		// 恢复redis库存
		redisAdd(seckillReq)
		return
	}
	log.Printf("mysql subStock")

	// 创建订单
	var orderReq model2.OrderReq
	utils.CopyStruct(seckillReq, orderReq)
	if err != nil {
		fmt.Printf("copyStruct err : %+v\n", err)
	}
	// 查询商品信息
	goods, err := h.goodsService.GetGoods(seckillReq.GoodsID)
	if err != nil {
		log.Printf("%s\n", service.GetErrMsg(service.ERR_GET_GOODS_FAILED))
	}
	if goods == nil || len(goods) == 0 {
		log.Printf("%s\n", service.GetErrMsg(service.ERR_GET_GOODS_EMPTY_FAILED))
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
	_, err = h.orderService.CreateOrder(tx, orderReq)

	if err != nil {
		log.Fatalln("%s\n", service.GetErrMsg(service.ERR_CREATE_ORDER_FAILED))
		// 回滚数据库
		tx.Rollback()
		log.Printf("5恢复redis")
		// 恢复redis库存
		redisAdd(seckillReq)
		return
	}

	log.Printf("create a seckillOrder SUCCESS!")
}

func redisAdd(seckillReq model2.SeckillReq) {
	scriptBytes, err := ioutil.ReadFile("./lua/redis_substock_add.lua")
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
	rdb.Eval(ctx, script, []string{strconv.Itoa(int(seckillReq.GoodsID))}, seckillReq.GoodsAmount).Result()
}

func MultiplyFloatsAsDecimals(price float32, amount uint32) float64 {
	// 将 float64 转换为 Decimal
	decA := decimal.NewFromFloat(float64(price))

	decB := decimal.NewFromFloat(float64(amount))
	// 进行乘法运算
	result := decA.Mul(decB)

	return result.InexactFloat64()
}
