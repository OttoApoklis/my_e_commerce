package serviceImpl

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"my_e_commerce/data/dal/model"
	"my_e_commerce/data/filter"
	model2 "my_e_commerce/data/resp"
)

type QuotaServiceImpl struct{}

func NewQuotaServiceImpl() *QuotaServiceImpl {
	return &QuotaServiceImpl{}
}

func (s *QuotaServiceImpl) CreateQuota(db *gorm.DB, quotum *model.Quotum) error {
	var quotas []model.Quotum

	if err := db.Where("goods_id = ?", quotum.GoodsID).Find(&quotas).Error; err != nil {
		log.Printf("err from quota find in createQuota: %+v", err)
		return err
	}
	if quotas != nil && len(quotas) != 0 {
		log.Printf("err from quota create because of repeateable")
		return errors.New("repeatable in quota create")
	}
	err := db.Save(quotum).Error
	if err != nil {
		log.Printf("error from Quota create: %+v", err)
		return err
	}
	return nil
}

func (s *QuotaServiceImpl) UpdateQuota(db *gorm.DB, req *model2.QuotumReq) error {
	var quotum model2.QuotumReq
	quotum = *req
	fmt.Println(quotum.GoodsID)
	dbMessage := db.Model(&model.Quotum{}).
		Where("goods_id = ?", *quotum.GoodsID).
		Limit(1).
		Update("num", quotum.Num)
	if dbMessage.RowsAffected == 0 {
		return errors.New("查不到该数据")
	}
	err := dbMessage.Error
	if err != nil {
		log.Printf("error from Quota update: %+v", err)
		return err
	}
	return nil
}

func (s *QuotaServiceImpl) DeleteQuotaById(db *gorm.DB, ID uint32) error {
	dbMessage := db.Where("id = ?", ID).Delete(&model.Quotum{})
	if dbMessage.RowsAffected == 0 {
		log.Printf("rows affected is zero in DeleteQuotaById")
		return errors.New("查不到数据")
	}
	err := dbMessage.Error
	if err != nil {
		log.Printf("error from Quota delete: %+v", err)
	}
	return nil
}

func (s *QuotaServiceImpl) DeleteQuotaByCondition(db *gorm.DB, filter *filter.QuotumFilter) error {
	err := db.Delete(filter)
	if err != nil {
		log.Printf("error form UserQuot delete: %+v", err)
	}
	return nil
}

func (s *QuotaServiceImpl) GetQuota(db *gorm.DB, filter *filter.QuotumFilter) ([]model.Quotum, error) {

	// 动态添加查询条件
	//query = utils.AddCondition(query, &filter.ID, "id")
	//query = utils.AddCondition(query, &filter.Num, "num")
	//query = utils.AddCondition(query, filter.GoodsID, "goods_id")
	//query = utils.AddCondition(query, &filter.UserID, "user_id")
	//query = utils.AddCondition(query, &filter.KilledNum, "killed_num")

	// 查询结果
	var quotas []model.Quotum
	fmt.Printf("\nfilter1: %+v\n", *filter.GoodsID)
	dbMessage := db.Table("quota").Where("goods_id = ?", *filter.GoodsID).Scan(&quotas)
	//db.Raw("select * from user_quota where user_id = ? and goods_id = ?", filter.UserID, *filter.GoodsID).Scan(&quotas)
	fmt.Printf("\nquotas: %+v\n", quotas)
	if quotas == nil && len(quotas) == 0 {
		log.Printf("rows affected is zero in GetUserQuoto")
		return nil, errors.New("查不到数据")
	}
	err := dbMessage.Error
	if err != nil {
		log.Printf("error from quota get: %+v", err)
		return nil, err
	}

	return quotas, nil
}
