package filter

import (
	"github.com/shopspring/decimal"
)

// Good 商品表
type GoodFilter struct {
	ID        uint32          `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	GoodsNum  *string         `gorm:"column:goods_num;type:varchar(128);comment:商品编号" json:"goods_num"`         // 商品编号
	GoodsName *string         `gorm:"column:goods_name;type:varchar(128);comment:商品名字" json:"goods_name"`       // 商品名字
	Price     decimal.Decimal `gorm:"column:price;type:decimal(10,2);not null;comment:商品单价" json:"price"`       // 商品单价
	PicURL    *string         `gorm:"column:pic_url;type:varchar(128);comment:商品图片" json:"pic_url"`             // 商品图片
	Seller    uint32          `gorm:"column:seller;type:bigint;not null;comment:卖家ID" json:"seller"`            // 卖家ID
}
