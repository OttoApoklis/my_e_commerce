package routes

import (
	"github.com/gin-gonic/gin"
	"my_e_commerce/controller"
	"my_e_commerce/service/serviceImpl"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	userGroup := router.Group("/users")
	userService := serviceImpl.NewUserServiceImpl()
	userHandler := controller.NewUserHandler(userService)
	{
		userGroup.POST("/save", userHandler.CreateUser)
		userGroup.POST("/get", userHandler.SelectByUserId)
		userGroup.POST("/update", userHandler.UpdateUserById)
	}
	seckillGroup := router.Group("/seckill")
	seckillRecordService := serviceImpl.NewSeckillServiceImpl()
	orderService := serviceImpl.NewOrderServiceImpl()
	stockService := serviceImpl.NewStockServiceImpl()
	goodsService := serviceImpl.NewGoodsServiceImpl()
	seckillHandler := controller.NewSeckillHanlder(seckillRecordService,
		orderService, stockService, goodsService)
	{
		seckillGroup.POST("/", seckillHandler.CreateSeckill)
	}

	return router
}
