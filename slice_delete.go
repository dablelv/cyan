package util

import (
	"errors"
	"reflect"
)

// DeleteStrSlice deletes string slice elements by indexes.
func DeleteStrSlice(src []string, indexes ...int) []string {
	return DeleteSlice(src, indexes...).([]string)
}

// DeleteIntSlice deletes int slice elements by indexes.
func DeleteIntSlice(src []int, indexes ...int) []int {
	return DeleteSlice(src, indexes...).([]int)
}

// DeleteInt8Slice deletes int8 slice elements by indexes.
func DeleteInt8Slice(src []int8, indexes ...int) []int8 {
	return DeleteSlice(src, indexes...).([]int8)
}

// DeleteInt16Slice deletes int16 slice elements by indexes.
func DeleteInt16Slice(src []int16, indexes ...int) []int16 {
	return DeleteSlice(src, indexes...).([]int16)
}

// DeleteInt32Slice deletes int32 slice elements by indexes.
func DeleteInt32Slice(src []int32, indexes ...int) []int32 {
	return DeleteSlice(src, indexes...).([]int32)
}

// DeleteInt64Slice deletes int64 slice elements by indexes.
func DeleteInt64Slice(src []int64, indexes ...int) []int64 {
	return DeleteSlice(src, indexes...).([]int64)
}

// DeleteUintSlice deletes uint slice elements by indexes.
func DeleteUintSlice(src []int, indexes ...int) []uint {
	return DeleteSlice(src, indexes...).([]uint)
}

// DeleteUint8Slice deletes uint8 slice elements by indexes.
func DeleteUint8Slice(src []int8, indexes ...int) []uint8 {
	return DeleteSlice(src, indexes...).([]uint8)
}

// DeleteUint16Slice deletes uint16 slice elements by indexes.
func DeleteUint16Slice(src []int, indexes ...int) []uint16 {
	return DeleteSlice(src, indexes...).([]uint16)
}

// DeleteUint32Slice deletes uint32 slice elements by indexes.
func DeleteUint32Slice(src []uint32, indexes ...int) []uint32 {
	return DeleteSlice(src, indexes...).([]uint32)
}

// DeleteUint64Slice deletes uint64 slice elements by indexes.
func DeleteUint64Slice(src []uint64, indexes ...int) []uint64 {
	return DeleteSlice(src, indexes...).([]uint64)
}

// DeleteSliceElms deletes the specified elements from the slice.
// Note that the original slice will not be modified.
func DeleteSliceElms(i interface{}, elms ...interface{}) interface{} {
	res, _ := DeleteSliceElmsE(i, elms...)
	return res
}

// DeleteSliceElmsE deletes the specified elements from the slice.
// Note that the original slice will not be modified.
func DeleteSliceElmsE(i interface{}, elms ...interface{}) (interface{}, error) {
	// check params
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Slice {
		return nil, errors.New("the input isn't a slice")
	}
	if v.Len() == 0 || len(elms) == 0 {
		return i, nil
	}
	if reflect.TypeOf(i).Elem() != reflect.TypeOf(elms[0]) {
		return nil, errors.New("element type is ill")
	}
	// convert the elements to map set
	m := make(map[interface{}]struct{})
	for _, v := range elms {
		m[v] = struct{}{}
	}
	// filter out specified elements
	t := reflect.MakeSlice(reflect.TypeOf(i), 0, v.Len())
	for i := 0; i < v.Len(); i++ {
		if _, ok := m[v.Index(i).Interface()]; !ok {
			t = reflect.Append(t, v.Index(i))
		}
	}
	return t.Interface(), nil
}

// DeleteSlice deletes the specified index element from the slice.
// Note that the original slice will not be modified.
func DeleteSlice(slice interface{}, indexes ...int) interface{} {
	res, _ := DeleteSliceE(slice, indexes...)
	return res
}

// DeleteSliceE deletes the specified index element from the slice with error.
// Note that the original slice will not be modified.
func DeleteSliceE(slice interface{}, indexes ...int) (interface{}, error) {
	// check params
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, errors.New("the input isn't a slice")
	}
	if v.Len() == 0 || len(indexes) == 0 {
		return slice, nil
	}
	// convert the indexes to map set
	m := make(map[int]struct{})
	for _, i := range indexes {
		m[i] = struct{}{}
	}
	// delete
	t := reflect.MakeSlice(reflect.TypeOf(slice), 0, v.Len())
	for i := 0; i < v.Len(); i++ {
		if _, ok := m[i]; !ok {
			t = reflect.Append(t, v.Index(i))
		}
	}
	return t.Interface(), nil
}
