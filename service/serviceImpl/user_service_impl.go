package serviceImpl

import (
	"fmt"
	"log"
	"my_e_commerce/config"
	"my_e_commerce/data/model"
	"my_e_commerce/utils"
)

type UserServiceImpl struct{}

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{}
}

func (s *UserServiceImpl) GetUserByID(id uint32) ([]model.User, error) {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("err: %+v", err)

		}
	}()
	db := config.GetDB()
	users := []model.User{}
	fmt.Println("id: ", id)
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("查询失败")
			fmt.Println("err: ", err)
		}
	}()
	err := db.Select("id", "username", "email", "first_name", "last_Name").Where("id = ?", id).Find(&users).Error
	if err != nil {
		fmt.Printf("Error: ", err)
		return nil, err
	}
	fmt.Println("Users:", users)
	return users, nil
}

func (s *UserServiceImpl) CreateUser(user *model.User) error {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("err: %+v", err)
		}
	}()
	// 配置数据库连接字符串，替换为实际的用户名、密码、数据库名等信息
	db := config.GetDB()
	db.Save(&user)
	return nil
}

func (s *UserServiceImpl) UpdateUser(userOld *model.User, user *model.User) error {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("err: %+v", err)
		}
	}()
	db := config.GetDB()
	// 对比需要更新的字段
	fmt.Println("compare")
	fmt.Println(user.Username)
	fmt.Println(userOld.Username)
	updates := utils.CompareAndCollectionChanges(*userOld, *user)
	fmt.Printf("updates : %+v", updates)

	fmt.Println("database operate")
	err := db.Model(updates).Where("id = ?", user.ID).Error

	if err != nil {
		log.Printf("user update err : %+v", err)
		return err
	}

	return nil
}
