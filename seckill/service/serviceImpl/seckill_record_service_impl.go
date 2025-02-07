package serviceImpl

import (
	"context"
	"errors"
	"fmt"
	"log"
	"my_e_commerce/config"
	model2 "my_e_commerce/data/dal/model"
	model "my_e_commerce/data/req"
	"my_e_commerce/enum"
	redis_test "my_e_commerce/redis_init"
	"my_e_commerce/service"
	"my_e_commerce/utils"
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

	db.Table("seckill_record").Save(&seckillRecordReq)
	return nil
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
