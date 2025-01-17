package serviceImpl

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"my_e_commerce/config"
	"my_e_commerce/data/dal/model"
	model2 "my_e_commerce/data/req"
)

type GoodsServiceImpl struct {
}

func NewGoodsServiceImpl() *GoodsServiceImpl { return &GoodsServiceImpl{} }

func (s *GoodsServiceImpl) CreateGoods(db *gorm.DB, goods *model.Good) error {
	var goodss []model.Good
	if err := db.Where("goods_num", goods.GoodsNum).Find(&goodss).Error; err != nil {
		log.Printf("err from quota find in createGoods: %+v", err)
		return err
	}
	if goodss != nil && len(goodss) != 0 {
		log.Printf("err from goods create because of repeateable")
		return errors.New("repeatable in goods create")
	}
	err := db.Save(goods).Error
	if err != nil {
		log.Printf("error from Goods create: %+v", err)
		return err
	}
	return nil
}

func (s *GoodsServiceImpl) GetGoods(goodsNum uint32) ([]model.Good, error) {
	db := config.GetDB()
	goods := []model.Good{}
	err := db.Select("id", "goods_num", "goods_name", "price",
		"pic_url", "seller").Where("goods_num = ?", goodsNum).Find(&goods).Error
	if err != nil {
		log.Printf(" err: %+v", err)
		return nil, err
	}
	if goods == nil {
		log.Printf("Get Goods is nil")
	}
	return goods, nil
}

func (s *GoodsServiceImpl) UpdateGoods(db *gorm.DB, req *model2.GoodsReq) error {
	var goods model2.GoodsReq
	goods = *req
	dbMessage := db.Model(&model.Good{}).
		Where("goods_num = ?", *goods.GoodsNum).
		Limit(1).
		Update("price", goods.Price)
	//if dbMessage.RowsAffected == 0 {
	//	return errors.New("查不到该数据")
	//}
	err := dbMessage.Error
	if err != nil {
		log.Printf("error from Goods update: %+v", err)
		return err
	}
	return nil
}

func (s *GoodsServiceImpl) DeleteGoodsById(db *gorm.DB, ID uint32) error {
	dbMessage := db.Where("id = ?", ID).Delete(&model.Good{})
	if dbMessage.RowsAffected == 0 {
		log.Printf("rows affected is zero in DeleteGoodsById")
		return errors.New("查不到数据")
	}
	err := dbMessage.Error
	if err != nil {
		log.Printf("error from Goods delete: %+v", err)
	}
	return nil
}
