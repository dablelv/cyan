package conv

import (
	"reflect"

	"github.com/spf13/cast"
)

// Struct2Map converts struct to map[string]interface{}.
// Such as struct{I int, S string}{I: 1, S: "a"} to map[I:1 S:a].
// Note that unexported fields of struct can't be converted.
func Struct2Map(obj interface{}) map[string]interface{} {
	// check params
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Struct {
		return nil
	}

	t := reflect.TypeOf(obj)
	var m = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).IsExported() {
			m[t.Field(i).Name] = v.Field(i).Interface()
		}
	}
	return m
}

// Struct2MapString converts struct to map[string]string.
// Such as struct{I int, S string}{I: 1, S: "a"} to map[I:1 S:a].
// Note that unexported fields of struct can't be converted.
func Struct2MapString(obj interface{}) map[string]string {
	// check params
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Struct {
		return nil
	}

	t := reflect.TypeOf(obj)
	var m = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).IsExported() {
			m[t.Field(i).Name] = cast.ToString(v.Field(i).Interface())
		}
	}
	return m
}
