package setup_mq

import (
	"log"
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
	connPool, err := config.GetRabbitmqConnectionPool(1)
	if err != nil {
		log.Fatal("GetRabbitmqConnectionPool err caused by %v", err)
	}
	// 获取连接
	conn, err := connPool.GetRabbitmqConn()
	if err != nil {
		log.Fatal("GetRabbitmqConn err caused by %v", err)
	}
	defer func() {
		// 回收连接
		connPool.PutRabbitmqConn(conn)
	}()
	go seckillHandler.ReceiveMessage(conn, "seckill", rate.Limit(1*time.Second))
}
