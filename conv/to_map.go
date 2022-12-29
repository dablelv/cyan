package conv

import (
	"encoding/json"
	"fmt"
	"reflect"
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
func Struct2MapString(obj any) map[string]string {
	// Check param.
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Struct {
		return nil
	}

	t := reflect.TypeOf(obj)
	var m = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).IsExported() {
			m[t.Field(i).Name] = ToAny[string](v.Field(i).Interface())
		}
	}
	return m
}

// ToMapStrStr casts any type to a map[string]string type.
func ToStringMapString(i interface{}) map[string]string {
	v, _ := ToMapStrStrE(i)
	return v
}

// ToMapStrStrE casts an interface to a map[string]string type.
func ToMapStrStrE(i any) (map[string]string, error) {
	var m = map[string]string{}

	switch v := i.(type) {
	case map[string]string:
		return v, nil
	case map[string]interface{}:
		for k, val := range v {
			m[k] = ToAny[string](val)
		}
		return m, nil
	case map[interface{}]string:
		for k, val := range v {
			m[ToAny[string](k)] = val
		}
		return m, nil
	case map[interface{}]interface{}:
		for k, val := range v {
			m[ToAny[string](k)] = ToAny[string](val)
		}
		return m, nil
	case string:
		err := jsonStringToObject(v, &m)
		return m, err
	default:
		return m, fmt.Errorf("unable to cast %#v of type %T to map[string]string", i, i)
	}
}

// jsonStringToObject attempts to unmarshall a string as JSON into
// the object passed as pointer.
func jsonStringToObject(s string, v interface{}) error {
	data := []byte(s)
	return json.Unmarshal(data, v)
}
