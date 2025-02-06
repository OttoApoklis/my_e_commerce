package enum

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	// 消费品大类
	CONSUMER_GOODS = "CONSUMER_GOODS"
	// 食品饮料
	FOOD_AND_BEVERAGE = "FOOD_AND_BEVERAGE"
	// 服装服饰
	CLOTHING_AND_ACCESSORIES = "FOOD_AND_BEVERAGE"
	// 家具用品
	HOUSEHOLD_ITEM = "FOOD_AND_BEVERAGE"
	// 数码产品
	DIGITAL_PRODUCT = "FOOD_AND_BEVERAGE"
	// 美妆个护
	BEATY_AND_PERSONAL_CARE_PRODUCT = "FOOD_AND_BEVERAGE"
	// 母婴用品
	MATERNITY_AND_BABY_PRODUCT = "FOOD_AND_BEVERAGE"
)

var dictNameToString = map[string]string{
	// 消费品大类
	"消费品大类": CONSUMER_GOODS,
	// 食品饮料
	"食品饮料": FOOD_AND_BEVERAGE,
	// 服装服饰
	"服装服饰": CLOTHING_AND_ACCESSORIES,
	// 家具用品
	"家具用品": HOUSEHOLD_ITEM,
	// 数码产品
	"数码产品": DIGITAL_PRODUCT,
	// 美妆个护
	"美妆个护": BEATY_AND_PERSONAL_CARE_PRODUCT,
	// 母婴用品
	"母婴用品": MATERNITY_AND_BABY_PRODUCT,
}

var dictStringToCode = map[string]int{
	// 消费品大类
	CONSUMER_GOODS: 1000000,
	// 食品饮料
	FOOD_AND_BEVERAGE: 1100000,
	// 服装服饰
	CLOTHING_AND_ACCESSORIES: 1200000,
	// 家具用品
	HOUSEHOLD_ITEM: 1300000,
	// 数码产品
	DIGITAL_PRODUCT: 1400000,
	// 美妆个护
	BEATY_AND_PERSONAL_CARE_PRODUCT: 1500000,
	// 母婴用品
	MATERNITY_AND_BABY_PRODUCT: 1600000,
}

func GetNumPrefix(code string) (string, error) {
	if msg, ok := dictStringToCode[code]; ok {
		return strconv.Itoa(msg), nil
	}
	return "", errors.New(fmt.Sprintf("unknown error code %d", code))
}

// 用于前端进行映射转换
func GetGoodsKid() map[string]string {
	return dictNameToString
}
