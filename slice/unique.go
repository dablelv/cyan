package slice

import (
	"fmt"
	"reflect"
)

//
// Desc: unique a slice, e.g. input []int32{1, 2, 2, 3} and output is []int32{1, 2, 3}.
//

func UniqueIntSlice(src []int) []int {
	dst, _ := UniqueSliceE(src)
	v, _ := dst.([]int)
	return v
}

func UniqueInt8Slice(src []int8) []int8 {
	dst, _ := UniqueSliceE(src)
	v, _ := dst.([]int8)
	return v
}

func UniqueInt16Slice(src []int16) []int16 {
	dst, _ := UniqueSliceE(src)
	v, _ := dst.([]int16)
	return v
}

func UniqueInt32Slice(src []int32) []int32 {
	dst, _ := UniqueSliceE(src)
	v, _ := dst.([]int32)
	return v
}

func UniqueInt64Slice(src []int64) []int64 {
	dst, _ := UniqueSliceE(src)
	v, _ := dst.([]int64)
	return v
}

func UniqueUintSlice(src []uint) []uint {
	dst, _ := UniqueSliceE(src)
	v, _ := dst.([]uint)
	return v
}

func UniqueUint8Slice(src []uint8) []uint8 {
	dst, _ := UniqueSliceE(src)
	v, _ := dst.([]uint8)
	return v
}

func UniqueUint16Slice(src []uint16) []uint16 {
	dst, _ := UniqueSliceE(src)
	v, _ := dst.([]uint16)
	return v
}

func UniqueUint32Slice(src []uint32) []uint32 {
	dst, _ := UniqueSliceE(src)
	v, _ := dst.([]uint32)
	return v
}

func UniqueUint64Slice(src []uint64) []uint64 {
	dst, _ := UniqueSliceE(src)
	v, _ := dst.([]uint64)
	return v
}

func UniqueFloat32Slice(src []float32) []float32 {
	dst, _ := UniqueSliceE(src)
	v, _ := dst.([]float32)
	return v
}

func UniqueFloat64Slice(src []float64) []float64 {
	dst, _ := UniqueSliceE(src)
	v, _ := dst.([]float64)
	return v
}

func UniqueStrSlice(src []string) []string {
	dst, _ := UniqueSliceE(src)
	v, _ := dst.([]string)
	return v
}

// UniqueSliceE deletes repeated elements in a slice with error.
// Note that the original slice will not be modified.
func UniqueSliceE(slice any) (any, error) {
	// Check params.
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("the input %#v of type %T isn't a slice", slice, slice)
	}
	// Unique the slice.
	dst := reflect.MakeSlice(reflect.TypeOf(slice), 0, v.Len())
	m := make(map[any]struct{})
	for i := 0; i < v.Len(); i++ {
		if _, ok := m[v.Index(i).Interface()]; !ok {
			dst = reflect.Append(dst, v.Index(i))
			m[v.Index(i).Interface()] = struct{}{}
		}
	}
	return dst.Interface(), nil
}
