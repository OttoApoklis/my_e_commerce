package serviceImpl

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"my_e_commerce/config"
	model2 "my_e_commerce/data/dal/model"
	model "my_e_commerce/data/req"
	"my_e_commerce/enum"
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

// db := config.GetDB()
// orders := []*model2.Order{}
// condition := map[string]interface{}{}
// if(order.Buyer!=0){
//condition["buyer = "] = order.Buyer
//}
//if order.GoodsNum!=nil{
//condition["goods_num like"] = order.GoodsNum
//}
// base := dao.NewDB(db)
// ctx := context.Background()
// err := base.GetByPage(ctx, "order", orders, condition, "create_time", 10, 1)

//func (s *OrderServiceImpl) GetOrderByUser(order model.OrderReq)([]*model2.Order, error){
//	db:= config.GetDB()
//	orders := []*model2.Order{}
//	base := dao.NewDB(db)
//	ctx := context.Background()
//	condition := map[string]interface{}{}
//	condition["buyer = "] = order.Buyer
//	if order.GoodsNum!=nil{
//		condition["goods_num like"] = order.GoodsNum
//	}
//	err := base.GetByPage(ctx, "order", orders, )
//}

func (s *OrderServiceImpl) GetOrder(order model.OrderReq) ([]*model2.Order, error) {
	db := config.GetDB()
	orders := []*model2.Order{}
	if order.Buyer == 0 {
		return orders, errors.New("买方id为空")
	} else {
		db = db.Where("buyer=?", order.Buyer)
	}
	if order.GoodsAmount != nil {
		db = db.Where("goods_amount=?", order.GoodsAmount)
	}
	if order.BeginTime != nil {
		db = db.Where("create_time>=?", *order.BeginTime)
	}
	if order.EndTime != nil {
		db = db.Where("create_time<=?", *order.EndTime)
	}
	if order.GoodsID != 0 {
		db = db.Where("goods_id = ?", order.GoodsID)
	}
	err := db.Find(&orders).Error
	if err != nil {
		log.Printf(" err: %+v", err)
		return nil, err
	}
	if orders == nil {
		log.Printf("Get Order is nil")
	}
	return orders, nil
}

func (s *OrderServiceImpl) GetOrderBySeller(order model.OrderReq) ([]*model2.Order, error) {
	db := config.GetDB()
	orders := []*model2.Order{}
	if order.Seller == 0 {
		return orders, errors.New("卖方id为空")
	} else {
		db = db.Where("seller=?", order.Seller)
	}
	if order.GoodsAmount != nil {
		db = db.Where("goods_amount=?", order.GoodsAmount)
	}
	if order.BeginTime != nil {
		db = db.Where("create_time>=?", *order.BeginTime)
	}
	if order.EndTime != nil {
		db = db.Where("create_time<=?", *order.EndTime)
	}
	if order.GoodsID != 0 {
		db = db.Where("goods_id = ?", order.GoodsID)
	}
	err := db.Find(&orders).Error
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

func (s *OrderServiceImpl) CreateOrder(db *gorm.DB, order model.OrderReq) (uint32, error) {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("err : %+v", err)
			return
		}
	}()
	var orderInsert model2.Order
	fmt.Printf("order %+v\n", order)
	utils.CopyStruct(&order, &orderInsert)
	fmt.Printf("orderInsert %+v\n", orderInsert)
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

// 修改创建态的订单状态
func (s *OrderServiceImpl) OrderStatusChange(orderNum string, status uint32) (bool, error) {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("err : %+v", err)
			return
		}
	}()
	db := config.GetDB()
	res := db.Table("order").Where("order_num = ? and status = ?", orderNum, enum.SECKILL_ORDER_CREATED).Update("status", status)
	if res.Error != nil {
		return false, res.Error
	}
	if res.RowsAffected == 0 {
		return false, errors.New("影响行数为0")
	}
	return true, nil
}
