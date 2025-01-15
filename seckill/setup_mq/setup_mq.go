package setup_mq

import (
	"my_e_commerce/config"
	"my_e_commerce/controller"
	"my_e_commerce/service/serviceImpl"
	"time"

	"golang.org/x/time/rate"
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
