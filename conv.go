package util

import (
	"errors"
	"reflect"

	"github.com/spf13/cast"
)

// map2SliceE converts keys and values of map to slice with error
func map2SliceE(i interface{})([]interface{}, []interface{}, error) {
	kind := reflect.TypeOf(i).Kind()
	if kind != reflect.Map {
		return nil, nil, errors.New("the input is not a map")
	}
	m := reflect.ValueOf(i)
	keys := m.MapKeys()
	slK, slV := make([]interface{}, 0, len(keys)), make([]interface{}, 0, len(keys))
	for _, k := range keys {
		slK = append(slK, k.Interface())
		v := m.MapIndex(k)
		slV = append(slV, v.Interface())
	}
	return slK, slV, nil
}

// Map2StrSliceE converts keys and values of map to string slice with error
func Map2StrSliceE(i interface{})([]string, []string, error) {
	slK, slV, err := map2SliceE(i)
	if err != nil {
		return nil, nil, err
	}
	k, err := cast.ToStringSliceE(slK)
	if err != nil {
		return nil, nil, err
	}
	v, err := cast.ToStringSliceE(slV)
	if err != nil {
		return nil, nil, err
	}
	return k, v, nil
}

// Map2StrSlice converts keys and values of map to string slice
func Map2StrSlice(i interface{})([]string, []string) {
	slK, slV, _ := Map2StrSliceE(i)
	return slK, slV
}