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
)

type UserQuotaHandler struct {
	userQuotaService service.UserQuotaService
}

func NewUserQuotaHandler(userQutoaService service.UserQuotaService) *UserQuotaHandler {
	return &UserQuotaHandler{userQuotaService: userQutoaService}
}

func (h *UserQuotaHandler) CreateUserQuota(c *gin.Context) {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("createUserQuota err: %+v", err)
		}
	}()
	var newUserQutoa model.UserQuotum
	if err := c.BindJSON(&newUserQutoa); err != nil {
		log.Printf("error from userQuota create BindJSON: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_JSON_BIND,
				response.GetErrMsg(response.ERR_JSON_BIND),
				err.Error()))
		return
	}
	db := config.GetDB()
	if err := h.userQuotaService.CreateUserQuota(db, &newUserQutoa); err != nil {
		log.Printf("err from create UserQuota: %+v", err)
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
			newUserQutoa))
	return
}

func (h *UserQuotaHandler) GetUserQuota(c *gin.Context) {
	var filter filter.UserQuotumFilter
	if err := c.BindJSON(&filter); err != nil {
		log.Printf("error frorm userQuota get BindJSON: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_JSON_BIND,
				response.GetErrMsg(response.ERR_JSON_BIND),
				err.Error()))
		return
	}
	fmt.Printf("\n filter: %+v", filter)
	db := config.GetDB()
	var userQuotas []model.UserQuotum
	userQuotas, err := h.userQuotaService.GetUserQuoto(db, &filter)
	if err != nil {
		log.Printf("err from get UserQuota: %+v", err)
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
			userQuotas))
	return
}

func (h *UserQuotaHandler) UpdateUserQuota(c *gin.Context) {
	var userQuotumReq model2.UserQuotumReq
	if err := c.BindJSON(&userQuotumReq); err != nil {
		log.Printf("error from userQuota update BindJSON: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_JSON_BIND,
				response.GetErrMsg(response.ERR_JSON_BIND),
				err.Error()))
		return
	}
	db := config.GetDB()
	if err := h.userQuotaService.UpdateUserQuota(db, &userQuotumReq); err != nil {
		log.Printf("err from update UserQuota: %+v", err)
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

func (h *UserQuotaHandler) DeleteUserQuotaById(c *gin.Context) {
	var quotum model.Quotum
	if err := c.BindJSON(&quotum); err != nil {
		log.Printf("error from userQuota delete BindJSON: %+v", err)
		c.JSON(200,
			response.GetResponse(
				response.ERR_JSON_BIND,
				response.GetErrMsg(response.ERR_JSON_BIND),
				err.Error()))
		return
	}
	db := config.GetDB()
	if err := h.userQuotaService.DeleteUserQuotaById(db, quotum.ID); err != nil {
		log.Printf("err from userQuota delete by id: %+v", err)
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
