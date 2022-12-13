package slice

import (
	"fmt"
	"reflect"
)

//
// Desc: reverse a slice, e.g. input []int32{1, 2, 3} and output is []int32{3, 2, 1}.
//

func ReverseIntSlice(src []int) []int {
	dst, _ := ReverseSliceE(src)
	v, _ := dst.([]int)
	return v
}

func ReverseInt8Slice(src []int8) []int8 {
	dst, _ := ReverseSliceE(src)
	v, _ := dst.([]int8)
	return v
}

func ReverseInt16Slice(src []int16) []int16 {
	dst, _ := ReverseSliceE(src)
	v, _ := dst.([]int16)
	return v
}

func ReverseInt32Slice(src []int32) []int32 {
	dst, _ := ReverseSliceE(src)
	v, _ := dst.([]int32)
	return v
}

func ReverseInt64Slice(src []int64) []int64 {
	dst, _ := ReverseSliceE(src)
	v, _ := dst.([]int64)
	return v
}

func ReverseUintSlice(src []uint) []uint {
	dst, _ := ReverseSliceE(src)
	v, _ := dst.([]uint)
	return v
}

func ReverseUint8Slice(src []uint8) []uint8 {
	dst, _ := ReverseSliceE(src)
	v, _ := dst.([]uint8)
	return v
}

func ReverseUint16Slice(src []uint16) []uint16 {
	dst, _ := ReverseSliceE(src)
	v, _ := dst.([]uint16)
	return v
}

func ReverseUint32Slice(src []uint32) []uint32 {
	dst, _ := ReverseSliceE(src)
	v, _ := dst.([]uint32)
	return v
}

func ReverseUint64Slice(src []uint64) []uint64 {
	dst, _ := ReverseSliceE(src)
	v, _ := dst.([]uint64)
	return v
}

func ReverseStrSlice(src []string) []string {
	dst, _ := ReverseSliceE(src)
	v, _ := dst.([]string)
	return v
}

// ReverseSliceE reverses the specified slice without modifying the original slice.
func ReverseSliceE(slice interface{}) (interface{}, error) {
	// check params
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("the input %#v of type %T isn't a slice", slice, slice)
	}
	// reverse the slice
	dst := reflect.MakeSlice(reflect.TypeOf(slice), 0, v.Len())
	for i := v.Len() - 1; i >= 0; i-- {
		dst = reflect.Append(dst, v.Index(i))
	}
	return dst.Interface(), nil
}
