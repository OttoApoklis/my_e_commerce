package utils

import (
	"log"
	"reflect"
)

func CompareAndCollectionChanges(oldObj, newObj interface{}) map[string]interface{} {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("err %+v", err)
		}
	}()
	oldValue := reflect.ValueOf(oldObj)
	newValue := reflect.ValueOf(newObj)
	if oldValue.Type() != newValue.Type() {
		panic("oldObj and newObj must have same type!")
	}
	changes := make(map[string]interface{})
	for i := 0; i < oldValue.NumField(); i++ {
		oldField := oldValue.Field(i)
		newField := newValue.Field(i)
		if !reflect.DeepEqual(oldField.Interface(), newField.Interface()) {
			filedName := newValue.Type().Field(i).Name
			changes[filedName] = newField.Interface()
		}
	}
	return changes
}
