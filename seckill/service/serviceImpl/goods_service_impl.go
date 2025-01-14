package serviceImpl

import (
	"log"
	"my_e_commerce/config"
	"my_e_commerce/data/dal/model"
	model2 "my_e_commerce/data/req"
)

type GoodsServiceImpl struct {
}

func (s *GoodsServiceImpl) CreateGoods(goodsReq *model2.GoodsReq) error {
	//TODO implement me
	panic("implement me")
}

func NewGoodsServiceImpl() *GoodsServiceImpl {
	return &GoodsServiceImpl{}
}

func (s *GoodsServiceImpl) GetGoods(goodsId uint32) ([]model.Good, error) {
	db := config.GetDB()
	goods := []model.Good{}
	err := db.Select("id", "goods_num", "goods_name", "price",
		"pic_url", "seller").Where("id = ?", goodsId).Find(&goods).Error
	if err != nil {
		log.Printf(" err: %+v", err)
		return nil, err
	}
	if goods == nil {
		log.Printf("Get Goods is nil")
	}
	return goods, nil
}
