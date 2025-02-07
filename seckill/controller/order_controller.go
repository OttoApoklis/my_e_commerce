package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"my_e_commerce/config"
	"my_e_commerce/data/dal/model"
	model2 "my_e_commerce/data/req"
	"my_e_commerce/data/response"
	"my_e_commerce/service"
	"my_e_commerce/utils"
)

type OrderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("createOrder err: %+v", err)
		}
	}()
	var newOrder model.Order
	if err := c.BindJSON(&newOrder); err != nil {
		log.Printf("error from order create BindJSON: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_JSON_BIND,
				response.GetErrMsg(response.ERR_JSON_BIND),
				err.Error()))
		return
	}
	fmt.Printf("controller order: %+v", newOrder)
	var orderReq model2.OrderReq
	utils.CopyStruct(&newOrder, &orderReq)
	db := config.GetDB()
	if _, err := h.orderService.CreateOrder(db, orderReq); err != nil {
		log.Printf("err from create Order: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_CREATE_ORDER_FAILED,
				response.GetErrMsg(response.ERR_CREATE_ORDER_FAILED),
				err.Error()))
		return
	}
	c.JSON(200,
		response.GetResponse(
			response.SUCCESS_CREATE_ORDER,
			response.GetErrMsg(response.SUCCESS_CREATE_ORDER),
			newOrder))
	return
}

func (h *OrderHandler) GetOrderByUser(c *gin.Context) {
	var orderGet model2.OrderUserGetReq
	if err := c.ShouldBind(&orderGet); err != nil {
		log.Printf("error frorm order get BindJSON: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_JSON_BIND,
				response.GetErrMsg(response.ERR_JSON_BIND),
				err.Error()))
		return
	}
	fmt.Printf("\n filter: %+v", orderGet)
	var orders []*model.Order
	var orderReq model2.OrderReq
	utils.CopyStruct(&orderGet, &orderReq)
	orders, err := h.orderService.GetOrder(orderReq)
	if err != nil {
		log.Printf("err from get Order: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_GET_ORDER_FAILED,
				response.GetErrMsg(response.ERR_GET_ORDER_FAILED),
				err.Error()))
		return
	}
	c.JSON(200,
		response.GetResponse(
			response.SUCCESS_GET_ORDER,
			response.GetErrMsg(response.SUCCESS_GET_ORDER),
			orders))
	return
}

func (h *OrderHandler) GetOrderBySeller(c *gin.Context) {
	var orderGet model2.OrderSellerGetReq
	if err := c.ShouldBind(&orderGet); err != nil {
		log.Printf("error frorm order get BindJSON: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_JSON_BIND,
				response.GetErrMsg(response.ERR_JSON_BIND),
				err.Error()))
		return
	}
	fmt.Printf("\n filter: %+v", orderGet)
	var orders []*model.Order
	var orderReq model2.OrderReq
	utils.CopyStruct(&orderGet, &orderReq)
	orders, err := h.orderService.GetOrderBySeller(orderReq)
	if err != nil {
		log.Printf("err from get Order: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_GET_ORDER_FAILED,
				response.GetErrMsg(response.ERR_GET_ORDER_FAILED),
				err.Error()))
		return
	}
	c.JSON(200,
		response.GetResponse(
			response.SUCCESS_GET_ORDER,
			response.GetErrMsg(response.SUCCESS_GET_ORDER),
			orders))
	return
}

func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	var orderReq model2.OrderReq
	if err := c.BindJSON(&orderReq); err != nil {
		log.Printf("error from order update BindJSON: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_JSON_BIND,
				response.GetErrMsg(response.ERR_JSON_BIND),
				err.Error()))
		return
	}
	fmt.Printf("/n orderReq :%+v \n", orderReq)
	if err := h.orderService.UpdateOrder(orderReq); err != nil {
		log.Printf("err from update Order: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_CREATE_ORDER_FAILED,
				response.GetErrMsg(response.ERR_UPDATE_ORDER_FAILED),
				err.Error()))
		return
	}
	c.JSON(200,
		response.GetResponse(
			response.SUCCESS_UPDATE_ORDER,
			response.GetErrMsg(response.SUCCESS_UPDATE_ORDER),
			nil))
	return
}

func (h *OrderHandler) DeleteOrderById(c *gin.Context) {
	var order model.Order
	if err := c.BindJSON(&order); err != nil {
		log.Printf("error from order delete BindJSON: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_JSON_BIND,
				response.GetErrMsg(response.ERR_JSON_BIND),
				err.Error()))
		return
	}
	db := config.GetDB()
	if err := h.orderService.DeleteOrderById(db, order.ID); err != nil {
		log.Printf("err from order delete by id: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_DELETE_ORDER_FAILED,
				response.GetErrMsg(response.ERR_DELETE_ORDER_FAILED),
				err.Error()))
		return
	}
	c.JSON(200,
		response.GetResponse(
			response.SUCCESS_DELETE_BY_ID_ORDER,
			response.GetErrMsg(response.SUCCESS_DELETE_BY_ID_ORDER),
			nil))
	return
}
