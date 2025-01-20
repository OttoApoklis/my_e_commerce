package service

import (
	"gorm.io/gorm"
	"my_e_commerce/data/dal/model"
)

type StockService interface {
	CreateStock(seckillStock *model.SeckillStock) error
	SubStock(db *gorm.DB, goodsId uint32, num uint32) (bool, error)
}
