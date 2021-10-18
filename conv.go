package util

import (
	"fmt"
	"reflect"

	"github.com/spf13/cast"
)

// Map2SliceE converts keys and values of map to slice with error
func Map2SliceE(i interface{})([]interface{}, []interface{}, error) {
	kind := reflect.TypeOf(i).Kind()
	if kind != reflect.Map {
		return nil, nil, fmt.Errorf("the input %#v of type %T isn't a map", i, i)
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
func Map2StrSliceE(i interface{})(k []string, v []string, err error) {
	slK, slV, err := Map2SliceE(i)
	if err != nil {
		return
	}
	k, err = cast.ToStringSliceE(slK)
	if err != nil {
		return
	}
	v, err = cast.ToStringSliceE(slV)
	return
}

// Map2StrSlice converts keys and values of map to string slice
func Map2StrSlice(i interface{})([]string, []string) {
	slK, slV, _ := Map2StrSliceE(i)
	return slK, slV
}

// ToMapSetE converts a slice or array to map[interface{}]struct{} with error
func ToMapSetE(i interface{}) (map[interface{}]struct{}, error) {
	// judge the validation of the input
	if i == nil {
		return nil, fmt.Errorf("unable to converts %#v of type %T to map[interface{}]struct{}", i, i)
	}
	kind := reflect.TypeOf(i).Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return nil, fmt.Errorf("the input %#v of type %T isn't a slice or array", i, i)
	}

	// execute the convert
	v := reflect.ValueOf(i)
	m := make(map[interface{}]struct{}, v.Len())
	for j := 0; j < v.Len(); j++ {
		m[v.Index(j).Interface()] = struct{}{}
	}
	return m, nil
}

// ToStrMapSetE converts a slice or array to map[string]struct{} with error
func ToStrMapSetE(i interface{}) (map[string]struct{}, error) {
	m, err := ToMapSetE(i)
	if err != nil {
		return nil, err
	}
	mm := make(map[string]struct{}, len(m))
	for k := range m {
		kk, err := cast.ToStringE(k)
		if err != nil {
			return nil, err
		}
		mm[kk] = struct{}{}
	}
	return mm, nil
}

// ToStrMapSet converts a slice or array to map[string]struct{}
func ToStrMapSet(i interface{}) map[string]struct{}{
	m, _ := ToStrMapSetE(i)
	return m
}
