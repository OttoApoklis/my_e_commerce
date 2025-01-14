package model

type OrderReq struct {
	ID          uint32  `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	Seller      uint32  `gorm:"column:seller;type:bigint;not null;comment:买方ID" json:"seller"`            // 买方ID
	Buyer       uint32  `gorm:"column:buyer;type:bigint;not null;comment:卖房ID" json:"buyer"`              // 卖房ID
	GoodsID     uint32  `gorm:"column:goods_id;type:bigint;not null;comment:商品ID" json:"goods_id"`        // 商品ID
	GoodsNum    *string `gorm:"column:goods_num;type:varchar(128);comment:商品编号" json:"goods_num"`         // 商品编号
	OrderNum    *string `gorm:"column:order_num;type:varchar(128);comment:订单号" json:"order_num"`          // 订单号
	GoodsAmount *uint32 `gorm:"column:goods_amount;type:int;comment:商品数量" json:"goods_amount"`            // 商品数量
	Price       float32 `gorm:"column:price;type:int;not null;comment:金额" json:"price"`                   // 金额
	Status      uint32  `gorm:"column:status;type:int;not null;comment:状态" json:"status"`                 // 状态
}
