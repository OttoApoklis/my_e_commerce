package service

import (
	"gorm.io/gorm"
	"my_e_commerce/data/dal/model"
	model2 "my_e_commerce/data/req"
)

type SeckillStockService interface {
	CreateSeckillStock(SeckillStock model2.SeckillStockReq) (uint32, error)
	UpdateSeckillStock(SeckillStock model2.SeckillStockReq) error
	GetSeckillStock(SeckillStock model2.SeckillStockReq) ([]*model.SeckillStock, error)
	DeleteSeckillStockById(db *gorm.DB, ID uint32) error
}
