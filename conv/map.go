package conv

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// StructToMap converts struct to map[string]any.
// Such as struct{I int, S string}{I: 1, S: "a"} to map["I":1 "S":"a"].
// Note that unexported fields of struct can't be converted.
func StructToMap(a any) map[string]any {
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.Struct {
		return nil
	}

	t := reflect.TypeOf(a)
	var m = make(map[string]any)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).IsExported() {
			m[t.Field(i).Name] = v.Field(i).Interface()
		}
	}
	return m
}

// StructToMapStr converts struct to map[string]string.
// Such as struct{I int, S string}{I: 1, S: "a"} to map["I":"1" "S":"a"].
// Note that unexported fields of struct can't be converted.
func StructToMapStr(a any) map[string]string {
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.Struct {
		return nil
	}

	t := reflect.TypeOf(a)
	var m = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).IsExported() {
			m[t.Field(i).Name] = ToAny[string](v.Field(i).Interface())
		}
	}
	return m
}

// ToMapStr converts any type to a map[string]string type.
func ToMapStr(a any) map[string]string {
	v, _ := ToMapStrE(a)
	return v
}

// ToMapStrE converts any type to a map[string]string type.
func ToMapStrE(a any) (map[string]string, error) {
	var m = map[string]string{}

	switch v := a.(type) {
	case map[string]string:
		return v, nil
	case map[string]any:
		for k, val := range v {
			val, err := ToAnyE[string](val)
			if err != nil {
				return nil, err
			}
			m[k] = val
		}
	case map[any]string:
		for k, val := range v {
			k, err := ToAnyE[string](k)
			if err != nil {
				return nil, err
			}
			m[k] = val
		}
	case map[any]any:
		for k, val := range v {
			k, err := ToAnyE[string](k)
			if err != nil {
				return nil, err
			}
			val, err := ToAnyE[string](val)
			if err != nil {
				return nil, err
			}
			m[k] = val
		}
	case string:
		if err := jsonStringToObject(v, &m); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unable to convert %#v of type %T to map[string]string", a, a)
	}
	return m, nil
}

// jsonStringToObject attempts to unmarshall a string as JSON into
// the object passed as pointer.
func jsonStringToObject(s string, v any) error {
	data := []byte(s)
	return json.Unmarshal(data, v)
}
