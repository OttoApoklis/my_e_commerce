package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"my_e_commerce/config"
	"my_e_commerce/data/dal/model"
	"my_e_commerce/data/filter"
	model2 "my_e_commerce/data/resp"
	"my_e_commerce/data/response"
	"my_e_commerce/service"
)

type QuotaHandler struct {
	quotaService service.QuotaService
}

func NewQuotaHandler(quotaService service.QuotaService) *QuotaHandler {
	return &QuotaHandler{quotaService: quotaService}
}

func getId(ctx context.Context) {
	// 隐式转换

}

func (h *QuotaHandler) CreateQuota(c *gin.Context) {
	getId(c)
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("createQuota err: %+v", err)
		}
	}()
	var newQuota model.Quotum
	if err := c.BindJSON(&newQuota); err != nil {
		log.Printf("error from quota create BindJSON: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_JSON_BIND,
				response.GetErrMsg(response.ERR_JSON_BIND),
				err.Error()))
		return
	}
	db := config.GetDB()
	if err := h.quotaService.CreateQuota(db, &newQuota); err != nil {
		log.Printf("err from create Quota: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_CREATE_USER_QUOTA_FAILED,
				response.GetErrMsg(response.ERR_CREATE_USER_QUOTA_FAILED),
				err.Error()))
		return
	}
	c.JSON(200,
		response.GetResponse(
			response.SUCCESS_CREATE_USER_QUOTA,
			response.GetErrMsg(response.SUCCESS_CREATE_USER_QUOTA),
			newQuota))
	return
}

func (h *QuotaHandler) GetQuota(c *gin.Context) {
	var filter filter.QuotumFilter
	if err := c.BindJSON(&filter); err != nil {
		log.Printf("error frorm quota get BindJSON: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_JSON_BIND,
				response.GetErrMsg(response.ERR_JSON_BIND),
				err.Error()))
		return
	}
	fmt.Printf("\n filter: %+v", filter)
	db := config.GetDB()
	var quotas []model.Quotum
	quotas, err := h.quotaService.GetQuota(db, &filter)
	if err != nil {
		log.Printf("err from get Quota: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_GET_USER_QUOTA_FAILED,
				response.GetErrMsg(response.ERR_GET_USER_QUOTA_FAILED),
				err.Error()))
		return
	}
	c.JSON(200,
		response.GetResponse(
			response.SUCCESS_GET_USER_QUOTA,
			response.GetErrMsg(response.SUCCESS_GET_USER_QUOTA),
			quotas))
	return
}

func (h *QuotaHandler) UpdateQuota(c *gin.Context) {
	var quotumReq model2.QuotumReq
	if err := c.BindJSON(&quotumReq); err != nil {
		log.Printf("error from quota update BindJSON: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_JSON_BIND,
				response.GetErrMsg(response.ERR_JSON_BIND),
				err.Error()))
		return
	}
	fmt.Printf("/n quotumReq :%+v \n", quotumReq)
	db := config.GetDB()
	if err := h.quotaService.UpdateQuota(db, &quotumReq); err != nil {
		log.Printf("err from update Quota: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_CREATE_USER_QUOTA_FAILED,
				response.GetErrMsg(response.ERR_UPDATE_USER_QUOTA_FAILED),
				err.Error()))
		return
	}
	c.JSON(200,
		response.GetResponse(
			response.SUCCESS_UPDATE_USER_QUOTA,
			response.GetErrMsg(response.SUCCESS_UPDATE_USER_QUOTA),
			nil))
	return
}

func (h *QuotaHandler) DeleteQuotaById(c *gin.Context) {
	var quotum model.Quotum
	if err := c.BindJSON(&quotum); err != nil {
		log.Printf("error from quota delete BindJSON: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_JSON_BIND,
				response.GetErrMsg(response.ERR_JSON_BIND),
				err.Error()))
		return
	}
	db := config.GetDB()
	if err := h.quotaService.DeleteQuotaById(db, quotum.ID); err != nil {
		log.Printf("err from quota delete by id: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_DELETE_BY_ID_USER_QUOTA_FAILED,
				response.GetErrMsg(response.ERR_DELETE_BY_ID_USER_QUOTA_FAILED),
				err.Error()))
		return
	}
	c.JSON(200,
		response.GetResponse(
			response.SUCCESS_DELETE_BY_ID_USER_QUOTA,
			response.GetErrMsg(response.SUCCESS_DELETE_BY_ID_USER_QUOTA),
			nil))
	return
}
