// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"github.com/shopspring/decimal"
	"time"
)

const TableNameGood = "goods"

// Good 商品表
type Good struct {
	ID         uint32    `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:ID" json:"id"`                            // ID
	GoodsNum   *string   `gorm:"column:goods_num;type:varchar(128);comment:商品编号" json:"goods_num"`                                    // 商品编号
	GoodsName  *string   `gorm:"column:goods_name;type:varchar(128);comment:商品名字" json:"goods_name"`                                  // 商品名字
	Price      decimal.Decimal   `gorm:"column:price;type:decimal(10,2);not null;comment:商品单价" json:"price"`                                  // 商品单价
	PicURL     *string   `gorm:"column:pic_url;type:varchar(128);comment:商品图片" json:"pic_url"`                                        // 商品图片
	Seller     uint32    `gorm:"column:seller;type:bigint;not null;comment:卖家ID" json:"seller"`                                       // 卖家ID
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	ModifyTime time.Time `gorm:"column:modify_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:修改时间" json:"modify_time"` // 修改时间
}

// TableName Good's table name
func (*Good) TableName() string {
	return TableNameGood
}