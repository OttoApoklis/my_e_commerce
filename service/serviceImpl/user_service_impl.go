package serviceImpl

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"my_e_commerce/data/model"
)

type UserServiceImpl struct{}

func (s *UserServiceImpl) CreateUser(user *model.User) error {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("err: %+v", err)
		}
	}()
	// 配置数据库连接字符串，替换为实际的用户名、密码、数据库名等信息
	dsn := "root:@(192.168.128.128:3307)/users?charset=utf8mb4&parseTime=True&loc=Local"
	var db *gorm.DB
	var err error
	// 建立数据库连接
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Save(&user)
	return nil
}
