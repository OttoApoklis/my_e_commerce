package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"golang.org/x/net/context"
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
	userID := 1
	// 商品id 购买数量
	// 先创建订单再创建秒杀订单
	var seckillReq model2.SeckillReq
	err := c.BindJSON(&seckillReq)
	if err != nil {
		log.Printf("bind json err: %+v\n", err)
		c.JSON(service.ERR_JSON_BIND, service.GetResponse(service.GetErrMsg(service.ERR_JSON_BIND), nil))
		return
	}
	// 校验参数
	if seckillReq.GoodsAmount <= 0 {
		c.JSON(service.ERR_INPUT_INVALID, service.GetResponse(service.GetErrMsg(service.ERR_INPUT_INVALID), nil))
	}
	// TODO 先去查redis库存数据
	// 读取.lua 文件的内容
	scriptBytes, err := ioutil.ReadFile("./lua/redis_substock.lua")
	if err != nil {
		fmt.Println("Error reading Lua script file:", err)
		return
	}
	script := string(scriptBytes)
	fmt.Printf("\n\n\nscript:" + script + "\n")
	rdb, err := config.GetRedisConnection()
	if err != nil {
		fmt.Println("Error GetRedisConnection:", err)
		return
	}
	ctx := context.Background()
	stock := rdb.Get(ctx, strconv.Itoa(int(seckillReq.GoodsID)))
	fmt.Printf("redis stock : %+v\n", stock)

	result, err := rdb.Eval(ctx, script, []string{strconv.Itoa(int(seckillReq.GoodsID))}, seckillReq.GoodsAmount).Result()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	stock = rdb.Get(ctx, strconv.Itoa(int(seckillReq.GoodsID)))
	fmt.Printf("redis stock : %+v\n", stock)
	fmt.Printf("redis subStock result: %+v", result)
	// 数据库扣库存
	ok, err := h.stockService.SubStock(seckillReq.GoodsID, seckillReq.GoodsAmount)
	if err != nil {
		log.Printf("this is err , %+v\n", err)
		c.JSON(service.ERR_DB, service.GetResponse(service.GetErrMsg(service.ERR_DB), nil))
		return
	}
	if !ok {
		c.JSON(service.ERR_DESC_STOCK_FAILED, service.GetResponse(service.GetErrMsg(service.ERR_DESC_STOCK_FAILED), nil))
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
		c.JSON(service.ERR_GET_GOODS_FAILED, service.GetResponse(service.GetErrMsg(service.ERR_GET_GOODS_FAILED), nil))
	}
	if goods == nil || len(goods) == 0 {
		c.JSON(service.ERR_GET_GOODS_EMPTY_FAILED, service.GetResponse(service.GetErrMsg(service.ERR_GET_GOODS_EMPTY_FAILED), nil))
	}
	good := goods[0]
	fmt.Printf("good : %+v\n", good)
	orderReq.GoodsID = seckillReq.GoodsID
	orderReq.GoodsNum = good.GoodsNum
	orderReq.Price = 18.00
	orderReq.Status = 0
	orderReq.GoodsAmount = &seckillReq.GoodsAmount
	orderReq.Buyer = uint32(userID)
	orderReq.Seller = good.Seller
	fmt.Printf("orderReq : %+v\n", orderReq)
	var orderID uint32
	orderID, err = h.orderService.CreateOrder(orderReq)
	if err != nil {
		log.Fatalln(err)
		c.JSON(service.ERR_CREATE_ORDER_FAILED, service.GetResponse(service.GetErrMsg(service.ERR_CREATE_ORDER_FAILED), nil))
		return
	}
	// TODO 创建秒杀单
	c.JSON(service.SUCCESS, service.GetResponse(service.GetErrMsg(service.SUCCESS), orderID))
	return

}

func MultiplyFloatsAsDecimals(price float32, amount uint32) float64 {
	// 将 float64 转换为 Decimal
	decA := decimal.NewFromFloat(float64(price))

	decB := decimal.NewFromFloat(float64(amount))
	// 进行乘法运算
	result := decA.Mul(decB)

	return result.InexactFloat64()
}
