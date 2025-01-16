// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"gorm.io/gorm"
	"time"
)

const TableNameSeckillStock = "seckill_stock"

// SeckillStock 秒杀库存表
type SeckillStock struct {
	ID         uint32    `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:ID" json:"id"`                            // ID
	GoodsID    *uint32   `gorm:"column:goods_id;type:bigint;comment:商品ID" json:"goods_id"`                                            // 商品ID
	Stock      uint32    `gorm:"column:stock;type:int;not null;comment:库存大小" json:"stock"`                                            // 库存大小
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	ModifyTime time.Time `gorm:"column:modify_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:修改时间" json:"modify_time"` // 修改时间
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:软删除，标记时间可用于恢复" json:"deleted_at"`                             // 软删除，标记时间可用于恢复
}

// TableName SeckillStock's table name
func (*SeckillStock) TableName() string {
	return TableNameSeckillStock
}
