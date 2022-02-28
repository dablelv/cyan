package util

import (
	"fmt"
	"reflect"

	"github.com/spf13/cast"
)

//
// Part 1: convert a slice or array to the specified type map set strictly.
// Note that the the element type of slice or array need to be equal to map key type.
// For example, []uint64{1, 2, 3} can be converted to map[uint64]struct{}{1:struct{}, 2:struct{},
// 3:struct{}} by calling ToUint64MapSetStrict() but can't be converted to map[string]struct{}{"1":struct{},
// "2":struct{}, "3":struct{}}.
//

// ToBoolMapSetStrict converts a slice or array to map[bool]struct{} strictly.
func ToBoolMapSetStrict(i interface{}) map[bool]struct{} {
	m, _ := ToBoolMapSetStrictE(i)
	return m
}

// ToBoolMapSetStrictE converts a slice or array to map[bool]struct{} with error.
func ToBoolMapSetStrictE(i interface{}) (map[bool]struct{}, error) {
	m, err := ToMapSetStrictE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[bool]struct{}); ok {
		return v, nil
	}
	return nil, fmt.Errorf("convert success but the type %T of result %#v of isn't map[bool]struct{}", m, m)
}

// ToIntMapSetStrict converts a slice or array to map[int]struct{}.
func ToIntMapSetStrict(i interface{}) map[int]struct{} {
	m, _ := ToIntMapSetStrictE(i)
	return m
}

// ToIntMapSetStrictE converts a slice or array to map[int]struct{} with error.
func ToIntMapSetStrictE(i interface{}) (map[int]struct{}, error) {
	m, err := ToMapSetStrictE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[int]struct{}); ok {
		return v, nil
	}
	return nil, fmt.Errorf("convert success but the type %T of result %#v isn't map[int]struct{}", m, m)
}

// ToInt8MapSetStrict converts a slice or array to map[int8]struct{}.
func ToInt8MapSetStrict(i interface{}) map[int8]struct{} {
	m, _ := ToInt8MapSetStrictE(i)
	return m
}

// ToInt8MapSetStrictE converts a slice or array to map[int8]struct{} with error.
func ToInt8MapSetStrictE(i interface{}) (map[int8]struct{}, error) {
	m, err := ToMapSetStrictE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[int8]struct{}); ok {
		return v, nil
	}
	return nil, fmt.Errorf("convert success but the type %T of result %#v of isn't map[int8]struct{}", m, m)
}

// ToInt16MapSetStrict converts a slice or array to map[int16]struct{}.
func ToInt16MapSetStrict(i interface{}) map[int16]struct{} {
	m, _ := ToInt16MapSetStrictE(i)
	return m
}

// ToInt16MapSetStrictE converts a slice or array to map[int16]struct{} with error.
func ToInt16MapSetStrictE(i interface{}) (map[int16]struct{}, error) {
	m, err := ToMapSetStrictE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[int16]struct{}); ok {
		return v, nil
	}
	return nil, fmt.Errorf("convert success but the type %T of result %#v of isn't map[int16]struct{}", m, m)
}

// ToInt32MapSetStrict converts a slice or array to map[int32]struct{}.
func ToInt32MapSetStrict(i interface{}) map[int32]struct{} {
	m, _ := ToInt32MapSetStrictE(i)
	return m
}

// ToInt32MapSetStrictE converts a slice or array to map[int32]struct{} with error.
func ToInt32MapSetStrictE(i interface{}) (map[int32]struct{}, error) {
	m, err := ToMapSetStrictE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[int32]struct{}); ok {
		return v, nil
	}
	return nil, fmt.Errorf("convert success but the type %T of result %#v of isn't map[int32]struct{}", m, m)
}

// ToInt64MapSetStrict converts a slice or array to map[int64]struct{}.
func ToInt64MapSetStrict(i interface{}) map[int64]struct{} {
	m, _ := ToInt64MapSetStrictE(i)
	return m
}

// ToInt64MapSetStrictE converts a slice or array to map[int64]struct{} with error.
func ToInt64MapSetStrictE(i interface{}) (map[int64]struct{}, error) {
	m, err := ToMapSetStrictE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[int64]struct{}); ok {
		return v, nil
	}
	return nil, fmt.Errorf("convert success but the type %T of result %#v of isn't map[int64]struct{}", m, m)
}

// ToUintMapSetStrict converts a slice or array to map[uint]struct{}.
func ToUintMapSetStrict(i interface{}) map[uint]struct{} {
	m, _ := ToUintMapSetStrictE(i)
	return m
}

// ToUintMapSetStrictE converts a slice or array to map[uint]struct{} with error.
func ToUintMapSetStrictE(i interface{}) (map[uint]struct{}, error) {
	m, err := ToMapSetStrictE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[uint]struct{}); ok {
		return v, nil
	}
	return nil, fmt.Errorf("convert success but the type %T of result %#v of isn't map[uint8]struct{}", m, m)
}

// ToUint8MapSetStrict converts a slice or array to map[uint8]struct{}.
func ToUint8MapSetStrict(i interface{}) map[uint8]struct{} {
	m, _ := ToUint8MapSetStrictE(i)
	return m
}

// ToUint8MapSetStrictE converts a slice or array to map[uint8]struct{} with error.
func ToUint8MapSetStrictE(i interface{}) (map[uint8]struct{}, error) {
	m, err := ToMapSetStrictE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[uint8]struct{}); ok {
		return v, nil
	}
	return nil, fmt.Errorf("convert success but the type %T of result %#v of isn't map[uint8]struct{}", m, m)
}

// ToUint16MapSetStrict converts a slice or array to map[uint16]struct{}.
func ToUint16MapSetStrict(i interface{}) map[uint16]struct{} {
	m, _ := ToUint16MapSetStrictE(i)
	return m
}

// ToUint16MapSetStrictE converts a slice or array to map[uint16]struct{} with error.
func ToUint16MapSetStrictE(i interface{}) (map[uint16]struct{}, error) {
	m, err := ToMapSetStrictE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[uint16]struct{}); ok {
		return v, nil
	}
	return nil, fmt.Errorf("convert success but the type %T of result %#v of isn't map[uint16]struct{}", m, m)
}

// ToUint32MapSetStrict converts a slice or array to map[uint32]struct{}.
func ToUint32MapSetStrict(i interface{}) map[uint32]struct{} {
	m, _ := ToUint32MapSetStrictE(i)
	return m
}

// ToUint32MapSetStrictE converts a slice or array to map[uint32]struct{} with error.
func ToUint32MapSetStrictE(i interface{}) (map[uint32]struct{}, error) {
	m, err := ToMapSetStrictE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[uint32]struct{}); ok {
		return v, nil
	}
	return nil, fmt.Errorf("convert success but the type %T of result %#v of isn't map[uint32]struct{}", m, m)
}

// ToUint64MapSetStrict converts a slice or array to map[uint64]struct{}.
func ToUint64MapSetStrict(i interface{}) map[uint64]struct{} {
	m, _ := ToUint64MapSetStrictE(i)
	return m
}

// ToUint64MapSetStrictE converts a slice or array to map[uint64]struct{} with error.
func ToUint64MapSetStrictE(i interface{}) (map[uint64]struct{}, error) {
	m, err := ToMapSetStrictE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[uint64]struct{}); ok {
		return v, nil
	}
	return nil, fmt.Errorf("convert success but the type %T of result %#v of isn't map[uint64]struct{}", m, m)
}

// ToStrMapSetStrict converts a slice or array to map[string]struct{}.
func ToStrMapSetStrict(i interface{}) map[string]struct{} {
	m, _ := ToStrMapSetStrictE(i)
	return m
}

// ToStrMapSetStrictE converts a slice or array to map[string]struct{} with error.
func ToStrMapSetStrictE(i interface{}) (map[string]struct{}, error) {
	m, err := ToMapSetStrictE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[string]struct{}); ok {
		return v, nil
	}
	return nil, fmt.Errorf("convert success but the type %T of result %#v of isn't map[string]struct{}", m, m)
}

// ToMapSetStrictE converts a slice or array to map set with error strictly.
// The result of map key type is equal to the type input element.
func ToMapSetStrictE(i interface{}) (interface{}, error) {
	// check params.
	if i == nil {
		return nil, fmt.Errorf("unable to converts nil to map[interface{}]struct{}")
	}
	t := reflect.TypeOf(i)
	kind := t.Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return nil, fmt.Errorf("the type %T of input %#v isn't a slice or array", i, i)
	}
	// execute the convert.
	v := reflect.ValueOf(i)
	mT := reflect.MapOf(t.Elem(), reflect.TypeOf(struct{}{}))
	mV := reflect.MakeMapWithSize(mT, v.Len())
	for j := 0; j < v.Len(); j++ {
		mV.SetMapIndex(v.Index(j), reflect.ValueOf(struct{}{}))
	}
	return mV.Interface(), nil
}

//
// Part 2: convert a any element type of slice or array to to specified type map set.
// Note that the the element type of input don't need to be equal to the map key type.
// For example, []uint64{1, 2, 3} can be converted to map[uint64]struct{}{1:struct{}, 2:struct{},
// 3:struct{}} and also can be converted to map[string]struct{}{"1":struct{}, "2":struct{}, "3":struct{}}.
//

// ToBoolMapSet converts a slice or array to map[bool]struct{}.
func ToBoolMapSet(i interface{}) map[bool]struct{} {
	m, _ := ToBoolMapSetE(i)
	return m
}

// ToBoolMapSetE converts a slice or array to map[bool]struct{} with error.
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

// ToIntMapSet converts a slice or array to map[int]struct{}.
func ToIntMapSet(i interface{}) map[int]struct{} {
	m, _ := ToIntMapSetE(i)
	return m
}

// ToIntMapSetE converts a slice or array to map[int]struct{} with error.
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

// ToInt8MapSet converts a slice or array to map[int8]struct{}.
func ToInt8MapSet(i interface{}) map[int8]struct{} {
	m, _ := ToInt8MapSetE(i)
	return m
}

// ToInt8MapSetE converts a slice or array to map[int8]struct{} with error.
func ToInt8MapSetE(i interface{}) (map[int8]struct{}, error) {
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

// ToInt16MapSet converts a slice or array to map[int16]struct{}.
func ToInt16MapSet(i interface{}) map[int16]struct{} {
	m, _ := ToInt16MapSetE(i)
	return m
}

// ToInt16MapSetE converts a slice or array to map[int16]struct{} with error.
func ToInt16MapSetE(i interface{}) (map[int16]struct{}, error) {
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

// ToInt32MapSet converts a slice or array to map[int32]struct{}.
func ToInt32MapSet(i interface{}) map[int32]struct{} {
	m, _ := ToInt32MapSetE(i)
	return m
}

// ToInt32MapSetE converts a slice or array to map[int32]struct{} with error.
func ToInt32MapSetE(i interface{}) (map[int32]struct{}, error) {
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

// ToInt64MapSet converts a slice or array to map[int64]struct{}.
func ToInt64MapSet(i interface{}) map[int64]struct{} {
	m, _ := ToInt64MapSetE(i)
	return m
}

// ToInt64MapSetE converts a slice or array to map[int64]struct{} with error.
func ToInt64MapSetE(i interface{}) (map[int64]struct{}, error) {
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

// ToUintMapSet converts a slice or array to map[uint]struct{}.
func ToUintMapSet(i interface{}) map[uint]struct{} {
	m, _ := ToUintMapSetE(i)
	return m
}

// ToUintMapSetE converts a slice or array to map[uint]struct{} with error.
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

// ToUint8MapSet converts a slice or array to map[uint8]struct{}.
func ToUint8MapSet(i interface{}) map[uint8]struct{} {
	m, _ := ToUint8MapSetE(i)
	return m
}

// ToUint8MapSetE converts a slice or array to map[uint8]struct{} with error.
func ToUint8MapSetE(i interface{}) (map[uint8]struct{}, error) {
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

// ToUint16MapSet converts a slice or array to map[uint16]struct{}.
func ToUint16MapSet(i interface{}) map[uint16]struct{} {
	m, _ := ToUint16MapSetE(i)
	return m
}

// ToUint16MapSetE converts a slice or array to map[uint16]struct{} with error.
func ToUint16MapSetE(i interface{}) (map[uint16]struct{}, error) {
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

// ToUint32MapSet converts a slice or array to map[uint32]struct{}.
func ToUint32MapSet(i interface{}) map[uint32]struct{} {
	m, _ := ToUint32MapSetE(i)
	return m
}

// ToUint32MapSetE converts a slice or array to map[uint32]struct{} with error.
func ToUint32MapSetE(i interface{}) (map[uint32]struct{}, error) {
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

// ToUint64MapSet converts a slice or array to map[uint64]struct{}.
func ToUint64MapSet(i interface{}) map[uint64]struct{} {
	m, _ := ToUint64MapSetE(i)
	return m
}

// ToUint64MapSetE converts a slice or array to map[uint64]struct{} with error.
func ToUint64MapSetE(i interface{}) (map[uint64]struct{}, error) {
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

// ToStrMapSet converts a slice or array to map[string]struct{}.
func ToStrMapSet(i interface{}) map[string]struct{} {
	m, _ := ToStrMapSetE(i)
	return m
}

// ToStrMapSetE converts a slice or array to map[string]struct{} with error.
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

// ToMapSetE converts a slice or array to map[interface{}]struct{} with error.
func ToMapSetE(i interface{}) (map[interface{}]struct{}, error) {
	// check param
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

//
// Part 3: convert a string to map set after split.
// For example, string "a,b,c" split by comma to map set is map[string]struct{}{"a":struct{}{}, "b":struct{},
// "c":struct{}}.
//

// SplitStrToMapSet convert a string to map set after split
func SplitStrToMapSet(v string, sep string) map[string]struct{} {
	m := make(map[string]struct{})
	for _, v := range Split(v, sep) {
		m[v] = struct{}{}
	}
	return m
}