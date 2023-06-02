package conv

import (
	"fmt"
	"reflect"
)

//
// Converts an any element type slice or array to the specified type mapping set.
// Note that the the element type of input don't need to be equal to the map key type.
// For example, []uint64{1, 2, 3} can be converted to map[uint64]struct{}{1:{}, 2:{},3:{}}
// and also can be converted to map[string]struct{}{"1":{}, "2":{}, "3":{}} if you want.
//

// ToSet converts a slice or array to map[T]struct{}.
// An error will be returned if an error occurs.
func ToSet[T comparable](i any) map[T]struct{} {
	m, _ := ToSetE[T](i)
	return m
}

// ToSetE converts a slice or array to map[T]struct{} and returns an error if occurred.
// Note that the the element type of input don't need to be equal to the map key type.
// For example, []uint64{1, 2, 3} can be converted to map[uint64]struct{}{1:{},2:{},3:{}}
// and also can be converted to map[string]struct{}{"1":{},"2":{},"3":{}} if you want.
// Note that this function is implemented through 1.18 generics, so the element type needs to
// be specified when calling it, e.g. ToSetE[int]([]int{1,2,3}).
func ToSetE[T comparable](a any) (map[T]struct{}, error) {
	if a == nil {
		return nil, nil
	}
	t := reflect.TypeOf(a)
	kind := t.Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return nil, fmt.Errorf("the type %T of input %#v isn't a slice or array", a, a)
	}

	// Execute the conversion.
	v := reflect.ValueOf(a)
	mapT := reflect.MapOf(t.Elem(), reflect.TypeOf(struct{}{}))
	mapV := reflect.MakeMapWithSize(mapT, v.Len())
	for i := 0; i < v.Len(); i++ {
		mapV.SetMapIndex(v.Index(i), reflect.ValueOf(struct{}{}))
	}
	if v, ok := mapV.Interface().(map[T]struct{}); ok {
		return v, nil
	}
	// Convert the element to the type T.
	set := make(map[T]struct{}, v.Len())
	for _, k := range mapV.MapKeys() {
		v, err := ToAnyE[T](k.Interface())
		if err != nil {
			return nil, err
		}
		set[v] = struct{}{}
	}
	return set, nil
}

// ToBoolSet converts a slice or array to map[bool]struct{}.
func ToBoolSet(a any) map[bool]struct{} {
	return ToSet[bool](a)
}

// ToBoolSetE converts a slice or array to map[bool]struct{} with error.
func ToBoolSetE(a any) (map[bool]struct{}, error) {
	return ToSetE[bool](a)
}

// ToIntSet converts a slice or array to map[int]struct{}.
func ToIntSet(a any) map[int]struct{} {
	return ToSet[int](a)
}

// ToIntSetE converts a slice or array to map[int]struct{Set} with error.
func ToIntSetE(a any) (map[int]struct{}, error) {
	return ToSetE[int](a)
}

// ToInt8Set converts a slice or array to map[int8]struct{}.
func ToInt8Set(a any) map[int8]struct{} {
	return ToSet[int8](a)
}

// ToInt8SetE converts a slice or array to map[int8]struct{} with error.
func ToInt8SetE(a any) (map[int8]struct{}, error) {
	return ToSetE[int8](a)
}

// ToInt16Set converts a slice or array to map[int16]struct{}.
func ToInt16Set(a any) map[int16]struct{} {
	return ToSet[int16](a)
}

// ToInt16SetE converts a slice or array to map[int16]struct{} with error.
func ToInt16SetE(a any) (map[int16]struct{}, error) {
	return ToSetE[int16](a)
}

// ToInt32Set converts a slice or array to map[int32]struct{}.
func ToInt32Set(a any) map[int32]struct{} {
	return ToSet[int32](a)
}

// ToInt32SetE converts a slice or array to map[int32]struct{} with error.
func ToInt32SetE(a any) (map[int32]struct{}, error) {
	return ToSetE[int32](a)
}

// ToInt64Set converts a slice or array to map[int64]struct{}.
func ToInt64Set(a any) map[int64]struct{} {
	return ToSet[int64](a)
}

// ToInt64SetE converts a slice or array to map[int64]struct{} with error.
func ToInt64SetE(a any) (map[int64]struct{}, error) {
	return ToSetE[int64](a)
}

// ToUintSet converts a slice or array to map[uint]struct{}.
func ToUintSet(a any) map[uint]struct{} {
	return ToSet[uint](a)
}

// ToUintSetE converts a slice or array to map[uint]struct{} with error.
func ToUintSetE(a any) (map[uint]struct{}, error) {
	return ToSetE[uint](a)
}

// ToUint8Set converts a slice or array to map[uint8]struct{}.
func ToUint8Set(a any) map[uint8]struct{} {
	return ToSet[uint8](a)
}

// ToUint8SetE converts a slice or array to map[uint8]struct{} with error.
func ToUint8SetE(a any) (map[uint8]struct{}, error) {
	return ToSetE[uint8](a)
}

// ToUint16Set converts a slice or array to map[uint16]struct{}.
func ToUint16Set(a any) map[uint16]struct{} {
	return ToSet[uint16](a)
}

// ToUint16SetE converts a slice or array to map[uint16]struct{} with error.
func ToUint16SetE(a any) (map[uint16]struct{}, error) {
	return ToSetE[uint16](a)
}

// ToUint32Set converts a slice or array to map[uint32]struct{}.
func ToUint32Set(a any) map[uint32]struct{} {
	return ToSet[uint32](a)
}

// ToUint32SetE converts a slice or array to map[uint32]struct{} with error.
func ToUint32SetE(a any) (map[uint32]struct{}, error) {
	return ToSetE[uint32](a)
}

// ToUint64Set converts a slice or array to map[uint64]struct{}.
func ToUint64Set(a any) map[uint64]struct{} {
	return ToSet[uint64](a)
}

// ToUint64SetE converts a slice or array to map[uint64]struct{} with error.
func ToUint64SetE(a any) (map[uint64]struct{}, error) {
	return ToSetE[uint64](a)
}

// ToFloat32Set converts a slice or array to map[float32]struct{}.
func ToFloat32Set(a any) map[float32]struct{} {
	return ToSet[float32](a)
}

// ToFloat32SetE converts a slice or array to map[float32]struct{} and returns an error if occurred.
func ToFloat32SetE(a any) (map[float32]struct{}, error) {
	return ToSetE[float32](a)
}

// ToFloat64Set converts a slice or array to map[float64]struct{}.
func ToFloat64Set(a any) map[float64]struct{} {
	return ToSet[float64](a)
}

// ToFloat64SetE converts a slice or array to map[float64]struct{} and returns an error if occurred.
func ToFloat64SetE(a any) (map[float64]struct{}, error) {
	return ToSetE[float64](a)
}

// ToStrSet converts a slice or array to map[string]struct{}.
func ToStrSet(a any) map[string]struct{} {
	return ToSet[string](a)
}

// ToStrSetE converts a slice or array to map[string]struct{} and returns an error if occurred.
func ToStrSetE(a any) (map[string]struct{}, error) {
	return ToSetE[string](a)
}
