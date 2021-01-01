package util

import (
	"reflect"

	"github.com/spf13/cast"
)

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var m = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		m[t.Field(i).Name] = v.Field(i).Interface()
	}
	return m
}

func Struct2MapString(obj interface{}) map[string]string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var m = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		m[t.Field(i).Name] = cast.ToString(v.Field(i).Interface())
	}
	return m
}
