package service

import (
	"gorm.io/gorm"
	model2 "my_e_commerce/data/dal/model"
	model "my_e_commerce/data/resp"
	"my_e_commerce/data/resp/page"
)

type GoodsService interface {
	CreateGoods(db *gorm.DB, goodsReq *model2.Good) error
	GetGoods(goodsId uint32) ([]model2.Good, error)
	GetGoodsInPage(goodsNum *string, size uint32, offset uint32) (page.GoodsRespPage, error)
	UpdateGoods(db *gorm.DB, quotum *model.GoodsReq) error
	DeleteGoodsById(db *gorm.DB, ID uint32) error
}
