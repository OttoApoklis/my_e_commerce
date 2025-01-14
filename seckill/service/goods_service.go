package service

import (
	model2 "my_e_commerce/data/dal/model"
	model "my_e_commerce/data/req"
)

type GoodsService interface {
	CreateGoods(goodsReq *model.GoodsReq) error
	GetGoods(goodsId uint32) ([]model2.Good, error)
}
