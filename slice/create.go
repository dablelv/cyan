package slice

import (
	"errors"
	"fmt"
	"reflect"
)

func InsertIntSlice(src []int, index, value int) []int {
	tmp, _ := InsertSliceE(src, index, value)
	v, _ := tmp.([]int)
	return v
}

func InsertInt8Slice(src []int8, index int, value int8) []int8 {
	tmp, _ := InsertSliceE(src, index, value)
	v, _ := tmp.([]int8)
	return v
}

func InsertInt16Slice(src []int, index int, value int16) []int16 {
	tmp, _ := InsertSliceE(src, index, value)
	v, _ := tmp.([]int16)
	return v
}

func InsertInt32Slice(src []int, index int, value int32) []int32 {
	tmp, _ := InsertSliceE(src, index, value)
	v, _ := tmp.([]int32)
	return v
}

func InsertInt64Slice(src []int, index int, value int64) []int64 {
	tmp, _ := InsertSliceE(src, index, value)
	v, _ := tmp.([]int64)
	return v
}

func InsertUintSlice(src []int, index int, value uint) []uint {
	tmp, _ := InsertSliceE(src, index, value)
	v, _ := tmp.([]uint)
	return v
}

func InsertUint8Slice(src []int8, index int, value uint8) []uint8 {
	tmp, _ := InsertSliceE(src, index, value)
	v, _ := tmp.([]uint8)
	return v
}

func InsertUint16Slice(src []int, index int, value uint16) []uint16 {
	tmp, _ := InsertSliceE(src, index, value)
	v, _ := tmp.([]uint16)
	return v
}

func InsertUint32Slice(src []int, index int, value uint32) []uint32 {
	tmp, _ := InsertSliceE(src, index, value)
	v, _ := tmp.([]uint32)
	return v
}

func InsertUint64Slice(src []int, index int, value uint64) []uint64 {
	tmp, _ := InsertSliceE(src, index, value)
	v, _ := tmp.([]uint64)
	return v
}

func InsertStrSlice(src []int, index int, value string) []string {
	tmp, _ := InsertSliceE(src, index, value)
	v, _ := tmp.([]string)
	return v
}

// InsertSliceE inserts a element to slice in the specified index.
// Note that the original slice will not be modified.
func InsertSliceE(slice interface{}, index int, value interface{}) (interface{}, error) {
	// check params
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("the input %#v of type %T isn't a slice", slice, slice)
	}
	t := reflect.TypeOf(slice)
	if index < 0 || index > v.Len() || t.Elem() != reflect.TypeOf(value) {
		return nil, errors.New("param is invalid")
	}

	dst := reflect.MakeSlice(t, 0, v.Len()+1)

	// add the element to the end of slice
	if index == v.Len() {
		dst = reflect.AppendSlice(dst, v)
		dst = reflect.Append(dst, reflect.ValueOf(value))
		return dst.Interface(), nil
	}

	dst = reflect.AppendSlice(dst, v.Slice(0, index+1))
	dst = reflect.AppendSlice(dst, v.Slice(index, v.Len()))
	dst.Index(index).Set(reflect.ValueOf(value))
	return dst.Interface(), nil
}
