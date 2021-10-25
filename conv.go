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

// ToBoolMapSet converts a slice or array to map[bool]struct{}
func ToBoolMapSet(i interface{}) map[bool]struct{}{
	m, _ := ToBoolMapSetE(i)
	return m
}

// ToIntMapSet converts a slice or array to map[int]struct{}
func ToIntMapSet(i interface{}) map[int]struct{}{
	m, _ := ToIntMapSetE(i)
	return m
}

// ToI8MapSet converts a slice or array to map[int8]struct{}
func ToI8MapSet(i interface{}) map[int8]struct{}{
	m, _ := ToI8MapSetE(i)
	return m
}

// ToI16MapSet converts a slice or array to map[int16]struct{}
func ToI16MapSet(i interface{}) map[int16]struct{}{
	m, _ := ToI16MapSetE(i)
	return m
}

// ToI32MapSet converts a slice or array to map[int32]struct{}
func ToI32MapSet(i interface{}) map[int32]struct{}{
	m, _ := ToI32MapSetE(i)
	return m
}

// ToI64MapSet converts a slice or array to map[int64]struct{}
func ToI64MapSet(i interface{}) map[int64]struct{}{
	m, _ := ToI64MapSetE(i)
	return m
}

// ToUintMapSet converts a slice or array to map[uint]struct{}
func ToUintMapSet(i interface{}) map[uint]struct{}{
	m, _ := ToUintMapSetE(i)
	return m
}

// ToU8MapSet converts a slice or array to map[uint8]struct{}
func ToU8MapSet(i interface{}) map[uint8]struct{}{
	m, _ := ToU8MapSetE(i)
	return m
}

// ToU16MapSet converts a slice or array to map[uint16]struct{}
func ToU16MapSet(i interface{}) map[uint16]struct{}{
	m, _ := ToU16MapSetE(i)
	return m
}

// ToU32MapSet converts a slice or array to map[uint32]struct{}
func ToU32MapSet(i interface{}) map[uint32]struct{}{
	m, _ := ToU32MapSetE(i)
	return m
}

// ToU64MapSet converts a slice or array to map[uint64]struct{}
func ToU64MapSet(i interface{}) map[uint64]struct{}{
	m, _ := ToU64MapSetE(i)
	return m
}

// ToStrMapSet converts a slice or array to map[string]struct{}
func ToStrMapSet(i interface{}) map[string]struct{}{
	m, _ := ToStrMapSetE(i)
	return m
}

// ToBoolMapSetE converts a slice or array to map[bool]struct{} with error
func ToBoolMapSetE(i interface{}) (map[bool]struct{}, error) {
	m, err := ToMapSetE(i)
	if err != nil {
		return nil, err
	}
	mm := make(map[bool]struct{}, len(m))
	for k := range m {
		v, err := cast.ToBoolE(k)
		if err != nil {
			return nil, err
		}
		mm[v] = struct{}{}
	}
	return mm, nil
}

// ToIntMapSetE converts a slice or array to map[int]struct{} with error
func ToIntMapSetE(i interface{}) (map[int]struct{}, error) {
	m, err := ToMapSetE(i)
	if err != nil {
		return nil, err
	}
	mm := make(map[int]struct{}, len(m))
	for k := range m {
		v, err := cast.ToIntE(k)
		if err != nil {
			return nil, err
		}
		mm[v] = struct{}{}
	}
	return mm, nil
}

// ToI8MapSetE converts a slice or array to map[int8]struct{} with error
func ToI8MapSetE(i interface{}) (map[int8]struct{}, error) {
	m, err := ToMapSetE(i)
	if err != nil {
		return nil, err
	}
	mm := make(map[int8]struct{}, len(m))
	for k := range m {
		v, err := cast.ToInt8E(k)
		if err != nil {
			return nil, err
		}
		mm[v] = struct{}{}
	}
	return mm, nil
}

// ToI16MapSetE converts a slice or array to map[int16]struct{} with error
func ToI16MapSetE(i interface{}) (map[int16]struct{}, error) {
	m, err := ToMapSetE(i)
	if err != nil {
		return nil, err
	}
	mm := make(map[int16]struct{}, len(m))
	for k := range m {
		v, err := cast.ToInt16E(k)
		if err != nil {
			return nil, err
		}
		mm[v] = struct{}{}
	}
	return mm, nil
}

// ToI32MapSetE converts a slice or array to map[int32]struct{} with error
func ToI32MapSetE(i interface{}) (map[int32]struct{}, error) {
	m, err := ToMapSetE(i)
	if err != nil {
		return nil, err
	}
	mm := make(map[int32]struct{}, len(m))
	for k := range m {
		v, err := cast.ToInt32E(k)
		if err != nil {
			return nil, err
		}
		mm[v] = struct{}{}
	}
	return mm, nil
}

// ToI64MapSetE converts a slice or array to map[int64]struct{} with error
func ToI64MapSetE(i interface{}) (map[int64]struct{}, error) {
	m, err := ToMapSetE(i)
	if err != nil {
		return nil, err
	}
	mm := make(map[int64]struct{}, len(m))
	for k := range m {
		v, err := cast.ToInt64E(k)
		if err != nil {
			return nil, err
		}
		mm[v] = struct{}{}
	}
	return mm, nil
}

// ToUintMapSetE converts a slice or array to map[uint]struct{} with error
func ToUintMapSetE(i interface{}) (map[uint]struct{}, error) {
	m, err := ToMapSetE(i)
	if err != nil {
		return nil, err
	}
	mm := make(map[uint]struct{}, len(m))
	for k := range m {
		v, err := cast.ToUintE(k)
		if err != nil {
			return nil, err
		}
		mm[v] = struct{}{}
	}
	return mm, nil
}

// ToU8MapSetE converts a slice or array to map[uint8]struct{} with error
func ToU8MapSetE(i interface{}) (map[uint8]struct{}, error) {
	m, err := ToMapSetE(i)
	if err != nil {
		return nil, err
	}
	mm := make(map[uint8]struct{}, len(m))
	for k := range m {
		v, err := cast.ToUint8E(k)
		if err != nil {
			return nil, err
		}
		mm[v] = struct{}{}
	}
	return mm, nil
}

// ToU16MapSetE converts a slice or array to map[uint16]struct{} with error
func ToU16MapSetE(i interface{}) (map[uint16]struct{}, error) {
	m, err := ToMapSetE(i)
	if err != nil {
		return nil, err
	}
	mm := make(map[uint16]struct{}, len(m))
	for k := range m {
		v, err := cast.ToUint16E(k)
		if err != nil {
			return nil, err
		}
		mm[v] = struct{}{}
	}
	return mm, nil
}

// ToU32MapSetE converts a slice or array to map[uint32]struct{} with error
func ToU32MapSetE(i interface{}) (map[uint32]struct{}, error) {
	m, err := ToMapSetE(i)
	if err != nil {
		return nil, err
	}
	mm := make(map[uint32]struct{}, len(m))
	for k := range m {
		v, err := cast.ToUint32E(k)
		if err != nil {
			return nil, err
		}
		mm[v] = struct{}{}
	}
	return mm, nil
}

// ToU64MapSetE converts a slice or array to map[uint64]struct{} with error
func ToU64MapSetE(i interface{}) (map[uint64]struct{}, error) {
	m, err := ToMapSetE(i)
	if err != nil {
		return nil, err
	}
	mm := make(map[uint64]struct{}, len(m))
	for k := range m {
		v, err := cast.ToUint64E(k)
		if err != nil {
			return nil, err
		}
		mm[v] = struct{}{}
	}
	return mm, nil
}

// ToStrMapSetE converts a slice or array to map[string]struct{} with error
func ToStrMapSetE(i interface{}) (map[string]struct{}, error) {
	m, err := ToMapSetE(i)
	if err != nil {
		return nil, err
	}
	mm := make(map[string]struct{}, len(m))
	for k := range m {
		v, err := cast.ToStringE(k)
		if err != nil {
			return nil, err
		}
		mm[v] = struct{}{}
	}
	return mm, nil
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