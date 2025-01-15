package setup_mq

import (
	"golang.org/x/time/rate"
	"my_e_commerce/config"
	"my_e_commerce/controller"
	"my_e_commerce/service/serviceImpl"
	"time"
)

func SetupMQ() {
	seckillRecordService := serviceImpl.NewSeckillServiceImpl()
	orderService := serviceImpl.NewOrderServiceImpl()
	stockService := serviceImpl.NewStockServiceImpl()
	goodsService := serviceImpl.NewGoodsServiceImpl()
	seckillHandler := controller.NewSeckillHanlder(seckillRecordService,
		orderService, stockService, goodsService)
	go seckillHandler.ReceiveMessage(config.GetRabbitmqConnection(), "seckill", rate.Limit(1*time.Second))
}
