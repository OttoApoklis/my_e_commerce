package serviceImpl

import (
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
	tx := db.Model(&model.SeckillStock{}).
		Where("goods_id = ? and stock >= ?", goodsId, num).
		Update("stock", gorm.Expr("stock - ?", num))
	if tx.Error != nil {
		return false, tx.Error
	}
	if tx.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}
