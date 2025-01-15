package serviceImpl

import (
	"fmt"
	"log"
	"my_e_commerce/config"
	model2 "my_e_commerce/data/dal/model"
	model "my_e_commerce/data/req"
	"my_e_commerce/utils"
)

type OrderServiceImpl struct{}

func (s *OrderServiceImpl) UpdateOrder(order model.OrderReq) error {
	//TODO implement me
	panic("implement me")
}

func (s *OrderServiceImpl) GetOrder(order model.OrderReq) ([]*model2.Order, error) {
	//TODO implement me
	panic("implement me")
}

func NewOrderServiceImpl() *OrderServiceImpl { return &OrderServiceImpl{} }

func (s *OrderServiceImpl) CreateOrder(order model.OrderReq) (uint32, error) {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("err : %+v", err)
			return
		}
	}()
	db := config.GetDB()
	var orderInsert model2.Order
	fmt.Printf("order %+v", order)
	utils.CopyStruct(&order, &orderInsert)
	fmt.Printf("orderInsert %+v", orderInsert)
	tx := db.Save(&orderInsert)
	return orderInsert.ID, tx.Error
}
