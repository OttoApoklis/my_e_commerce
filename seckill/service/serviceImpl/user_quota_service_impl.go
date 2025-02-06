package serviceImpl

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"my_e_commerce/data/dal/model"
	"my_e_commerce/data/filter"
	model2 "my_e_commerce/data/req"
)

type UserQuotaServiceImpl struct{}

func NewUserQuotaServiceImpl() *UserQuotaServiceImpl {
	return &UserQuotaServiceImpl{}
}

func (s *UserQuotaServiceImpl) CreateUserQuota(db *gorm.DB, quotum *model.UserQuotum) error {
	var userQuotas []model.UserQuotum

	if err := db.Select("goods_id", "user_id").Find(&userQuotas).Error; err != nil {
		log.Printf("err from userQuota find in createUserQuota: %+v", err)
		return err
	}
	if userQuotas != nil && len(userQuotas) != 0 {
		log.Printf("err from userQuota create because of repeateable")
		return errors.New("repeatable in userQuota create")
	}
	err := db.Save(quotum).Error
	if err != nil {
		log.Printf("error from UserQuota create: %+v", err)
		return err
	}
	return nil
}

func (s *UserQuotaServiceImpl) UpdateUserQuota(db *gorm.DB, req *model2.UserQuotumReq) error {
	var quotum model2.UserQuotumReq
	quotum = *req
	fmt.Println(quotum.GoodsID)
	fmt.Println("\n")
	fmt.Println(quotum.UserID)
	dbMessage := db.Model(&model.UserQuotum{}).
		Where("goods_id = ? and user_id = ?", *quotum.GoodsID, quotum.UserID).
		Limit(1).
		Update("num", quotum.Num)
	if dbMessage.RowsAffected == 0 {
		return errors.New("查不到该数据")
	}
	err := dbMessage.Error
	if err != nil {
		log.Printf("error from UserQuota update: %+v", err)
		return err
	}
	return nil
}

func (s *UserQuotaServiceImpl) DeleteUserQuotaById(db *gorm.DB, ID uint32) error {
	dbMessage := db.Where("id = ?", ID).Delete(&model.UserQuotum{})
	if dbMessage.RowsAffected == 0 {
		log.Printf("rows affected is zero in DeleteUserQuotaById")
		return errors.New("查不到数据")
	}
	err := dbMessage.Error
	if err != nil {
		log.Printf("error from UserQuota delete: %+v", err)
	}
	return nil
}

func (s *UserQuotaServiceImpl) DeleteUserQuotaByCondition(db *gorm.DB, filter *filter.UserQuotumFilter) error {
	err := db.Delete(filter)
	if err != nil {
		log.Printf("error form UserQuot delete: %+v", err)
	}
	return nil
}

func (s *UserQuotaServiceImpl) GetUserQuoto(db *gorm.DB, filter *filter.UserQuotumFilter) ([]model.UserQuotum, error) {

	// 动态添加查询条件
	//query = utils.AddCondition(query, &filter.ID, "id")
	//query = utils.AddCondition(query, &filter.Num, "num")
	//query = utils.AddCondition(query, filter.GoodsID, "goods_id")
	//query = utils.AddCondition(query, &filter.UserID, "user_id")
	//query = utils.AddCondition(query, &filter.KilledNum, "killed_num")

	// 查询结果
	var userQuotas []model.UserQuotum
	fmt.Printf("\nfilter: %+v\n", filter)
	dbMessage := db.Table("user_quota").Where("user_id = ? and goods_id = ?", filter.UserID, *filter.GoodsID).Scan(&userQuotas)
	//db.Raw("select * from user_quota where user_id = ? and goods_id = ?", filter.UserID, *filter.GoodsID).Scan(&userQuotas)
	if dbMessage.RowsAffected == 0 {
		log.Printf("rows affected is zero in GetUserQuoto")
		return nil, errors.New("查不到数据")
	}
	err := dbMessage.Error
	if err != nil {
		log.Printf("error from userQuota get: %+v", err)
		return nil, err
	}
	return userQuotas, nil
}
