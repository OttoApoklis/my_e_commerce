package model

type SeckillReq struct {
	GoodsID     uint32  `gorm:"column:goods_id;type:bigint;not null;comment:商品ID" json:"goods_id"`
	GoodsAmount uint32  `gorm:"column:goods_amount;type:int;not null;comment:商品数量" json:"goods_amount"`
	SecNum      *string `json:"sec_num"`
}

type SeckillBuyReq struct {
	SeckillNum string `json:"seckill_num" gorm:"seckill_num" binding:"required"`
}

type SeckillCancelReq struct {
	SeckillNum string `json:"seckill_num" gorm:"seckill_num" binding:"required"`
}

type SeckillRecordGetReq struct {
	UserID    *uint32                 `json:"user_id" form:"user_id"`       // 用户ID
	Status    *uint32                 `json:"status"`                       // 状态
	BeginTime *string                 `json:"begin_time" form:"begin_time"` // 开始时间
	EndTime   *string                 `json:"end_time"   form:"end_time"`   // 结束时间
	Order     *map[string]interface{} `json:"order" form:"order"`           // 排序字段
}

type SeckillRecordBySecNumReq struct {
	SecNum *string `json:"sec_num" form:"sec_num"`
}
