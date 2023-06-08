package slice

import (
	"fmt"
	"reflect"
)

//
// Reverse a slice, e.g. input []int32{1, 2, 3} and output is []int32{3, 2, 1}.
//

// Reverse reverses the specified slice without modifying the original slice.
// Reverse implemented by generics is recommended to be used.
func Reverse[E any, S ~[]E](s S) S {
	if s == nil {
		return s
	}
	r := make([]E, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return r
}

func ReverseInt(src []int) []int {
	if src == nil {
		return nil
	}
	dst, _ := ReverseE(src)
	v, _ := dst.([]int)
	return v
}

func ReverseInt8(src []int8) []int8 {
	if src == nil {
		return nil
	}
	dst, _ := ReverseE(src)
	v, _ := dst.([]int8)
	return v
}

func ReverseInt16(src []int16) []int16 {
	if src == nil {
		return nil
	}
	dst, _ := ReverseE(src)
	v, _ := dst.([]int16)
	return v
}

func ReverseInt32(src []int32) []int32 {
	if src == nil {
		return nil
	}
	dst, _ := ReverseE(src)
	v, _ := dst.([]int32)
	return v
}

func ReverseInt64(src []int64) []int64 {
	if src == nil {
		return nil
	}
	dst, _ := ReverseE(src)
	v, _ := dst.([]int64)
	return v
}

func ReverseUint(src []uint) []uint {
	if src == nil {
		return nil
	}
	dst, _ := ReverseE(src)
	v, _ := dst.([]uint)
	return v
}

func ReverseUint8(src []uint8) []uint8 {
	if src == nil {
		return nil
	}
	dst, _ := ReverseE(src)
	v, _ := dst.([]uint8)
	return v
}

func ReverseUint16(src []uint16) []uint16 {
	if src == nil {
		return nil
	}
	dst, _ := ReverseE(src)
	v, _ := dst.([]uint16)
	return v
}

func ReverseUint32(src []uint32) []uint32 {
	if src == nil {
		return nil
	}
	dst, _ := ReverseE(src)
	v, _ := dst.([]uint32)
	return v
}

func ReverseUint64(src []uint64) []uint64 {
	if src == nil {
		return nil
	}
	dst, _ := ReverseE(src)
	v, _ := dst.([]uint64)
	return v
}

func ReverseStr(src []string) []string {
	if src == nil {
		return nil
	}
	dst, _ := ReverseE(src)
	v, _ := dst.([]string)
	return v
}

// ReverseRefE implemented by reflect reverses the specified slice without modifying the original slice.
func ReverseE(slice any) (any, error) {
	// Check params.
	if slice == nil {
		return slice, nil
	}
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("the input %#v of type %T isn't a slice", slice, slice)
	}
	// Reverse slice.
	dst := reflect.MakeSlice(reflect.TypeOf(slice), 0, v.Len())
	for i := v.Len() - 1; i >= 0; i-- {
		dst = reflect.Append(dst, v.Index(i))
	}
	return dst.Interface(), nil
}
