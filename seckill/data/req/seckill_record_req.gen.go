// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model


const TableNameSeckillRecord = "seckill_record"

// SeckillRecord 秒杀记录表
type SeckillRecordReq struct {
	ID         uint32    `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:ID" json:"id"`                            // ID
	UserID     uint32    `gorm:"column:user_id;type:bigint;not null;comment:用户ID" json:"user_id"`                                     // 用户ID
	GoodsID    uint32    `gorm:"column:goods_id;type:bigint;not null;comment:商品ID" json:"goods_id"`                                   // 商品ID
	SecNum     *string   `gorm:"column:sec_num;type:varchar(128);comment:秒杀号" json:"sec_num"`                                         // 秒杀号
	OrderNum   *string   `gorm:"column:order_num;type:varchar(128);comment:订单号" json:"order_num"`                                     // 订单号
	Price      float32    `gorm:"column:price;type:int;not null;comment:金额" json:"price"`                                              // 金额
	Status     uint32    `gorm:"column:status;type:int;not null;comment:状态" json:"status"`                                            // 状态
}

