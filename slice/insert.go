package slice

import (
	"errors"
	"fmt"
	"reflect"
)

//
// All of the functions in this file are the supplement to the standard library slices package.
// The standard library functions https://pkg.go.dev/golang.org/x/exp/slices#Insert
// implemented by generics should be used first.
//

func InsertInts(src []int, i, v int) []int {
	tmp, _ := InsertE(src, i, v)
	s, _ := tmp.([]int)
	return s
}

func InsertInt8s(src []int8, i int, v int8) []int8 {
	tmp, _ := InsertE(src, i, v)
	s, _ := tmp.([]int8)
	return s
}

func InsertInt16s(src []int16, i int, v int16) []int16 {
	tmp, _ := InsertE(src, i, v)
	s, _ := tmp.([]int16)
	return s
}

func InsertInt32s(src []int32, i int, v int32) []int32 {
	tmp, _ := InsertE(src, i, v)
	s, _ := tmp.([]int32)
	return s
}

func InsertInt64s(src []int64, i int, v int64) []int64 {
	tmp, _ := InsertE(src, i, v)
	s, _ := tmp.([]int64)
	return s
}

func InsertUints(src []uint, i int, v uint) []uint {
	tmp, _ := InsertE(src, i, v)
	s, _ := tmp.([]uint)
	return s
}

func InsertUint8s(src []uint8, i int, v uint8) []uint8 {
	tmp, _ := InsertE(src, i, v)
	s, _ := tmp.([]uint8)
	return s
}

func InsertUint16s(src []uint16, i int, v uint16) []uint16 {
	tmp, _ := InsertE(src, i, v)
	s, _ := tmp.([]uint16)
	return s
}

func InsertUint32s(src []uint32, i int, v uint32) []uint32 {
	tmp, _ := InsertE(src, i, v)
	s, _ := tmp.([]uint32)
	return s
}

func InsertUint64s(src []uint64, i int, v uint64) []uint64 {
	tmp, _ := InsertE(src, i, v)
	s, _ := tmp.([]uint64)
	return s
}

func InsertStrs(src []string, i int, v string) []string {
	tmp, _ := InsertE(src, i, v)
	s, _ := tmp.([]string)
	return s
}

// InsertE inserts a element to slice in the specified index.
// Note that the original slice will not be modified.
func InsertE(slice any, i int, v any) (any, error) {
	// Check arguments.
	val := reflect.ValueOf(slice)
	if val.Kind() != reflect.Slice {
		return nil, fmt.Errorf("The input %#v of type %T isn't a slice", slice, slice)
	}
	t := reflect.TypeOf(slice)
	if i < 0 || i > val.Len() || t.Elem() != reflect.TypeOf(v) {
		return nil, errors.New("Arguments are invalid")
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

// Insert inserts the values v... into s at index i and returns the result slice.
// Note that unlike the standard library function Insert, it does not modify the original slice.
func Insert[S ~[]E, E any](s S, i int, v ...E) S {

}
