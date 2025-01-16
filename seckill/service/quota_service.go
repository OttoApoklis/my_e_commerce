package service

import (
	"gorm.io/gorm"
	"my_e_commerce/data/dal/model"
	"my_e_commerce/data/filter"
	model2 "my_e_commerce/data/req"
)

type QuotaService interface {
	CreateQuota(db *gorm.DB, quotum *model.Quotum) error
	UpdateQuota(db *gorm.DB, quotum *model2.QuotumReq) error
	DeleteQuotaById(db *gorm.DB, ID uint32) error
	DeleteQuotaByCondition(db *gorm.DB, filter *filter.QuotumFilter) error
	GetQuota(db *gorm.DB, filter *filter.QuotumFilter) ([]model.Quotum, error)
}
