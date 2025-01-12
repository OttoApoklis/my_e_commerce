package service

import "my_e_commerce/data/model"

type UserService interface {
	CreateUser(user *model.User) error
	GetUserByID(id uint32) error
}
