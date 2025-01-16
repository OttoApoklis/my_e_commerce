package serviceImpl

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"my_e_commerce/config"
	model2 "my_e_commerce/data/dal/model"
	model "my_e_commerce/data/req"
	"my_e_commerce/utils"
)

type OrderServiceImpl struct{}

func (s *OrderServiceImpl) UpdateOrder(req model.OrderReq) error {
	// TODO 更新做的更全面一些
	db := config.GetDB()
	fmt.Printf("orderreq: %+v\n", *req.OrderNum)
	dbMessage := db.Table("order").
		Where("order_num = ?", *req.OrderNum).
		Update("price", req.Price)
	fmt.Printf("affect row: %+v\n", dbMessage.RowsAffected)
	if dbMessage.RowsAffected == 0 {
		return errors.New("查不到该数据")
	}
	err := dbMessage.Error
	if err != nil {
		log.Printf("error from Order update: %+v", err)
		return err
	}
	return nil
}

func (s *OrderServiceImpl) GetOrder(order model.OrderReq) ([]*model2.Order, error) {
	db := config.GetDB()
	orders := []*model2.Order{}
	err := db.Where("order_num = ?", order.OrderNum).Find(&orders).Error
	if err != nil {
		log.Printf(" err: %+v", err)
		return nil, err
	}
	if orders == nil {
		log.Printf("Get Order is nil")
	}
	return orders, nil
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
	fmt.Println(orderInsert)
	tx := db.Save(&orderInsert)
	fmt.Printf("affect :%+v", tx.RowsAffected)
	return orderInsert.ID, tx.Error
}

func (s *OrderServiceImpl) DeleteOrderById(db *gorm.DB, ID uint32) error {
	dbMessage := db.Where("id = ?", ID).Delete(&model2.Order{})
	if dbMessage.RowsAffected == 0 {
		log.Printf("rows affected is zero in DeleteOrderById")
		return errors.New("查不到数据")
	}
	err := dbMessage.Error
	if err != nil {
		log.Printf("error from Orders delete: %+v", err)
	}
	return nil
}
