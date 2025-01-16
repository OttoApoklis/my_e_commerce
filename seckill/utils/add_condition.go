package utils

import (
	"fmt"
	"gorm.io/gorm"
)

// 添加单个条件
func AddCondition[T any](query *gorm.DB, condition *T, columnName string) *gorm.DB {
	if condition != nil {
		fmt.Printf("condition: %+v", *condition)
		return query.Where(columnName+" = ?", *condition)
	}
	return query
}
