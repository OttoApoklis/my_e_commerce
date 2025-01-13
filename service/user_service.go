package service

import "my_e_commerce/data/model"

type UserService interface {
	CreateUser(user *model.User) error
	GetUserByID(uint32) ([]model.User, error)
	UpdateUser(userOld *model.User, user *model.User) error
}
