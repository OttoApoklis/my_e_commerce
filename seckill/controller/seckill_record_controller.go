package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/shopspring/decimal"
	"github.com/streadway/amqp"
	"golang.org/x/net/context"
	"golang.org/x/time/rate"
	"io/ioutil"
	"log"
	"my_e_commerce/config"
	"my_e_commerce/data/dal/model"
	model2 "my_e_commerce/data/req"
	model3 "my_e_commerce/data/resp"
	service "my_e_commerce/data/response"
	"my_e_commerce/enum"
	redis_test "my_e_commerce/redis_init"
	service2 "my_e_commerce/service"
	"my_e_commerce/utils"
	"strconv"
	"strings"
	"time"
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
	secNum := strconv.FormatUint(uint64(utils.GetSnowCode()), 10)
	seckillReq.SecNum = &secNum
	utils.SendMessageOnce("seckill", seckillReq)
	// TODO 创建秒杀单
	c.JSON(200, service.GetResponse(service.SUCCESS, service.GetErrMsg(service.SUCCESS), secNum))
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
		true,   // auto - ack
		false,  // exclusive
		false,  // no - local
		false,  // no - wait
		nil,    // args
	)
	if err != nil {
		log.Printf("mq err3 : %+v", err)
	}

	limiter := rate.NewLimiter(limit, 10)
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
			h.CreateSeckillByRabbitmq(data)
			//if err = d.Ack(false); err != nil {
			//	log.Printf("err caused by ack: %v.", err)
			//}
		} else {
			log.Println("Rate limit exceeded, skipping message")
		}
	}
}

func (h *SeckillHandler) CreateSeckillByRabbitmq(seckillReq model2.SeckillReq) {
	// 读取.lua 文件的内容
	defer func() {
		if err := recover(); err != nil {
			log.Printf("ERROR recover: %s", err)
		}
	}()
	if seckillReq.SecNum == nil {
		log.Println("secNum is nil!")
		return
	}
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
	// 先去查redis库存数据
	result, err := rdb.Eval(ctx, script, []string{strconv.Itoa(int(seckillReq.GoodsID))}, seckillReq.GoodsAmount).Result()
	res, ok := result.(int64)
	if !ok {
		log.Printf("秒杀接口断言转换失败！")
		return
	}
	if res == 0 {
		log.Printf("redis扣库存失败， message:%+v", seckillReq)
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
			log.Printf("2恢复redis并回滚mysql")
			// 恢复redis并回滚mysql库存
			redisAdd(seckillReq)
			log.Printf("err from db: %+v", r)
		} else {

			if err := tx.Commit().Error; err != nil {
				if !strings.Contains(err.Error(), "transaction has already been committed or rolled back") {
					// 回滚数据库
					tx.Rollback()
					log.Printf("1恢复redis并回滚mysql")
					// 恢复redis并回滚mysql库存
					redisAdd(seckillReq)
				}
				log.Printf("Failed to commit transaction: %+v", err)
				return
			} else {
				log.Printf("事务提交")
			}
		}
	}()

	// 数据库扣库存
	ok, err = h.stockService.SubStock(tx, seckillReq.GoodsID, seckillReq.GoodsAmount)
	if err != nil {
		log.Printf("this is err , %+v\n", err)
		// 回滚数据库
		tx.Rollback()
		log.Printf("3恢复redis并回滚mysql")
		// 恢复redis并回滚mysql库存
		redisAdd(seckillReq)
		return
	}
	if !ok {
		log.Printf("%s\n", service.GetErrMsg(service.ERR_DESC_STOCK_FAILED))
		// 回滚数据库
		tx.Rollback()
		log.Printf("4恢复redis并回滚mysql")
		// 恢复redis并回滚mysql库存
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
	orderReq.Status = enum.SECKILL_ORDER_CREATED
	orderReq.GoodsAmount = &seckillReq.GoodsAmount
	orderReq.Buyer = uint32(userID)
	orderReq.Seller = good.Seller
	orderNum, err := utils.GetUUID()
	orderReq.OrderNum = &orderNum
	if err != nil {
		log.Printf("getUUID err caused by %v", err)
		return
	}
	fmt.Printf("orderReq : %+v\n", orderReq)
	_, err = h.orderService.CreateOrder(tx, orderReq)

	if err != nil {
		log.Printf("%s\n", service.GetErrMsg(service.ERR_CREATE_ORDER_FAILED))
		// 回滚数据库
		tx.Rollback()
		log.Printf("5恢复redis并回滚mysql")
		// 恢复redis并回滚mysql库存
		redisAdd(seckillReq)
		return
	}

	var seckillRecord model2.SeckillRecordReq
	utils.CopyStruct(&orderReq, &seckillRecord)
	seckillRecord.UserID = orderReq.Buyer
	seckillRecord.SecNum = seckillReq.SecNum
	seckillRecord.OrderNum = &orderNum
	seckillRecord.Status = enum.SECKILL_ORDER_CREATED
	log.Printf("seckillRecord:%v.", &seckillRecord)
	err = h.seckillRecordService.CreateSeckillRecord(&seckillRecord)

	if err != nil {
		log.Printf("%s\n", service.GetErrMsg(service.ERR_CREATE_ORDER_FAILED))
		// 回滚数据库
		tx.Rollback()
		log.Printf("6恢复redis并回滚mysql")
		// 恢复redis并回滚mysql库存
		redisAdd(seckillReq)
		return
	}
	duration := 15 * time.Minute
	// 将秒杀号和订单号添加到redis有序集合中
	members1 := []*redis.Z{}
	member1 := redis.Z{Score: float64(time.Now().Add(duration).Unix()), Member: *seckillRecord.SecNum}
	members1 = append(members1, &member1)
	ok, err = utils.AddToSortedSet(ctx, rdb, redis_test.ZSETKEY_SECKILL_RECORD, members1)
	if err != nil {
		log.Printf("AddToSortedSet err caused by %v.", err)
	}

	members2 := []*redis.Z{}
	member2 := redis.Z{Score: float64(time.Now().Add(duration).Unix()), Member: *seckillRecord.OrderNum}
	members2 = append(members2, &member2)
	ok, err = utils.AddToSortedSet(ctx, rdb, redis_test.ZSETKEY_ORDER, members2)
	if err != nil {
		log.Printf("AddToSortedSet err caused by %v.", err)
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

func (h *SeckillHandler) Buy(c *gin.Context) {
	/*TODO 付款操作*/
	var seckillBuyReq model2.SeckillBuyReq
	err := c.ShouldBind(&seckillBuyReq)
	if err != nil {
		log.Printf("bind json err: %+v\n", err)
		c.JSON(200, service.GetResponse(service.ERR_JSON_BIND, service.GetErrMsg(service.ERR_JSON_BIND), nil))
		return
	}
	// 设置秒杀记录状态
	ok, err := h.seckillRecordService.SeckillRecordStatusChange(seckillBuyReq.SeckillNum, enum.SECKILL_ORDER_PAYMENT_SUCCESSFUL)
	if !ok {
		c.JSON(200, service.GetResponse(service.SUCCESS, service.GetErrMsg(service.ERR_UPDATE_STATUS_ORDER), nil))
		return
	}
	// 根据秒杀号获取秒杀记录，获取订单号
	log.Printf("seckill %v", seckillBuyReq.SeckillNum)
	records, err := h.seckillRecordService.GetSeckillRecord(seckillBuyReq.SeckillNum)
	if err != nil {
		return
	}
	log.Printf("records:%v.", records)
	orderNum := *records[0].OrderNum

	// 设置订单状态
	ok, err = h.orderService.OrderStatusChange(orderNum, enum.SECKILL_ORDER_PAYMENT_SUCCESSFUL)
	if !ok {
		c.JSON(200, service.GetResponse(service.SUCCESS, service.GetErrMsg(service.ERR_UPDATE_STATUS_ORDER), nil))
		return
	}
	// 删除redis记录
	ok, err = h.seckillRecordService.DeleteRedisSeckillRecord(seckillBuyReq.SeckillNum)
	if !ok {
		log.Printf("DeleteRedisSeckillRecord err caused by %v.", err)
		c.JSON(200, service.GetResponse(service.ERR_BUY_ORDER, service.GetErrMsg(service.ERR_BUY_ORDER), nil))
		return
	}
	c.JSON(200, service.GetResponse(service.SUCCESS, service.GetErrMsg(service.SUCCESS), nil))
	return
}

func (h *SeckillHandler) Cancel(c *gin.Context) {
	var seckillCancelReq model2.SeckillCancelReq
	err := c.ShouldBind(&seckillCancelReq)
	if err != nil {
		log.Printf("bind json err: %+v\n", err)
		c.JSON(service.ERR_JSON_BIND, service.GetResponse(service.ERR_JSON_BIND, service.GetErrMsg(service.ERR_JSON_BIND), nil))
		return
	}
	// 设置秒杀记录状态
	ok, err := h.seckillRecordService.SeckillRecordStatusChange(seckillCancelReq.SeckillNum, enum.SECKILL_ORDER_CANCELLED)
	if !ok {
		c.JSON(200, service.GetResponse(service.SUCCESS, service.GetErrMsg(service.ERR_UPDATE_STATUS_ORDER), nil))
		return
	}
	// 根据秒杀号获取秒杀记录，获取订单号
	log.Printf("seckill %v", seckillCancelReq.SeckillNum)
	records, err := h.seckillRecordService.GetSeckillRecord(seckillCancelReq.SeckillNum)
	if err != nil {
		return
	}
	log.Printf("records:%v.", records)
	orderNum := *records[0].OrderNum

	// 设置订单状态
	ok, err = h.orderService.OrderStatusChange(orderNum, enum.SECKILL_ORDER_CANCELLED)
	if !ok {
		c.JSON(200, service.GetResponse(service.SUCCESS, service.GetErrMsg(service.ERR_UPDATE_STATUS_ORDER), nil))
		return
	}
	// 删除redis记录
	ok, err = h.seckillRecordService.DeleteRedisSeckillRecord(seckillCancelReq.SeckillNum)
	if !ok {
		log.Printf("DeleteRedisSeckillRecord err caused by %v.", err)
		c.JSON(200, service.GetResponse(service.ERR_BUY_ORDER, service.GetErrMsg(service.ERR_BUY_ORDER), nil))
		return
	}
	c.JSON(200, service.GetResponse(service.SUCCESS, service.GetErrMsg(service.SUCCESS), nil))
	return
}

func (h *SeckillHandler) GetSeckillRecord(c *gin.Context) {
	var seckillRecordGetReq model2.SeckillRecordGetReq
	var seckillRecords []*model3.SeckillRecordResp
	if err := c.ShouldBind(&seckillRecordGetReq); err != nil {
		log.Printf("bind json err: %+v\n", err)
		c.JSON(200, service.GetResponse(service.ERR_JSON_BIND, service.GetErrMsg(service.ERR_JSON_BIND), seckillRecords))
		return
	}
	log.Printf("seckillRecordGetReq %+v", seckillRecordGetReq)
	seckillRecords, err := h.seckillRecordService.GetSeckillRecordByUser(seckillRecordGetReq)
	if err != nil {
		log.Printf("GetSeckillRecordByUser err caused by: %+v\n", err)
		c.JSON(200, service.GetResponse(service.ERR_GET_SECKILL_RECORD_FAILED, service.GetErrMsg(service.ERR_GET_SECKILL_RECORD_FAILED), seckillRecords))
		return
	}
	c.JSON(200, service.GetResponse(service.SUCCESS, service.GetErrMsg(service.SUCCESS), seckillRecords))
	return
}

func (h *SeckillHandler) GetSeckillRecordLast(c *gin.Context) {
	var seckillRecordGetReq model2.SeckillRecordGetReq
	var seckillRecords []*model.SeckillRecord
	if err := c.ShouldBind(&seckillRecordGetReq); err != nil {
		log.Printf("bind json err: %+v\n", err)
		c.JSON(200, service.GetResponse(service.ERR_JSON_BIND, service.GetErrMsg(service.ERR_JSON_BIND), seckillRecords))
		return
	}
	log.Printf("seckillRecordGetReq %+v", seckillRecordGetReq)
	seckillRecords, err := h.seckillRecordService.GetSeckillRecordByUserLast(seckillRecordGetReq)
	if err != nil {
		log.Printf("GetSeckillRecordByUser err caused by: %+v\n", err)
		c.JSON(200, service.GetResponse(service.ERR_GET_SECKILL_RECORD_FAILED, service.GetErrMsg(service.ERR_GET_SECKILL_RECORD_FAILED), seckillRecords))
		return
	}
	c.JSON(200, service.GetResponse(service.SUCCESS, service.GetErrMsg(service.SUCCESS), seckillRecords))
	return
}

func (h *SeckillHandler) GetSeckillRecordBySecNum(c *gin.Context) {
	var seckillRecordGetReq model2.SeckillRecordBySecNumReq
	var seckillRecords []*model.SeckillRecord
	if err := c.ShouldBind(&seckillRecordGetReq); err != nil {
		log.Printf("bind json err: %+v\n", err)
		c.JSON(200, service.GetResponse(service.ERR_JSON_BIND, service.GetErrMsg(service.ERR_JSON_BIND), seckillRecords))
		return
	}
	if seckillRecordGetReq.SecNum == nil {
		c.JSON(200, service.GetResponse(service.ERR_SECNUM_IS_NIL, service.GetErrMsg(service.ERR_SECNUM_IS_NIL), nil))
		return
	}
	log.Printf("seckillRecordGetReq %+v", seckillRecordGetReq)
	seckillRecords, err := h.seckillRecordService.GetSeckillRecordBySecNum(seckillRecordGetReq)
	if err != nil {
		log.Printf("GetSeckillRecordByUser err caused by: %+v\n", err)
		c.JSON(200, service.GetResponse(service.ERR_GET_SECKILL_RECORD_FAILED, service.GetErrMsg(service.ERR_GET_SECKILL_RECORD_FAILED), seckillRecords))
		return
	}
	c.JSON(200, service.GetResponse(service.SUCCESS, service.GetErrMsg(service.SUCCESS), seckillRecords))
	return
}
