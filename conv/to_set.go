package conv

import (
	"fmt"
	"reflect"
)

//
// Convert an any element type of slice or array to to specified type map set.
// Note that the the element type of input don't need to be equal to the map key type.
// For example, []uint64{1, 2, 3} can be converted to map[uint64]struct{}{1:struct{}, 2:struct{},
// 3:struct{}} and also can be converted to map[string]struct{}{"1":struct{}, "2":struct{}, "3":struct{}}
// if you want.
//

// ToBoolSet converts a slice or array to map[bool]struct{}.
func ToBoolSet(i any) map[bool]struct{} {
	m, _ := ToBoolSetE(i)
	return m
}

// ToBoolSetE converts a slice or array to map[bool]struct{} with error.
func ToBoolSetE(i any) (map[bool]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[bool]struct{}); ok {
		return v, nil
	}
	dst := make(map[bool]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToBoolE(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// ToIntSet converts a slice or array to map[int]struct{}.
func ToIntSet(i any) map[int]struct{} {
	m, _ := ToIntSetE(i)
	return m
}

// ToIntSetE converts a slice or array to map[int]struct{Set} with error.
func ToIntSetE(i any) (map[int]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[int]struct{}); ok {
		return v, nil
	}
	dst := make(map[int]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToIntE(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// ToInt8Set converts a slice or array to map[int8]struct{}.
func ToInt8Set(i any) map[int8]struct{} {
	m, _ := ToInt8SetE(i)
	return m
}

// ToInt8SetE converts a slice or array to map[int8]struct{} with error.
func ToInt8SetE(i any) (map[int8]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[int8]struct{}); ok {
		return v, nil
	}
	dst := make(map[int8]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToInt8E(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// ToInt16Set converts a slice or array to map[int16]struct{}.
func ToInt16Set(i any) map[int16]struct{} {
	m, _ := ToInt16SetE(i)
	return m
}

// ToInt16SetE converts a slice or array to map[int16]struct{} with error.
func ToInt16SetE(i any) (map[int16]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[int16]struct{}); ok {
		return v, nil
	}
	dst := make(map[int16]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToInt16E(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// ToInt32Set converts a slice or array to map[int32]struct{}.
func ToInt32Set(i any) map[int32]struct{} {
	m, _ := ToInt32SetE(i)
	return m
}

// ToInt32SetE converts a slice or array to map[int32]struct{} with error.
func ToInt32SetE(i any) (map[int32]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[int32]struct{}); ok {
		return v, nil
	}
	dst := make(map[int32]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToInt32E(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// ToInt64Set converts a slice or array to map[int64]struct{}.
func ToInt64Set(i any) map[int64]struct{} {
	m, _ := ToInt64SetE(i)
	return m
}

// ToInt64SetE converts a slice or array to map[int64]struct{} with error.
func ToInt64SetE(i any) (map[int64]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[int64]struct{}); ok {
		return v, nil
	}
	dst := make(map[int64]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToInt64E(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// ToUintSet converts a slice or array to map[uint]struct{}.
func ToUintSet(i any) map[uint]struct{} {
	m, _ := ToUintSetE(i)
	return m
}

// ToUintSetE converts a slice or array to map[uint]struct{} with error.
func ToUintSetE(i any) (map[uint]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[uint]struct{}); ok {
		return v, nil
	}
	dst := make(map[uint]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToUintE(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// ToUint8Set converts a slice or array to map[uint8]struct{}.
func ToUint8Set(i any) map[uint8]struct{} {
	m, _ := ToUint8SetE(i)
	return m
}

// ToUint8SetE converts a slice or array to map[uint8]struct{} with error.
func ToUint8SetE(i any) (map[uint8]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[uint8]struct{}); ok {
		return v, nil
	}
	dst := make(map[uint8]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToUint8E(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// ToUint16Set converts a slice or array to map[uint16]struct{}.
func ToUint16Set(i any) map[uint16]struct{} {
	m, _ := ToUint16SetE(i)
	return m
}

// ToUint16SetE converts a slice or array to map[uint16]struct{} with error.
func ToUint16SetE(i any) (map[uint16]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[uint16]struct{}); ok {
		return v, nil
	}
	dst := make(map[uint16]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToUint16E(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// ToUint32Set converts a slice or array to map[uint32]struct{}.
func ToUint32Set(i any) map[uint32]struct{} {
	m, _ := ToUint32SetE(i)
	return m
}

// ToUint32SetE converts a slice or array to map[uint32]struct{} with error.
func ToUint32SetE(i any) (map[uint32]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[uint32]struct{}); ok {
		return v, nil
	}
	dst := make(map[uint32]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToUint32E(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// ToUint64Set converts a slice or array to map[uint64]struct{}.
func ToUint64Set(i any) map[uint64]struct{} {
	m, _ := ToUint64SetE(i)
	return m
}

// ToUint64SetE converts a slice or array to map[uint64]struct{} with error.
func ToUint64SetE(i any) (map[uint64]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[uint64]struct{}); ok {
		return v, nil
	}
	dst := make(map[uint64]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToUint64E(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// ToFloat32Set converts a slice or array to map[float32]struct{}.
func ToFloat32Set(i any) map[float32]struct{} {
	m, _ := ToFloat32SetE(i)
	return m
}

// ToFloat32SetE converts a slice or array to map[float32]struct{} and returns an error if occurred.
func ToFloat32SetE(i any) (map[float32]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[float32]struct{}); ok {
		return v, nil
	}
	dst := make(map[float32]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToFloat32E(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// ToFloat64Set converts a slice or array to map[float64]struct{}.
func ToFloat64Set(i any) map[float64]struct{} {
	m, _ := ToFloat64SetE(i)
	return m
}

// ToFloat64SetE converts a slice or array to map[float64]struct{} and returns an error if occurred.
func ToFloat64SetE(i any) (map[float64]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[float64]struct{}); ok {
		return v, nil
	}
	dst := make(map[float64]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToFloat64E(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// ToStrSet converts a slice or array to map[string]struct{}.
func ToStrSet(i any) map[string]struct{} {
	m, _ := ToStrSetE(i)
	return m
}

// ToStrSetE converts a slice or array to map[string]struct{} and returns an error if occurred.
func ToStrSetE(i any) (map[string]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[string]struct{}); ok {
		return v, nil
	}
	dst := make(map[string]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToStringE(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// toSetE converts a slice or array to map[any]struct{} and returns an error if occurred.
func toSetE(i any) (any, error) {
	// Check params.
	if i == nil {
		return nil, fmt.Errorf("the input i is nil")
	}
	t := reflect.TypeOf(i)
	kind := t.Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return nil, fmt.Errorf("the input %#v of type %T isn't a slice or array", i, i)
	}

	// Execute the conversion.
	v := reflect.ValueOf(i)
	mT := reflect.MapOf(t.Elem(), reflect.TypeOf(struct{}{}))
	mV := reflect.MakeMapWithSize(mT, v.Len())
	for j := 0; j < v.Len(); j++ {
		mV.SetMapIndex(v.Index(j), reflect.ValueOf(struct{}{}))
	}
	return mV.Interface(), nil
}
