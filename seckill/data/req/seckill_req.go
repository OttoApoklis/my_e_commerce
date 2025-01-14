package model

type SeckillReq struct {
	GoodsID     uint32 `gorm:"column:goods_id;type:bigint;not null;comment:商品ID" json:"goods_id"`
	GoodsAmount uint32 `gorm:"column:goods_amount;type:int;not null;comment:商品数量" json:"goods_amount"`
}
