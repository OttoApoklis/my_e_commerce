package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"my_e_commerce/config"
	"my_e_commerce/data/dal/model"
	"my_e_commerce/data/filter"
	model2 "my_e_commerce/data/req"
	"my_e_commerce/data/response"
	"my_e_commerce/service"
	"my_e_commerce/utils"
)

type SeckillStockHandler struct {
	SeckillStockService service.SeckillStockService
}

func NewSeckillStockHandler(SeckillStockService service.SeckillStockService) *SeckillStockHandler {
	return &SeckillStockHandler{SeckillStockService: SeckillStockService}
}

func (h *SeckillStockHandler) CreateSeckillStock(c *gin.Context) {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("createSeckillStock err: %+v", err)
		}
	}()
	var newSeckillStock model.SeckillStock
	if err := c.BindJSON(&newSeckillStock); err != nil {
		log.Printf("error from SeckillStock create BindJSON: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_JSON_BIND,
				response.GetErrMsg(response.ERR_JSON_BIND),
				err.Error()))
		return
	}
	fmt.Printf("controller SeckillStock: %+v", newSeckillStock)
	var SeckillStockReq model2.SeckillStockReq
	utils.CopyStruct(&newSeckillStock, &SeckillStockReq)
	if _, err := h.SeckillStockService.CreateSeckillStock(SeckillStockReq); err != nil {
		log.Printf("err from create SeckillStock: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_CREATE_SECKILL_STOCK_FAILED,
				response.GetErrMsg(response.ERR_CREATE_SECKILL_STOCK_FAILED),
				err.Error()))
		return
	}
	c.JSON(200,
		response.GetResponse(
			response.SUCCESS_CREATE_SECKILL_STOCK,
			response.GetErrMsg(response.SUCCESS_CREATE_SECKILL_STOCK),
			newSeckillStock))
	return
}

func (h *SeckillStockHandler) GetSeckillStock(c *gin.Context) {
	var filter filter.SeckillStockFilter
	if err := c.BindJSON(&filter); err != nil {
		log.Printf("error frorm SeckillStock get BindJSON: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_JSON_BIND,
				response.GetErrMsg(response.ERR_JSON_BIND),
				err.Error()))
		return
	}
	fmt.Printf("\n filter: %+v", filter)
	var SeckillStocks []*model.SeckillStock
	var SeckillStockReq model2.SeckillStockReq
	utils.CopyStruct(&filter, &SeckillStockReq)
	SeckillStocks, err := h.SeckillStockService.GetSeckillStock(SeckillStockReq)
	if err != nil {
		log.Printf("err from get SeckillStock: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_GET_SECKILL_STOCK_FAILED,
				response.GetErrMsg(response.ERR_GET_SECKILL_STOCK_FAILED),
				err.Error()))
		return
	}
	c.JSON(200,
		response.GetResponse(
			response.SUCCESS_GET_SECKILL_STOCK,
			response.GetErrMsg(response.SUCCESS_GET_SECKILL_STOCK),
			SeckillStocks))
	return
}

func (h *SeckillStockHandler) UpdateSeckillStock(c *gin.Context) {
	var SeckillStockReq model2.SeckillStockReq
	if err := c.BindJSON(&SeckillStockReq); err != nil {
		log.Printf("error from SeckillStock update BindJSON: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_JSON_BIND,
				response.GetErrMsg(response.ERR_JSON_BIND),
				err.Error()))
		return
	}
	fmt.Printf("/n SeckillStockReq :%+v \n", SeckillStockReq)
	if err := h.SeckillStockService.UpdateSeckillStock(SeckillStockReq); err != nil {
		log.Printf("err from update SeckillStock: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_CREATE_SECKILL_STOCK_FAILED,
				response.GetErrMsg(response.ERR_UPDATE_SECKILL_STOCK_FAILED),
				err.Error()))
		return
	}
	c.JSON(200,
		response.GetResponse(
			response.SUCCESS_UPDATE_SECKILL_STOCK,
			response.GetErrMsg(response.SUCCESS_UPDATE_SECKILL_STOCK),
			nil))
	return
}

func (h *SeckillStockHandler) DeleteSeckillStockById(c *gin.Context) {
	var SeckillStock model.SeckillStock
	if err := c.BindJSON(&SeckillStock); err != nil {
		log.Printf("error from SeckillStock delete BindJSON: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_JSON_BIND,
				response.GetErrMsg(response.ERR_JSON_BIND),
				err.Error()))
		return
	}
	db := config.GetDB()
	if err := h.SeckillStockService.DeleteSeckillStockById(db, *SeckillStock.GoodsID); err != nil {
		log.Printf("err from SeckillStock delete by id: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_DELETE_SECKILL_STOCK_FAILED,
				response.GetErrMsg(response.ERR_DELETE_SECKILL_STOCK_FAILED),
				err.Error()))
		return
	}
	c.JSON(200,
		response.GetResponse(
			response.SUCCESS_DELETE_BY_ID_SECKILL_STOCK,
			response.GetErrMsg(response.SUCCESS_DELETE_BY_ID_SECKILL_STOCK),
			nil))
	return
}
