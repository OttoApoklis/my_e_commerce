package model

type UserQuotumReq struct {
	ID        uint32  `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	UserID    uint32  `gorm:"column:user_id;type:bigint;not null;comment:用户ID" json:"user_id"`          // 用户ID
	GoodsID   *uint32 `gorm:"column:goods_id;type:bigint;comment:商品ID" json:"goods_id"`                 // 商品ID
	Num       uint32  `gorm:"column:num;type:int;not null;comment:限额" json:"num"`                       // 限额
	KilledNum uint32  `gorm:"column:killed_num;type:int;not null;comment:已经消耗的额度" json:"killed_num"`    // 已经消耗的额度
}
