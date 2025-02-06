package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"my_e_commerce/config"
	"my_e_commerce/data/dal/model"
	"my_e_commerce/data/filter"
	model2 "my_e_commerce/data/req"
	"my_e_commerce/data/req/page"
	"my_e_commerce/data/response"
	"my_e_commerce/enum"
	"my_e_commerce/service"
	"my_e_commerce/utils"
)

type GoodsHandler struct {
	goodsService        service.GoodsService
	seckillStockService service.SeckillStockService
}

func NewGoodsHandler(goodsService service.GoodsService, seckillStockService service.SeckillStockService) *GoodsHandler {
	return &GoodsHandler{goodsService: goodsService, seckillStockService: seckillStockService}
}

func (h *GoodsHandler) CreateGoods(c *gin.Context) {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("createGoods err: %+v", err)
		}
	}()
	var newGoods model2.GoodsCreateReq
	if err := c.ShouldBind(&newGoods); err != nil {
		log.Printf("error from goods create BindJSON: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_JSON_BIND,
				response.GetErrMsg(response.ERR_JSON_BIND),
				err.Error()))
		return
	}
	db := config.GetDB()
	var insertGoods model.Good
	utils.CopyStruct(newGoods, insertGoods)
	snowCode := utils.GetSnowCode()
	numPrefix, err := enum.GetNumPrefix(*newGoods.GoodsType)
	if err != nil {
		log.Printf("getNumPrefix failed cause by: %v.", err)
	}
	GoodsNum := fmt.Sprintf("%s%d", numPrefix, snowCode)
	insertGoods.GoodsNum = &GoodsNum
	if err := h.goodsService.CreateGoods(db, &insertGoods); err != nil {
		log.Printf("err from create Goods: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_CREATE_GOODS_FAILED,
				response.GetErrMsg(response.ERR_CREATE_GOODS_FAILED),
				err.Error()))
		return
	}
	id := insertGoods.ID
	seckillStock := model2.SeckillStockReq{Stock: newGoods.Stock, GoodsID: &id}
	if _, err := h.seckillStockService.CreateSeckillStock(seckillStock); err != nil {
		log.Printf("Create Goods SUCCESS, but err from create SeckillStock: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_CREATE_GOODS_FAILED,
				response.GetErrMsg(response.ERR_CREATE_GOODS_FAILED),
				err.Error()))
		return
	}
	c.JSON(200,
		response.GetResponse(
			response.SUCCESS_CREATE_GOODS,
			response.GetErrMsg(response.SUCCESS_CREATE_GOODS),
			insertGoods))
	return
}

func (h *GoodsHandler) GetGoods(c *gin.Context) {
	var filter filter.GoodFilter
	if err := c.BindJSON(&filter); err != nil {
		log.Printf("error frorm goods get BindJSON: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_JSON_BIND,
				response.GetErrMsg(response.ERR_JSON_BIND),
				err.Error()))
		return
	}
	fmt.Printf("\n filter: %+v", filter)
	var goodss []model.Good
	goodss, err := h.goodsService.GetGoods(filter.ID)
	if err != nil {
		log.Printf("err from get Goods: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_GET_GOODS_FAILED,
				response.GetErrMsg(response.ERR_GET_GOODS_FAILED),
				err.Error()))
		return
	}
	c.JSON(200,
		response.GetResponse(
			response.SUCCESS_GET_GOODS,
			response.GetErrMsg(response.SUCCESS_GET_GOODS),
			goodss))
	return
}

func (h *GoodsHandler) GetGoodsInPage(c *gin.Context) {
	var req model2.GoodsGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		log.Printf("error frorm goods get BindJSON: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_JSON_BIND,
				response.GetErrMsg(response.ERR_JSON_BIND),
				err.Error()))
		return
	}
	var goodsPage page.GoodsRespPage
	goodsPage, err := h.goodsService.GetGoodsInPage(req.GoodsNum, req.PageSize, req.PageNum)
	if err != nil {
		log.Printf("err from get Goods: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_GET_GOODS_FAILED,
				response.GetErrMsg(response.ERR_GET_GOODS_FAILED),
				err.Error()))
		return
	}
	c.JSON(200,
		response.GetResponse(
			response.SUCCESS_GET_GOODS,
			response.GetErrMsg(response.SUCCESS_GET_GOODS),
			goodsPage))
	return
}

func (h *GoodsHandler) UpdateGoods(c *gin.Context) {
	var goodReq model2.GoodsUpdateReq
	if err := c.BindJSON(&goodReq); err != nil {
		log.Printf("error from goods update BindJSON: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_JSON_BIND,
				response.GetErrMsg(response.ERR_JSON_BIND),
				err.Error()))
		return
	}
	fmt.Printf("/n goodReq :%+v \n", goodReq)
	db := config.GetDB()
	if err := h.goodsService.UpdateGoods(db, &goodReq); err != nil {
		log.Printf("err from update Goods: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_CREATE_GOODS_FAILED,
				response.GetErrMsg(response.ERR_UPDATE_GOODS_FAILED),
				err.Error()))
		return
	}
	c.JSON(200,
		response.GetResponse(
			response.SUCCESS_UPDATE_GOODS,
			response.GetErrMsg(response.SUCCESS_UPDATE_GOODS),
			nil))
	return
}

func (h *GoodsHandler) DeleteGoodsById(c *gin.Context) {
	var good model.Good
	if err := c.BindJSON(&good); err != nil {
		log.Printf("error from goods delete BindJSON: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_JSON_BIND,
				response.GetErrMsg(response.ERR_JSON_BIND),
				err.Error()))
		return
	}
	db := config.GetDB()
	if err := h.goodsService.DeleteGoodsById(db, good.ID); err != nil {
		log.Printf("err from goods delete by id: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_DELETE_GOODS_FAILED,
				response.GetErrMsg(response.ERR_DELETE_GOODS_FAILED),
				err.Error()))
		return
	}
	c.JSON(200,
		response.GetResponse(
			response.SUCCESS_DELETE_BY_ID_GOODS,
			response.GetErrMsg(response.SUCCESS_DELETE_BY_ID_GOODS),
			nil))
	return
}
