package model

import "github.com/shopspring/decimal"

type GoodsGetPageReq struct {
	ID        uint32          `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:ID" json:"id"`          // ID
	GoodsNum  *string         `form:"goods_num" gorm:"column:goods_num;type:varchar(128);comment:商品编号" json:"goods_num"` // 商品编号
	GoodsType *string         `json:"goods_type" form:"goods_type" `                                                     // 商品类型
	GoodsName *string         `gorm:"column:goods_name;type:varchar(128);comment:商品名字" json:"goods_name"`                // 商品名字
	Price     decimal.Decimal `gorm:"column:price;type:float;not null;comment:商品单价" json:"price"`                        // 商品单价
	PicURL    *string         `gorm:"column:pic_url;type:varchar(128);comment:商品图片" json:"pic_url"`                      // 商品图片
	Seller    uint32          `gorm:"column:seller;type:bigint;not null;comment:卖家ID" json:"seller"`                     // 卖家ID
	PageSize  uint32          `json:"page_size" form:"page_size" `
	PageNum   uint32          `json:"page_num" form:"page_num"`
}

type GoodsGetReq struct {
	ID        uint32          `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:ID" json:"id"`                             // ID
	GoodsNum  *string         `form:"goods_num" binding:"required" gorm:"column:goods_num;type:varchar(128);comment:商品编号" json:"goods_num"` // 商品编号
	GoodsName *string         `gorm:"column:goods_name;type:varchar(128);comment:商品名字" json:"goods_name"`                                   // 商品名字
	Price     decimal.Decimal `gorm:"column:price;type:float;not null;comment:商品单价" json:"price"`                                           // 商品单价
	PicURL    *string         `gorm:"column:pic_url;type:varchar(128);comment:商品图片" json:"pic_url"`                                         // 商品图片
	Seller    uint32          `gorm:"column:seller;type:bigint;not null;comment:卖家ID" json:"seller"`                                        // 卖家ID
}

type GoodsGetUserReq struct {
	GoodsType *string         `json:"goods_type" gorm:"goods_type"`                                       // 商品类型
	GoodsName *string         `gorm:"column:goods_name;type:varchar(128);comment:商品名字" json:"goods_name"` // 商品名字
	GoddsNum  *string         `json:"godds_num"`                                                          // 商品编码
	Price     decimal.Decimal `gorm:"column:price;type:float;not null;comment:商品单价" json:"price"`         // 商品单价
	Seller    uint32          `gorm:"column:seller;type:bigint;not null;comment:卖家ID" json:"seller"`      // 卖家ID
}

type GoodsCreateReq struct {
	GoodsType *string         `json:"goods_type" form:"goods_type" binding:"required" `  // 商品类型
	GoodsName *string         `json:"goods_name" form:"goods_name" binding:"required"  ` // 商品名字
	Price     decimal.Decimal `json:"price"      form:"price"      binding:"required"  ` // 商品单价
	PicURL    *string         `json:"pic_url"    form:"picURL"  `                        // 商品图片
	Seller    uint32          `json:"seller"     form:"seller"     binding:"required"  ` // 卖家ID
	Stock     uint32          `json:"stock"      form:"stock"      binding:"required"`   // 库存
}

type GoodsUpdateReq struct {
	GoodsNum  *string         `form:"goods_num" binding:"required" gorm:"column:goods_num;type:varchar(128);comment:商品编号" json:"goods_num"`    // 商品编号
	GoodsName *string         `form:"goods_name" binding:"required" gorm:"column:goods_name;type:varchar(128);comment:商品名字" json:"goods_name"` // 商品名字
	Price     decimal.Decimal `form:"price" binding:"required" gorm:"column:price;type:float;not null;comment:商品单价" json:"price"`              // 商品单价
	PicURL    *string         `form:"picURL" gorm:"column:pic_url;type:varchar(128);comment:商品图片" json:"pic_url"`                              // 商品图片
	Seller    uint32          `form:"seller" binding:"required" gorm:"column:seller;type:bigint;not null;comment:卖家ID" json:"seller"`          // 卖家ID
}
