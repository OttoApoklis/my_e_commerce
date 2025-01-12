package dao

import (
	"gorm.io/gorm"
	"log"
	"my_e_commerce/data/model"
)

func InsertUser(db *gorm.DB, user model.User) error {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("err: %+v", err)
		}
	}()

	return nil
}
