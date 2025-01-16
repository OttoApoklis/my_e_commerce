package service

import (
	"gorm.io/gorm"
	"my_e_commerce/data/dal/model"
	model2 "my_e_commerce/data/req"
)

type OrderService interface {
	CreateOrder(order model2.OrderReq) (uint32, error)
	UpdateOrder(order model2.OrderReq) error
	GetOrder(order model2.OrderReq) ([]*model.Order, error)
	DeleteOrderById(db *gorm.DB, ID uint32) error
}
