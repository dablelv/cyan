package slice

import (
	"errors"
	"fmt"
	"reflect"
)

//
// Note that after Go 1.18, this file is deprecated.
// Please use the standard lib function https://pkg.go.dev/golang.org/x/exp/slices#Insert
// implemented by generics.
//

func InsertIntSlice(src []int, i, v int) []int {
	tmp, _ := InsertSliceE(src, i, v)
	s, _ := tmp.([]int)
	return s
}

func InsertInt8Slice(src []int8, i int, v int8) []int8 {
	tmp, _ := InsertSliceE(src, i, v)
	s, _ := tmp.([]int8)
	return s
}

func InsertInt16Slice(src []int16, i int, v int16) []int16 {
	tmp, _ := InsertSliceE(src, i, v)
	s, _ := tmp.([]int16)
	return s
}

func InsertInt32Slice(src []int32, i int, v int32) []int32 {
	tmp, _ := InsertSliceE(src, i, v)
	s, _ := tmp.([]int32)
	return s
}

func InsertInt64Slice(src []int64, i int, v int64) []int64 {
	tmp, _ := InsertSliceE(src, i, v)
	s, _ := tmp.([]int64)
	return s
}

func InsertUintSlice(src []uint, i int, v uint) []uint {
	tmp, _ := InsertSliceE(src, i, v)
	s, _ := tmp.([]uint)
	return s
}

func InsertUint8Slice(src []uint8, i int, v uint8) []uint8 {
	tmp, _ := InsertSliceE(src, i, v)
	s, _ := tmp.([]uint8)
	return s
}

func InsertUint16Slice(src []uint16, i int, v uint16) []uint16 {
	tmp, _ := InsertSliceE(src, i, v)
	s, _ := tmp.([]uint16)
	return s
}

func InsertUint32Slice(src []uint32, i int, v uint32) []uint32 {
	tmp, _ := InsertSliceE(src, i, v)
	s, _ := tmp.([]uint32)
	return s
}

func InsertUint64Slice(src []uint64, i int, v uint64) []uint64 {
	tmp, _ := InsertSliceE(src, i, v)
	s, _ := tmp.([]uint64)
	return s
}

func InsertStrSlice(src []string, i int, v string) []string {
	tmp, _ := InsertSliceE(src, i, v)
	s, _ := tmp.([]string)
	return s
}

// InsertSliceE inserts a element to slice in the specified index.
// Note that the original slice will not be modified.
func InsertSliceE(slice any, i int, v any) (any, error) {
	// Check params.
	val := reflect.ValueOf(slice)
	if val.Kind() != reflect.Slice {
		return nil, fmt.Errorf("the input %#v of type %T isn't a slice", slice, slice)
	}
	t := reflect.TypeOf(slice)
	if i < 0 || i > val.Len() || t.Elem() != reflect.TypeOf(v) {
		return nil, errors.New("param is invalid")
	}

	dst := reflect.MakeSlice(t, 0, val.Len()+1)

	// Add the element to the end of slice.
	if i == val.Len() {
		dst = reflect.AppendSlice(dst, val)
		dst = reflect.Append(dst, reflect.ValueOf(v))
		return dst.Interface(), nil
	}

	dst = reflect.AppendSlice(dst, val.Slice(0, i+1))
	dst = reflect.AppendSlice(dst, val.Slice(i, val.Len()))
	dst.Index(i).Set(reflect.ValueOf(v))
	return dst.Interface(), nil
}
