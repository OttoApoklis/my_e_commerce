package utils

import (
	"fmt"
	"reflect"
)

func SetField(obj interface{}, name string, value interface{}) error {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("expecting a pointer to a struct")
	}
	v = v.Elem()
	field := v.FieldByName(name)
	if !field.IsValid() {
		return fmt.Errorf("field not found: %s", name)
	}
	if !field.CanSet() {
		return fmt.Errorf("cannot set field: %s", name)
	}

	val := reflect.ValueOf(value)
	if field.Type() != val.Type() {
		return fmt.Errorf("provided value type didn't match target field type")
	}

	field.Set(val)
	return nil
}

func CopyStruct(src, dst interface{}) error {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst)

	// 确保src和dst都是指针
	if srcVal.Kind() != reflect.Ptr || dstVal.Kind() != reflect.Ptr {
		return fmt.Errorf("both src and dst must be pointers to structs")
	}

	// 获取实际的结构体值
	srcElem := srcVal.Elem()
	dstElem := dstVal.Elem()

	// 确保src和dst都是结构体
	if srcElem.Kind() != reflect.Struct || dstElem.Kind() != reflect.Struct {
		return fmt.Errorf("both src and dst must be pointers to structs")
	}

	dstNames := make(map[string]bool)
	for i := 0; i < dstElem.NumField(); i++ {
		dstNames[dstElem.Type().Field(i).Name] = true
	}

	for i := 0; i < srcElem.NumField(); i++ {
		fieldName := srcElem.Type().Field(i).Name
		if _, exists := dstNames[fieldName]; exists {
			srcField := srcElem.Field(i)
			if !srcField.CanInterface() {
				continue
			}
			err := SetField(dst, fieldName, srcField.Interface())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

type OrderReq struct {
	ID      int
	Details string
}

func TestCopy() {
	srcOrder := &OrderReq{
		ID:      10,
		Details: "Sample order",
	}

	dstOrder := &OrderReq{}

	err := CopyStruct(srcOrder, dstOrder)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("Source Order: %+v\n", srcOrder)
	fmt.Printf("Destination Order: %+v\n", dstOrder)
}
