package conv

import (
	"fmt"
	"reflect"

	"github.com/spf13/cast"
)

//
// Part 1: convert a any element type of slice or array to to specified type map set.
// Note that the the element type of input don't need to be equal to the map key type.
// For example, []uint64{1, 2, 3} can be converted to map[uint64]struct{}{1:struct{}, 2:struct{},
// 3:struct{}} and also can be converted to map[string]struct{}{"1":struct{}, "2":struct{}, "3":struct{}}.
//

// ToBoolSet converts a slice or array to map[bool]struct{}.
func ToBoolSet(i any) map[bool]struct{} {
	m, _ := ToBoolSetE(i)
	return m
}

// ToBoolSetE converts a slice or array to map[bool]struct{} with error.
func ToBoolSetE(i any) (map[bool]struct{}, error) {
	m, err := ToSetE(i)
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

// ToIntSet converts a slice or array to map[int]struct{}.
func ToIntSet(i any) map[int]struct{} {
	m, _ := ToIntSetE(i)
	return m
}

// ToIntSetE converts a slice or array to map[int]struct{Set} with error.
func ToIntSetE(i any) (map[int]struct{}, error) {
	m, err := ToSetE(i)
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

// ToInt8Set converts a slice or array to map[int8]struct{}.
func ToInt8Set(i any) map[int8]struct{} {
	m, _ := ToInt8SetE(i)
	return m
}

// ToInt8SetE converts a slice or array to map[int8]struct{} with error.
func ToInt8SetE(i any) (map[int8]struct{}, error) {
	m, err := ToSetE(i)
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

// ToInt16Set converts a slice or array to map[int16]struct{}.
func ToInt16Set(i any) map[int16]struct{} {
	m, _ := ToInt16SetE(i)
	return m
}

// ToInt16SetE converts a slice or array to map[int16]struct{} with error.
func ToInt16SetE(i any) (map[int16]struct{}, error) {
	m, err := ToSetE(i)
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

// ToInt32Set converts a slice or array to map[int32]struct{}.
func ToInt32Set(i any) map[int32]struct{} {
	m, _ := ToInt32SetE(i)
	return m
}

// ToInt32SetE converts a slice or array to map[int32]struct{} with error.
func ToInt32SetE(i any) (map[int32]struct{}, error) {
	m, err := ToSetE(i)
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

// ToInt64Set converts a slice or array to map[int64]struct{}.
func ToInt64Set(i any) map[int64]struct{} {
	m, _ := ToInt64SetE(i)
	return m
}

// ToInt64SetE converts a slice or array to map[int64]struct{} with error.
func ToInt64SetE(i any) (map[int64]struct{}, error) {
	m, err := ToSetE(i)
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

// ToUintSet converts a slice or array to map[uint]struct{}.
func ToUintSet(i any) map[uint]struct{} {
	m, _ := ToUintSetE(i)
	return m
}

// ToUintSetE converts a slice or array to map[uint]struct{} with error.
func ToUintSetE(i any) (map[uint]struct{}, error) {
	m, err := ToSetE(i)
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

// ToUint8Set converts a slice or array to map[uint8]struct{}.
func ToUint8Set(i any) map[uint8]struct{} {
	m, _ := ToUint8SetE(i)
	return m
}

// ToUint8SetE converts a slice or array to map[uint8]struct{} with error.
func ToUint8SetE(i any) (map[uint8]struct{}, error) {
	m, err := ToSetE(i)
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

// ToUint16Set converts a slice or array to map[uint16]struct{}.
func ToUint16Set(i any) map[uint16]struct{} {
	m, _ := ToUint16SetE(i)
	return m
}

// ToUint16SetE converts a slice or array to map[uint16]struct{} with error.
func ToUint16SetE(i any) (map[uint16]struct{}, error) {
	m, err := ToSetE(i)
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

// ToUint32Set converts a slice or array to map[uint32]struct{}.
func ToUint32Set(i any) map[uint32]struct{} {
	m, _ := ToUint32SetE(i)
	return m
}

// ToUint32SetE converts a slice or array to map[uint32]struct{} with error.
func ToUint32SetE(i any) (map[uint32]struct{}, error) {
	m, err := ToSetE(i)
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

// ToUint64Set converts a slice or array to map[uint64]struct{}.
func ToUint64Set(i any) map[uint64]struct{} {
	m, _ := ToUint64SetE(i)
	return m
}

// ToUint64SetE converts a slice or array to map[uint64]struct{} with error.
func ToUint64SetE(i any) (map[uint64]struct{}, error) {
	m, err := ToSetE(i)
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

// ToStrSet converts a slice or array to map[string]struct{}.
func ToStrSet(i any) map[string]struct{} {
	m, _ := ToStrSetE(i)
	return m
}

// ToStrSetE converts a slice or array to map[string]struct{} with error.
func ToStrSetE(i any) (map[string]struct{}, error) {
	m, err := ToSetE(i)
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

// ToSetE converts a slice or array to map[any]struct{} with error.
func ToSetE(i any) (map[any]struct{}, error) {
	// check param
	if i == nil {
		return nil, fmt.Errorf("the input i is nil")
	}
	kind := reflect.TypeOf(i).Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return nil, fmt.Errorf("the input %#v of type %T isn't a slice or array", i, i)
	}
	// execute the convert
	v := reflect.ValueOf(i)
	m := make(map[any]struct{}, v.Len())
	for j := 0; j < v.Len(); j++ {
		m[v.Index(j).Interface()] = struct{}{}
	}
	return m, nil
}

//
// Part 2: convert a slice or array to the specified type map set strictly.
// Note that the the element type of slice or array need to be equal to map key type.
// For example, []uint64{1, 2, 3} can be converted to
// map[uint64]struct{}{1:struct{}, 2:struct{}, 3:struct{}}
// by calling ToUint64SetStrict() but can't be converted to
// map[string]struct{}{"1":struct{},"2":struct{}, "3":struct{}}.
//

// ToBoolSetStrict converts a slice or array to map[bool]struct{} strictly.
func ToBoolSetStrict(i any) map[bool]struct{} {
	m, _ := ToBoolSetStrictE(i)
	return m
}

// ToBoolSetStrictE converts a slice or array to map[bool]struct{} with error.
func ToBoolSetStrictE(i any) (map[bool]struct{}, error) {
	m, err := ToSetStrictE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[bool]struct{}); ok {
		return v, nil
	}
	return nil, fmt.Errorf("convert success but the type %T of result %#v of isn't map[bool]struct{}", m, m)
}

// ToIntSetStrict converts a slice or array to map[int]struct{}.
func ToIntSetStrict(i any) map[int]struct{} {
	m, _ := ToIntSetStrictE(i)
	return m
}

// ToIntSetStrictE converts a slice or array to map[int]struct{} with error.
func ToIntSetStrictE(i any) (map[int]struct{}, error) {
	m, err := ToSetStrictE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[int]struct{}); ok {
		return v, nil
	}
	return nil, fmt.Errorf("convert success but the type %T of result %#v isn't map[int]struct{}", m, m)
}

// ToInt8SetStrict converts a slice or array to map[int8]struct{}.
func ToInt8SetStrict(i any) map[int8]struct{} {
	m, _ := ToInt8SetStrictE(i)
	return m
}

// ToInt8SetStrictE converts a slice or array to map[int8]struct{} with error.
func ToInt8SetStrictE(i any) (map[int8]struct{}, error) {
	m, err := ToSetStrictE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[int8]struct{}); ok {
		return v, nil
	}
	return nil, fmt.Errorf("convert success but the type %T of result %#v of isn't map[int8]struct{}", m, m)
}

// ToInt16SetStrict converts a slice or array to map[int16]struct{}.
func ToInt16SetStrict(i any) map[int16]struct{} {
	m, _ := ToInt16SetStrictE(i)
	return m
}

// ToInt16SetStrictE converts a slice or array to map[int16]struct{} with error.
func ToInt16SetStrictE(i any) (map[int16]struct{}, error) {
	m, err := ToSetStrictE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[int16]struct{}); ok {
		return v, nil
	}
	return nil, fmt.Errorf("convert success but the type %T of result %#v of isn't map[int16]struct{}", m, m)
}

// ToInt32SetStrict converts a slice or array to map[int32]struct{}.
func ToInt32SetStrict(i any) map[int32]struct{} {
	m, _ := ToInt32SetStrictE(i)
	return m
}

// ToInt32SetStrictE converts a slice or array to map[int32]struct{} with error.
func ToInt32SetStrictE(i any) (map[int32]struct{}, error) {
	m, err := ToSetStrictE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[int32]struct{}); ok {
		return v, nil
	}
	return nil, fmt.Errorf("convert success but the type %T of result %#v of isn't map[int32]struct{}", m, m)
}

// ToInt64SetStrict converts a slice or array to map[int64]struct{}.
func ToInt64SetStrict(i any) map[int64]struct{} {
	m, _ := ToInt64SetStrictE(i)
	return m
}

// ToInt64SetStrictE converts a slice or array to map[int64]struct{} with error.
func ToInt64SetStrictE(i any) (map[int64]struct{}, error) {
	m, err := ToSetStrictE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[int64]struct{}); ok {
		return v, nil
	}
	return nil, fmt.Errorf("convert success but the type %T of result %#v of isn't map[int64]struct{}", m, m)
}

// ToUintSetStrict converts a slice or array to map[uint]struct{}.
func ToUintSetStrict(i any) map[uint]struct{} {
	m, _ := ToUintSetStrictE(i)
	return m
}

// ToUintSetStrictE converts a slice or array to map[uint]struct{} with error.
func ToUintSetStrictE(i any) (map[uint]struct{}, error) {
	m, err := ToSetStrictE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[uint]struct{}); ok {
		return v, nil
	}
	return nil, fmt.Errorf("convert success but the type %T of result %#v of isn't map[uint8]struct{}", m, m)
}

// ToUint8SetStrict converts a slice or array to map[uint8]struct{}.
func ToUint8SetStrict(i any) map[uint8]struct{} {
	m, _ := ToUint8SetStrictE(i)
	return m
}

// ToUint8SetStrictE converts a slice or array to map[uint8]struct{} with error.
func ToUint8SetStrictE(i any) (map[uint8]struct{}, error) {
	m, err := ToSetStrictE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[uint8]struct{}); ok {
		return v, nil
	}
	return nil, fmt.Errorf("convert success but the type %T of result %#v of isn't map[uint8]struct{}", m, m)
}

// ToUint16SetStrict converts a slice or array to map[uint16]struct{}.
func ToUint16SetStrict(i any) map[uint16]struct{} {
	m, _ := ToUint16SetStrictE(i)
	return m
}

// ToUint16SetStrictE converts a slice or array to map[uint16]struct{} with error.
func ToUint16SetStrictE(i any) (map[uint16]struct{}, error) {
	m, err := ToSetStrictE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[uint16]struct{}); ok {
		return v, nil
	}
	return nil, fmt.Errorf("convert success but the type %T of result %#v of isn't map[uint16]struct{}", m, m)
}

// ToUint32SetStrict converts a slice or array to map[uint32]struct{}.
func ToUint32SetStrict(i any) map[uint32]struct{} {
	m, _ := ToUint32SetStrictE(i)
	return m
}

// ToUint32SetStrictE converts a slice or array to map[uint32]struct{} with error.
func ToUint32SetStrictE(i any) (map[uint32]struct{}, error) {
	m, err := ToSetStrictE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[uint32]struct{}); ok {
		return v, nil
	}
	return nil, fmt.Errorf("convert success but the type %T of result %#v of isn't map[uint32]struct{}", m, m)
}

// ToUint64SetStrict converts a slice or array to map[uint64]struct{}.
func ToUint64SetStrict(i any) map[uint64]struct{} {
	m, _ := ToUint64SetStrictE(i)
	return m
}

// ToUint64SetStrictE converts a slice or array to map[uint64]struct{} with error.
func ToUint64SetStrictE(i any) (map[uint64]struct{}, error) {
	m, err := ToSetStrictE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[uint64]struct{}); ok {
		return v, nil
	}
	return nil, fmt.Errorf("convert success but the type %T of result %#v of isn't map[uint64]struct{}", m, m)
}

// ToStrSetStrict converts a slice or array to map[string]struct{}.
func ToStrSetStrict(i any) map[string]struct{} {
	m, _ := ToStrSetStrictE(i)
	return m
}

// ToStrSetStrictE converts a slice or array to map[string]struct{} with error.
func ToStrSetStrictE(i any) (map[string]struct{}, error) {
	m, err := ToSetStrictE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[string]struct{}); ok {
		return v, nil
	}
	return nil, fmt.Errorf("convert success but the type %T of result %#v of isn't map[string]struct{}", m, m)
}

// ToSetStrictE converts a slice or array to set strictly and returns an error if occurred.
// The result of map key type is equal to the type input element.
func ToSetStrictE(i any) (any, error) {
	// check params.
	if i == nil {
		return nil, fmt.Errorf("the input i is nil")
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
