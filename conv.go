package util

import (
	"errors"
	"reflect"

	"github.com/spf13/cast"
)

// Map2StrSliceE converts key and value of map to string slice with error
func Map2StrSliceE(i interface{})([]string, []string, error) {
	kind := reflect.TypeOf(i).Kind()
	if kind != reflect.Map {
		return nil, nil, errors.New("the input is not a map")
	}
	m := reflect.ValueOf(i)
	keys := m.MapKeys()
	slK, slV := make([]string, 0, len(keys)), make([]string, 0, len(keys))
	for _, k := range keys {
		// convert the key to string
		sK, err := cast.ToStringE(k.Interface())
		if err != nil {
			return nil, nil, err
		}
		slK = append(slK, sK)

		// convert the value to string
		v := m.MapIndex(k)
		sV, err := cast.ToStringE(v.Interface())
		if err != nil {
			return nil, nil, err
		}
		slV = append(slV, sV)
	}
	return slK, slV, nil
}

// Map2StrSlice converts key and value of map to string slice
func Map2StrSlice(i interface{})([]string, []string) {
	slK, slV, _ := Map2StrSliceE(i)
	return slK, slV
}
