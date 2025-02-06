package service

import (
	"gorm.io/gorm"
	"my_e_commerce/data/dal/model"
	"my_e_commerce/data/filter"
	model2 "my_e_commerce/data/resp"
)

type UserQuotaService interface {
	CreateUserQuota(db *gorm.DB, quotum *model.UserQuotum) error
	UpdateUserQuota(db *gorm.DB, quotum *model2.UserQuotumReq) error
	DeleteUserQuotaById(db *gorm.DB, ID uint32) error
	DeleteUserQuotaByCondition(db *gorm.DB, filter *filter.UserQuotumFilter) error
	GetUserQuoto(db *gorm.DB, filter *filter.UserQuotumFilter) ([]model.UserQuotum, error)
}
