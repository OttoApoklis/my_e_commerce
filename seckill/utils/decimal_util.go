package utils

import (
	"github.com/shopspring/decimal"
)

// MultiplyFloatsAsDecimals 接受两个 float64 类型的参数，
// 将它们转换为 Decimal 类型并进行乘法运算，最后返回结果。
func MultiplyFloatsAsDecimals(a, b float64) decimal.Decimal {
	// 将 float64 转换为 Decimal
	decA := decimal.NewFromFloat(a)

	decB := decimal.NewFromFloat(b)
	// 进行乘法运算
	result := decA.Mul(decB)

	return result
}
