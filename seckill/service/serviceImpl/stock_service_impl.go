package serviceImpl

import (
	"log"
	"my_e_commerce/data/dal/model"

	"gorm.io/gorm"
)

type StockServiceImpl struct{}

func (s *StockServiceImpl) CreateStock(seckillStock *model.SeckillStock) error {
	//TODO implement me
	panic("implement me")
}

func NewStockServiceImpl() *StockServiceImpl {
	return &StockServiceImpl{}
}

func (s *StockServiceImpl) SubStock(db *gorm.DB, goodsId uint32, num uint32) (bool, error) {
	log.Printf("goodsId:%d, %d", goodsId, num)
	tx := db.Model(&model.SeckillStock{}).
		Where("goods_id = ? and stock >= ?", goodsId, num).
		Update("stock", gorm.Expr("stock - ?", num))
	if tx.Error != nil {
		log.Printf("更新失败")
		return false, tx.Error
	}
	if tx.RowsAffected == 0 {
		log.Printf("影响行数为0")
		return false, nil
	}

	return true, nil
}
