package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// 后段接口接收JSON后反序列化后的结构体
type JSONMap map[string]interface{}

/*
功能是将JSONMap转化为数据库可存储的值
实现了数据库接口——driver.Value GO语言数据类型->数据库数据
具体实现：如果值是[]byte类型，直接返回该字符切片
如果值为nil，则返回nil
*/
func (m JSONMap) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	return json.Marshal(m)
}

/*
功能是将数据库的值转换为JSONMap
实现了sql.Scanner接口 数据库数据->GO语言数据类型
将数据库值扫描并反序列化为JSONMap
支持字符切片[]byte和字符串string类型的值
如果传入的值类型不支持，会返回错误
*/
func (m *JSONMap) Scan(value interface{}) error {
	if value == nil {
		*m = make(map[string]interface{})
		return nil
	}
	var err error
	switch value.(type) {
	case []byte:
		err = json.Unmarshal(value.([]byte), m)
	case string:
		err = json.Unmarshal([]byte(value.(string)), m)
	default:
		err = errors.New("basetypes.JSONMap.Scan: invalid value type")
	}
	if err != nil {
		return err
	}
	return nil
}

type TreeNode[T any] interface {
	GetChilden() []T
	SetChilden(children T)
	GetID() int
	GetParentID() int
}
