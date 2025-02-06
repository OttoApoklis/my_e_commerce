package model

type SeckillStockReq struct {
	ID      uint32  `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	GoodsID *uint32 `gorm:"column:goods_id;type:bigint;comment:商品ID" json:"goods_id"`                 // 商品ID
	Stock   uint32  `gorm:"column:stock;type:int;not null;comment:库存大小" json:"stock"`                 // 库存大小
}
