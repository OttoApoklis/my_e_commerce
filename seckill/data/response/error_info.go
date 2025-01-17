package response

import (
	"fmt"
)

var (
	SUCCESS                    = 0
	ERR_DB                     = 10
	ERR_INPUT_INVALID          = 10020
	ERR_SHOULD_BIND            = 10021
	ERR_JSON_MARSHAL           = 10022
	ERR_JSON_BIND              = 10023
	ERR_FIND_GOODS_FAILED      = 10101
	ERR_GOODS_STOCK_NOT_ENOUGH = 10102
	//ERR_CREATE_SECKILL_RECORD_FAILED = 10104

	ERR_RECORD_USER_KILLED_NUM_FAILED = 10105
	ERR_DESC_STOCK_FAILED             = 10106
	ERR_USER_QUOTA_NOT_ENOUGH         = 10108
	ERR_FIND_USER_QUOTA_FAILED        = 10109

	ERR_CREATE_GOODS_FAILED    = 30001
	ERR_GET_GOODS_FAILED       = 30002
	ERR_UPDATE_GOODS_FAILED    = 30003
	ERR_DELETE_GOODS_FAILED    = 30004
	ERR_GET_GOODS_EMPTY_FAILED = 30005
	SUCCESS_CREATE_GOODS       = 30006
	SUCCESS_GET_GOODS          = 30007
	SUCCESS_UPDATE_GOODS       = 30008
	SUCCESS_DELETE_BY_ID_GOODS = 30009

	// order 4
	ERR_CREATE_ORDER_FAILED    = 40001
	ERR_GET_ORDER_FAILED       = 40002
	ERR_UPDATE_ORDER_FAILED    = 40003
	ERR_DELETE_ORDER_FAILED    = 40004
	ERR_GET_ORDER_EMPTY_FAILED = 40005
	SUCCESS_CREATE_ORDER       = 40006
	SUCCESS_GET_ORDER          = 40007
	SUCCESS_UPDATE_ORDER       = 40008
	SUCCESS_DELETE_BY_ID_ORDER = 40009

	// quota 5
	ERR_CREATE_QUOTA_FAILED    = 50001
	ERR_GET_QUOTA_FAILED       = 50002
	ERR_UPDATE_QUOTA_FAILED    = 50003
	ERR_DELETE_QUOTA_FAILED    = 50004
	ERR_GET_QUOTA_EMPTY_FAILED = 50005
	SUCCESS_CREATE_QUOTA       = 50006
	SUCCESS_GET_QUOTA          = 50007
	SUCCESS_UPDATE_QUOTA       = 50008
	SUCCESS_DELETE_BY_ID_QUOTA = 50009

	// seckill_record 6
	ERR_CREATE_SECKILL_RECORD_FAILED    = 60001
	ERR_GET_SECKILL_RECORD_FAILED       = 60002
	ERR_UPDATE_SECKILL_RECORD_FAILED    = 60003
	ERR_DELETE_SECKILL_RECORD_FAILED    = 60004
	ERR_GET_SECKILL_RECORD_EMPTY_FAILED = 60005
	SUCCESS_CREATE_SECKILL_RECORD       = 60006
	SUCCESS_GET_SECKILL_RECORD          = 60007
	SUCCESS_UPDATE_SECKILL_RECORD       = 60008
	SUCCESS_DELETE_BY_ID_SECKILL_RECORD = 60009

	// seckill_stock 7
	ERR_CREATE_SECKILL_STOCK_FAILED    = 70001
	ERR_GET_SECKILL_STOCK_FAILED       = 70002
	ERR_UPDATE_SECKILL_STOCK_FAILED    = 70003
	ERR_DELETE_SECKILL_STOCK_FAILED    = 70004
	ERR_GET_SECKILL_STOCK_EMPTY_FAILED = 70005
	SUCCESS_CREATE_SECKILL_STOCK       = 70006
	SUCCESS_GET_SECKILL_STOCK          = 70007
	SUCCESS_UPDATE_SECKILL_STOCK       = 70008
	SUCCESS_DELETE_BY_ID_SECKILL_STOCK = 70009

	// user_quota 8
	ERR_CREATE_USER_QUOTA_FAILED       = 80001
	ERR_GET_USER_QUOTA_FAILED          = 80002
	ERR_UPDATE_USER_QUOTA_FAILED       = 80003
	ERR_DELETE_BY_ID_USER_QUOTA_FAILED = 80004
	ERR_GET_USER_QUOTA_EMPTY_FAILED    = 80005
	SUCCESS_CREATE_USER_QUOTA          = 80006
	SUCCESS_GET_USER_QUOTA             = 80007
	SUCCESS_UPDATE_USER_QUOTA          = 80008
	SUCCESS_DELETE_BY_ID_USER_QUOTA    = 80009
)

var errMsgDic = map[int]string{
	SUCCESS:                    "success",
	ERR_INPUT_INVALID:          "input invalid",
	ERR_SHOULD_BIND:            "should bind failed",
	ERR_JSON_BIND:              "JSONBind failed",
	ERR_JSON_MARSHAL:           "json marshal failed",
	ERR_DB:                     "db failed",
	ERR_FIND_GOODS_FAILED:      "商品查询失败",
	ERR_GOODS_STOCK_NOT_ENOUGH: "商品库存不足",

	ERR_CREATE_SECKILL_RECORD_FAILED:  "秒杀记录创建失败",
	ERR_RECORD_USER_KILLED_NUM_FAILED: "记录用户已经秒杀到的名额数失败",
	ERR_DESC_STOCK_FAILED:             "库存扣减失败",
	ERR_USER_QUOTA_NOT_ENOUGH:         "用户额度不足",
	ERR_FIND_USER_QUOTA_FAILED:        "查询用户额度失败",

	ERR_CREATE_GOODS_FAILED:    "商品创建失败",
	ERR_GET_GOODS_FAILED:       "商品查询失败",
	ERR_UPDATE_GOODS_FAILED:    "商品更新失败",
	ERR_DELETE_GOODS_FAILED:    "商品删除失败",
	ERR_GET_GOODS_EMPTY_FAILED: "商品获取为空",
	SUCCESS_CREATE_GOODS:       "商品创建成功",
	SUCCESS_GET_GOODS:          "商品查询成功",
	SUCCESS_UPDATE_GOODS:       "商品更新成功",
	SUCCESS_DELETE_BY_ID_GOODS: "商品删除成功",

	ERR_CREATE_ORDER_FAILED:    "订单创建失败",
	ERR_GET_ORDER_FAILED:       "订单获取失败",
	ERR_UPDATE_ORDER_FAILED:    "订单更新失败",
	ERR_DELETE_ORDER_FAILED:    "订单删除失败",
	ERR_GET_ORDER_EMPTY_FAILED: "订单获取为空",
	SUCCESS_CREATE_ORDER:       "订单创建成功",
	SUCCESS_GET_ORDER:          "订单获取成功",
	SUCCESS_UPDATE_ORDER:       "订单更新成功",
	SUCCESS_DELETE_BY_ID_ORDER: "订单删除成功",

	ERR_CREATE_QUOTA_FAILED:    "限额创建失败",
	ERR_GET_QUOTA_FAILED:       "限额获取失败",
	ERR_UPDATE_QUOTA_FAILED:    "限额更新失败",
	ERR_DELETE_QUOTA_FAILED:    "限额删除失败",
	ERR_GET_QUOTA_EMPTY_FAILED: "限额获取为空",
	SUCCESS_CREATE_QUOTA:       "限额创建成功",
	SUCCESS_GET_QUOTA:          "限额获取成功",
	SUCCESS_UPDATE_QUOTA:       "限额更新成功",
	SUCCESS_DELETE_BY_ID_QUOTA: "限额删除成功",

	ERR_CREATE_SECKILL_RECORD_FAILED:    "秒杀记录创建失败",
	ERR_GET_SECKILL_RECORD_FAILED:       "秒杀记录获取失败",
	ERR_UPDATE_SECKILL_RECORD_FAILED:    "秒杀记录更新失败",
	ERR_DELETE_SECKILL_RECORD_FAILED:    "秒杀记录删除失败",
	ERR_GET_SECKILL_RECORD_EMPTY_FAILED: "秒杀记录获取为空",
	SUCCESS_CREATE_SECKILL_RECORD:       "秒杀记录创建成功",
	SUCCESS_GET_SECKILL_RECORD:          "秒杀记录获取成功",
	SUCCESS_UPDATE_SECKILL_RECORD:       "秒杀记录更新成功",
	SUCCESS_DELETE_BY_ID_SECKILL_RECORD: "秒杀记录删除成功",

	ERR_CREATE_SECKILL_STOCK_FAILED:    "秒杀库存创建失败",
	ERR_GET_SECKILL_STOCK_FAILED:       "秒杀库存获取失败",
	ERR_UPDATE_SECKILL_STOCK_FAILED:    "秒杀库存更新失败",
	ERR_DELETE_SECKILL_STOCK_FAILED:    "秒杀库存删除失败",
	ERR_GET_SECKILL_STOCK_EMPTY_FAILED: "秒杀库存获取为空",
	SUCCESS_CREATE_SECKILL_STOCK:       "秒杀库存创建成功",
	SUCCESS_GET_SECKILL_STOCK:          "秒杀库存获取成功",
	SUCCESS_UPDATE_SECKILL_STOCK:       "秒杀库存更新成功",
	SUCCESS_DELETE_BY_ID_SECKILL_STOCK: "秒杀库存删除成功",

	ERR_CREATE_USER_QUOTA_FAILED:       "用户限额创建失败",
	ERR_GET_USER_QUOTA_FAILED:          "用户限额获取失败",
	ERR_UPDATE_USER_QUOTA_FAILED:       "用户限额更新失败",
	ERR_DELETE_BY_ID_USER_QUOTA_FAILED: "用户限额删除失败",
	ERR_GET_USER_QUOTA_EMPTY_FAILED:    "用户限额获取为空",
	SUCCESS_CREATE_USER_QUOTA:          "用户限额创建成功",
	SUCCESS_GET_USER_QUOTA:             "用户限额获取成功",
	SUCCESS_UPDATE_USER_QUOTA:          "用户限额更新成功",
	SUCCESS_DELETE_BY_ID_USER_QUOTA:    "用户限额删除成功",
}

// GetErrMsg 获取错误描述
func GetErrMsg(code int) string {
	if msg, ok := errMsgDic[code]; ok {
		return msg
	}
	return fmt.Sprintf("unknown error code %d", code)
}
