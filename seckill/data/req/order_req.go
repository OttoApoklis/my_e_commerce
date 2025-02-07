package model

import (
	"github.com/shopspring/decimal"
)

type OrderReq struct {
	ID          uint32          `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	Seller      uint32          `gorm:"column:seller;type:bigint;not null;comment:买方ID" json:"seller"`            // 卖方ID
	Buyer       uint32          `gorm:"column:buyer;type:bigint;not null;comment:卖房ID" json:"buyer"`              // 买方ID
	GoodsID     uint32          `gorm:"column:goods_id;type:bigint;not null;comment:商品ID" json:"goods_id"`        // 商品ID
	GoodsNum    *string         `gorm:"column:goods_num;type:varchar(128);comment:商品编号" json:"goods_num"`         // 商品编号
	OrderNum    *string         `gorm:"column:order_num;type:varchar(128);comment:订单号" json:"order_num"`          // 订单号
	GoodsAmount *uint32         `gorm:"column:goods_amount;type:int;comment:商品数量" json:"goods_amount"`            // 商品数量
	Price       decimal.Decimal `gorm:"column:price;type:int;not null;comment:金额" json:"price"`                   // 金额
	Status      uint32          `gorm:"column:status;type:int;not null;comment:状态" json:"status"`                 // 状态
	BeginTime   *string         `gorm:"column:begin_time" json:"begin_time"`                                      // 查询事件范围起点
	EndTime     *string         `gorm:"column:end_time" json:"end_time"`                                          // 查询事件范围起点
}

type OrderGetReq struct {
	ID          uint32          `json:"id"`           // ID
	Seller      uint32          `json:"seller"`       // 卖方ID
	Buyer       uint32          `json:"buyer"`        // 买方ID
	GoodsID     uint32          `json:"goods_id"`     // 商品ID
	GoodsNum    *string         `json:"goods_num"`    // 商品编号
	OrderNum    *string         `json:"order_num"`    // 订单号
	GoodsAmount *uint32         `json:"goods_amount"` // 商品数量
	Price       decimal.Decimal `json:"price"`        // 金额
	Status      uint32          `json:"status"`       // 状态
}

type OrderUserGetReq struct {
	ID          uint32          `json:"id"`                                    // ID
	Seller      uint32          `json:"seller"`                                // 卖方ID
	Buyer       uint32          `json:"buyer" form:"buyer" binding:"required"` // 买方ID
	GoodsID     uint32          `json:"goods_id"`                              // 商品ID
	GoodsNum    *string         `json:"goods_num"`                             // 商品编号
	OrderNum    *string         `json:"order_num"`                             // 订单号
	GoodsAmount *uint32         `json:"goods_amount"`                          // 商品数量
	Price       decimal.Decimal `json:"price"`                                 // 金额
	Status      uint32          `json:"status"`                                // 状态
	BeginTime   *string         `form:"begin_time" json:"begin_time"`          // 查询事件范围起点
	EndTime     *string         `form:"end_time" json:"end_time"`              // 查询事件范围起点
}

type OrderSellerGetReq struct {
	ID          uint32          `json:"id"`                                      // ID
	Seller      uint32          `json:"seller" form:"seller" binding:"required"` // 卖方ID
	Buyer       uint32          `json:"buyer" form:"buyer"`                      // 买方ID
	GoodsID     uint32          `json:"goods_id"`                                // 商品ID
	GoodsNum    *string         `json:"goods_num"`                               // 商品编号
	OrderNum    *string         `json:"order_num"`                               // 订单号
	GoodsAmount *uint32         `json:"goods_amount"`                            // 商品数量
	Price       decimal.Decimal `json:"price"`                                   // 金额
	Status      uint32          `json:"status"`                                  // 状态
	BeginTime   *string         `form:"begin_time" json:"begin_time"`            // 查询事件范围起点
	EndTime     *string         `form:"end_time" json:"end_time"`                // 查询事件范围起点
}
