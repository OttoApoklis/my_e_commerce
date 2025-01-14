package service

import (
	"fmt"
)

var (
	SUCCESS                           = 0
	ERR_INPUT_INVALID                 = 10020
	ERR_SHOULD_BIND                   = 10021
	ERR_JSON_MARSHAL                  = 10022
	ERR_FIND_GOODS_FAILED             = 10101
	ERR_GOODS_STOCK_NOT_ENOUGH        = 10102
	ERR_CREATE_ORDER_FAILED           = 10103
	ERR_CREATE_SECKILL_RECORD_FAILED  = 10104
	ERR_RECORD_USER_KILLED_NUM_FAILED = 10105
	ERR_DESC_STOCK_FAILED             = 10106
	ERR_CREATER_USER_QUOTA_FAILED     = 10107
	ERR_USER_QUOTA_NOT_ENOUGH         = 10108
	ERR_FIND_USER_QUOTA_FAILED        = 10109
)

var errMsgDic = map[int]string{
	SUCCESS:                           "success",
	ERR_INPUT_INVALID:                 "input invalid",
	ERR_SHOULD_BIND:                   "should bind failed",
	ERR_JSON_MARSHAL:                  "json marshal failed",
	ERR_FIND_GOODS_FAILED:             "商品查询失败",
	ERR_GOODS_STOCK_NOT_ENOUGH:        "商品库存不足",
	ERR_CREATE_ORDER_FAILED:           "订单创建失败",
	ERR_CREATE_SECKILL_RECORD_FAILED:  "秒杀记录创建失败",
	ERR_RECORD_USER_KILLED_NUM_FAILED: "记录用户已经秒杀到的名额数失败",
	ERR_DESC_STOCK_FAILED:             "库存扣减失败",
	ERR_CREATER_USER_QUOTA_FAILED:     "插入用户限额记录失败",
	ERR_USER_QUOTA_NOT_ENOUGH:         "用户额度不足",
	ERR_FIND_USER_QUOTA_FAILED:        "查询用户额度失败",
}

// GetErrMsg 获取错误描述
func GetErrMsg(code int) string {
	if msg, ok := errMsgDic[code]; ok {
		return msg
	}
	return fmt.Sprintf("unknown error code %d", code)
}
