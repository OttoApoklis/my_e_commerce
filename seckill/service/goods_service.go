package service

import (
	"gorm.io/gorm"
	model2 "my_e_commerce/data/dal/model"
	model "my_e_commerce/data/req"
	"my_e_commerce/data/req/page"
)

type GoodsService interface {
	CreateGoods(db *gorm.DB, goodsReq *model2.Good) error
	GetGoods(goodID uint32) ([]model2.Good, error)
	GetGoodsByUser(goodsReq model.GoodsGetUserReq) ([]model2.Good, error)
	GetGoodsInPage(req model.GoodsGetPageReq) (page.GoodsRespPage, error)
	UpdateGoods(db *gorm.DB, quotum *model.GoodsUpdateReq) error
	DeleteGoodsById(db *gorm.DB, ID uint32) error
}
