package service

import "my_e_commerce/data/dal/model"

type StockService interface {
	CreateStock(seckillStock *model.SeckillStock) error
	SubStock(goodsId uint32, num uint32) (bool, error)
}
