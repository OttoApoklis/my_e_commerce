package routes

import (
	"my_e_commerce/controller"
	"my_e_commerce/service/serviceImpl"

	"github.com/gin-gonic/gin"
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
	seckillStockService := serviceImpl.NewSeckillStockServiceImpl()
	seckillHandler := controller.NewSeckillHanlder(seckillRecordService,
		orderService, stockService, goodsService)
	{
		seckillGroup.POST("/", seckillHandler.CreateSeckill)
		seckillGroup.POST("/buy", seckillHandler.Buy)
		seckillGroup.POST("/cancel", seckillHandler.Cancel)
		seckillGroup.POST("/get", seckillHandler.GetSeckillRecord)
		seckillGroup.POST("/getLast", seckillHandler.GetSeckillRecordLast)
		seckillGroup.POST("/getBySecNum", seckillHandler.GetSeckillRecordBySecNum)
	}

	userQuotaGroup := router.Group("/userQuota")
	userQuotaService := serviceImpl.NewUserQuotaServiceImpl()
	userQuotaHandler := controller.NewUserQuotaHandler(userQuotaService)
	{
		userQuotaGroup.POST("/create", userQuotaHandler.CreateUserQuota)
		userQuotaGroup.POST("/get", userQuotaHandler.GetUserQuota)
		userQuotaGroup.POST("/update", userQuotaHandler.UpdateUserQuota)
		userQuotaGroup.POST("/deleteById", userQuotaHandler.DeleteUserQuotaById)
	}

	quotaGroup := router.Group("/quota")
	quotaService := serviceImpl.NewQuotaServiceImpl()
	quotaHandler := controller.NewQuotaHandler(quotaService)
	{
		quotaGroup.POST("/create", quotaHandler.CreateQuota)
		quotaGroup.POST("/get", quotaHandler.GetQuota)
		quotaGroup.POST("/update", quotaHandler.UpdateQuota)
		quotaGroup.POST("/deleteById", quotaHandler.DeleteQuotaById)
	}

	goodsGroup := router.Group("/goods")
	goodsHandler := controller.NewGoodsHandler(goodsService, seckillStockService)
	{
		goodsGroup.POST("/create", goodsHandler.CreateGoods)
		goodsGroup.POST("/get", goodsHandler.GetGoods)
		goodsGroup.POST("/getPage", goodsHandler.GetGoodsInPage)
		goodsGroup.POST("/update", goodsHandler.UpdateGoods)
		goodsGroup.POST("/deleteById", goodsHandler.DeleteGoodsById)
	}

	orderGroup := router.Group("/order")
	orderHandler := controller.NewOrderHandler(orderService)
	{
		orderGroup.POST("/create", orderHandler.CreateOrder)
		orderGroup.POST("/get", orderHandler.GetOrderByUser)
		orderGroup.POST("/getBySeller", orderHandler.GetOrderBySeller)
		orderGroup.POST("/update", orderHandler.UpdateOrder)
		orderGroup.POST("/deleteById", orderHandler.DeleteOrderById)
	}

	SeckillStockGroup := router.Group("/seckillStock")

	SeckillStockHandler := controller.NewSeckillStockHandler(seckillStockService)
	{
		SeckillStockGroup.POST("/create", SeckillStockHandler.CreateSeckillStock)
		SeckillStockGroup.POST("/get", SeckillStockHandler.GetSeckillStock)
		SeckillStockGroup.POST("/update", SeckillStockHandler.UpdateSeckillStock)
		SeckillStockGroup.POST("/deleteById", SeckillStockHandler.DeleteSeckillStockById)
	}
	return router
}
