package serviceImpl

import (
	"context"
	"errors"
	"fmt"
	"log"
	"my_e_commerce/config"
	model2 "my_e_commerce/data/dal/model"
	model "my_e_commerce/data/req"
	model3 "my_e_commerce/data/resp"
	"my_e_commerce/enum"
	redis_test "my_e_commerce/redis_init"
	"my_e_commerce/service"
	"my_e_commerce/utils"
	"strings"
)

type SeckillServiceImpl struct{}

func (s *SeckillServiceImpl) CreateSecRecord(service *service.SeckillRecordService) error {
	//TODO implement me
	panic("implement me")
}

func (s *SeckillServiceImpl) GetSeckillRecord(seckillNum string) ([]*model2.SeckillRecord, error) {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("err: %+v", err)
		}
	}()
	db := config.GetDB()
	var seckillRecords []*model2.SeckillRecord
	err := db.Table("seckill_record").Where("sec_num = ?", seckillNum).Find(&seckillRecords).Error
	if err != nil {
		return seckillRecords, err
	}
	return seckillRecords, nil
}

func NewSeckillServiceImpl() *SeckillServiceImpl { return &SeckillServiceImpl{} }

func (s *SeckillServiceImpl) GetSeckillRecordByID(id uint32) ([]model2.SeckillRecord, error) {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("err: %+v", err)

		}
	}()
	db := config.GetDB()
	seckillRecords := []model2.SeckillRecord{}
	fmt.Println("id: ", id)
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("查询失败")
			fmt.Println("err: ", err)
		}
	}()
	err := db.Select("id", "username", "email", "first_name", "last_Name").Where("id = ?", id).Find(&seckillRecords).Error
	if err != nil {
		fmt.Printf("Error: ", err)
		return nil, err
	}
	fmt.Println("Users:", s)
	return seckillRecords, nil
}

func (s *SeckillServiceImpl) CreateSeckillRecord(seckillRecordReq *model.SeckillRecordReq) error {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("err: %+v", err)
		}
	}()
	// 配置数据库连接字符串，替换为实际的用户名、密码、数据库名等信息
	db := config.GetDB()

	err := db.Table("seckill_record").Save(&seckillRecordReq).Error
	return err
}

func (s *SeckillServiceImpl) UpdateSeckillRecord(seckillRecordReq *model.SeckillRecordReq) error {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("err: %+v", err)
		}
	}()
	db := config.GetDB()
	var userOld model2.SeckillRecord
	err := db.Select("id", "username", "email", "first_name", "last_Name").Where("id = ?", seckillRecordReq.ID).Find(&userOld).Error
	if err != nil {
		return err
	}
	seckillRecordOldReq := model.SeckillRecordReq{}
	utils.CopyStruct(userOld, seckillRecordOldReq)
	updates := utils.CompareAndCollectionChanges(seckillRecordOldReq, *seckillRecordReq)
	err = db.Model(updates).Where("id = ?", seckillRecordReq.ID).Error
	if err != nil {
		log.Printf("user update err : %+v", err)
		return err
	}
	return nil
}

func (s *SeckillServiceImpl) SeckillRecordStatusChange(seckillRecordNum string, status uint32) (bool, error) {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("err: %+v", err)
		}
	}()
	db := config.GetDB()
	res := db.Table("seckill_record").Where("sec_num = ? and status = ?", seckillRecordNum, enum.SECKILL_ORDER_CREATED).Update("status", status)
	if res.Error != nil {
		return false, res.Error
	}
	if res.RowsAffected == 0 {
		return false, errors.New("影响行数为0")
	}
	if res.Error != nil {
		return false, res.Error
	}
	return true, nil
}

func (s *SeckillServiceImpl) DeleteRedisSeckillRecord(seckillRecordNum string) (bool, error) {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("err: %+v", err)
		}
	}()
	rdb, err := config.GetRedisConnection()
	if err != nil {
		return false, err
	}
	ctx := context.Background()
	member := seckillRecordNum
	ok, err := utils.DeleteToSortedSet(ctx, rdb, redis_test.ZSETKEY_ORDER, member)
	if !ok {
		return false, err
	}
	return true, nil
}

func (s *SeckillServiceImpl) GetSeckillRecordByUser(seckillGetRecord model.SeckillRecordGetReq) ([]*model3.SeckillRecordResp, error) {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("err: %+v", err)
		}
	}()
	db := config.GetDB()
	var seckillRecords []*model2.SeckillRecord
	var seckillRecordsResps []*model3.SeckillRecordResp
	if seckillGetRecord.UserID != nil {
		db = db.Where("user_id = ?", *seckillGetRecord.UserID)
	} else {
		return seckillRecordsResps, errors.New("用户id缺失")
	}
	if seckillGetRecord.Status != nil {
		db = db.Where("status = ?", *seckillGetRecord.Status)
	}
	if seckillGetRecord.BeginTime != nil {
		db = db.Where("create_time >= ?", seckillGetRecord.BeginTime)
	}
	if seckillGetRecord.EndTime != nil {
		db = db.Where("end_time <= ?", seckillGetRecord.EndTime)
	}
	if seckillGetRecord.Order != nil {
		orders := seckillGetRecord.Order
		for order := range *orders {
			if !strings.Contains(order, "asc") {
				db = db.Order(fmt.Sprintf("%s desc", order))
			} else {
				db = db.Order(fmt.Sprintf("%s", order))
			}
		}
	}
	if seckillGetRecord.Limit != nil {
		limit := seckillGetRecord.Limit
		db = db.Limit(*limit)
	} else {
		db = db.Limit(1000)
	}
	if err := db.Find(&seckillRecords).Error; err != nil {
		log.Printf("查询秒杀记录失败！")
		return seckillRecordsResps, err
	}
	// 获取商品和订单信息
	for _, seckillRecord := range seckillRecords {
		log.Printf("seckillRecord %+v\n", seckillRecord)
		var seckillRecordResp model3.SeckillRecordResp
		utils.CopyStruct(seckillRecord, &seckillRecordResp)
		log.Printf("seckillRecordResp %+v\n", seckillRecordResp)
		if seckillRecord.OrderNum != nil {
			var order model2.Order
			db1 := config.GetDB()
			err := db1.Table("order").Where("order_num = ?", *seckillRecord.OrderNum).Find(&order).Error
			if err != nil {
				log.Printf("获取秒杀信息中，获取order失败！")
			} else {
				seckillRecordResp.Order = order
			}
		}
		if seckillRecord.GoodsID != 0 {
			var goods model2.Good
			db2 := config.GetDB()
			err := db2.Table("goods").Where("id = ?", seckillRecord.GoodsID).Find(&goods).Error
			if err != nil {
				log.Printf("获取秒杀信息中，获取goods失败！")

			} else {
				seckillRecordResp.Goods = goods
			}
		}
		seckillRecordsResps = append(seckillRecordsResps, &seckillRecordResp)
	}
	return seckillRecordsResps, nil
}

func (s *SeckillServiceImpl) GetSeckillRecordByUserLast(seckillGetRecord model.SeckillRecordGetReq) ([]*model2.SeckillRecord, error) {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("err: %+v", err)
		}
	}()
	db := config.GetDB()
	var seckillRecords []*model2.SeckillRecord
	if seckillGetRecord.UserID != nil {
		db = db.Where("user_id = ?", *seckillGetRecord.UserID)
	} else {
		return seckillRecords, errors.New("用户id缺失")
	}
	if seckillGetRecord.Status != nil {
		db = db.Where("status = ?", *seckillGetRecord.Status)
	}
	if seckillGetRecord.BeginTime != nil {
		db = db.Where("create_time >= ?", seckillGetRecord.BeginTime)
	}
	if seckillGetRecord.EndTime != nil {
		db = db.Where("end_time <= ?", seckillGetRecord.EndTime)
	}
	if seckillGetRecord.Order != nil {
		orders := seckillGetRecord.Order
		for order := range *orders {
			db = db.Order(fmt.Sprintf("%s desc", order))
		}
	}
	if err := db.Limit(1).Find(&seckillRecords).Error; err != nil {
		log.Printf("查询秒杀记录失败！")
		return seckillRecords, err
	}
	return seckillRecords, nil
}

func (s *SeckillServiceImpl) GetSeckillRecordBySecNum(seckillGetRecord model.SeckillRecordBySecNumReq) ([]*model2.SeckillRecord, error) {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("err: %+v", err)
		}
	}()
	db := config.GetDB()
	var seckillRecords []*model2.SeckillRecord
	if seckillGetRecord.SecNum != nil {
		db = db.Select("sec_num").Where("sec_num = ?", *seckillGetRecord.SecNum)
	} else {
		return seckillRecords, errors.New("找不到秒杀单！")
	}
	if err := db.Limit(1).Find(&seckillRecords).Error; err != nil {
		log.Printf("查询秒杀记录失败！")
		return seckillRecords, err
	}
	return seckillRecords, nil
}
