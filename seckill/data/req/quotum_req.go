package model

type QuotumReq struct {
	ID      uint32  `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	GoodsID *uint32 `gorm:"column:goods_id;type:bigint;comment:商品ID" json:"goods_id"`                 // 商品ID
	Num     uint32  `gorm:"column:num;type:int;not null;comment:限额" json:"num"`                       // 限额
}
