package serviceImpl

import (
	"fmt"
	"log"
	"my_e_commerce/config"
	model2 "my_e_commerce/data/dal/model"
	model "my_e_commerce/data/resp"
	"my_e_commerce/service"
	"my_e_commerce/utils"
)

type SeckillServiceImpl struct{}

func (s *SeckillServiceImpl) CreateSecRecord(service *service.SeckillRecordService) error {
	//TODO implement me
	panic("implement me")
}

func (s *SeckillServiceImpl) GetSeckillRecord(service *service.SeckillRecordService) ([]*service.SeckillRecordService, error) {
	//TODO implement me
	panic("implement me")
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

func (s *SeckillServiceImpl) CreateSeckillRecord(seckRecord *model.SeckillRecordReq) error {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("err: %+v", err)
		}
	}()
	// 配置数据库连接字符串，替换为实际的用户名、密码、数据库名等信息
	db := config.GetDB()
	db.Save(&seckRecord)
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
