package serviceImpl

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"my_e_commerce/config"
	model2 "my_e_commerce/data/dal/model"
	model "my_e_commerce/data/resp"
	"my_e_commerce/utils"
	"strconv"
)

type SeckillStockServiceImpl struct{}

func (s *SeckillStockServiceImpl) UpdateSeckillStock(req model.SeckillStockReq) error {
	// TODO 更新做的更全面一些
	db := config.GetDB()
	fmt.Printf("goods_id %+v", *req.GoodsID)
	dbMessage := db.Table("seckill_stock").
		Where("goods_id = ?", *req.GoodsID).
		Update("stock", req.Stock)
	// TODO 修改判断条件，应该是查不到该行时报出错误，因为如果请求数据与当前数据一致也是影响0行
	if dbMessage.RowsAffected == 0 {
		return errors.New("查不到该数据")
	}
	err := dbMessage.Error
	if err != nil {
		log.Printf("error from SeckillStock update: %+v", err)
		return err
	}
	return nil
}

func (s *SeckillStockServiceImpl) GetSeckillStock(SeckillStock model.SeckillStockReq) ([]*model2.SeckillStock, error) {
	db := config.GetDB()
	SeckillStocks := []*model2.SeckillStock{}
	err := db.Where("goods_id = ?", SeckillStock.GoodsID).Find(&SeckillStocks).Error
	if err != nil {
		log.Printf(" err: %+v", err)
		return nil, err
	}
	if SeckillStocks == nil {
		log.Printf("Get SeckillStock is nil")
	}
	return SeckillStocks, nil
}

func NewSeckillStockServiceImpl() *SeckillStockServiceImpl { return &SeckillStockServiceImpl{} }

func (s *SeckillStockServiceImpl) CreateSeckillStock(SeckillStock model.SeckillStockReq) (uint32, error) {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("err : %+v", err)
			return
		}
	}()
	db := config.GetDB()
	var seckillStockInsert model2.SeckillStock
	fmt.Printf("SeckillStock %+v", SeckillStock)
	utils.CopyStruct(&SeckillStock, &seckillStockInsert)
	fmt.Printf("seckillStockInsert %+v", seckillStockInsert)
	fmt.Println(seckillStockInsert)
	tx := db.Save(&seckillStockInsert)
	utils.RedisSetStock(strconv.FormatUint(uint64(*seckillStockInsert.GoodsID), 10), int(seckillStockInsert.Stock))
	fmt.Printf("affect :%+v", tx.RowsAffected)
	return seckillStockInsert.ID, tx.Error
}

func (s *SeckillStockServiceImpl) DeleteSeckillStockById(db *gorm.DB, ID uint32) error {
	dbMessage := db.Where("goods_id = ?", ID).Delete(&model2.SeckillStock{})
	if dbMessage.RowsAffected == 0 {
		log.Printf("rows affected is zero in DeleteSeckillStockById")
		return errors.New("查不到数据")
	}
	err := dbMessage.Error
	if err != nil {
		log.Printf("error from SeckillStocks delete: %+v", err)
	}
	return nil
}
