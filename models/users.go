package models

import (
	"database/sql/driver"
	"encoding/json"
)

type JSONMap map[string]interface{}

func (m JSONMap) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	return json.Marshal(m)
}

//func (m *JSONMap) Scan(value interface{}) error {
//	if value == nil {
//		*m = make(map[string]interface{})
//		return nil
//	}
//}
