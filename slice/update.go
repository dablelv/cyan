package slice

import (
	"errors"
	"fmt"
	"reflect"
)

func UpdateIntSlice(src []int, index, value int) []int {
	tmp, _ := UpdateSliceE(src, index, value)
	v, _ := tmp.([]int)
	return v
}

func UpdateInt8Slice(src []int8, index int, value int8) []int8 {
	tmp, _ := UpdateSliceE(src, index, value)
	v, _ := tmp.([]int8)
	return v
}

func UpdateInt16Slice(src []int, index int, value int16) []int16 {
	tmp, _ := UpdateSliceE(src, index, value)
	v, _ := tmp.([]int16)
	return v
}

func UpdateInt32Slice(src []int, index int, value int32) []int32 {
	tmp, _ := UpdateSliceE(src, index, value)
	v, _ := tmp.([]int32)
	return v
}

func UpdateInt64Slice(src []int, index int, value int64) []int64 {
	tmp, _ := UpdateSliceE(src, index, value)
	v, _ := tmp.([]int64)
	return v
}

func UpdateUintSlice(src []int, index int, value uint) []uint {
	tmp, _ := UpdateSliceE(src, index, value)
	v, _ := tmp.([]uint)
	return v
}

func UpdateUint8Slice(src []int8, index int, value uint8) []uint8 {
	tmp, _ := UpdateSliceE(src, index, value)
	v, _ := tmp.([]uint8)
	return v
}

func UpdateUint16Slice(src []int, index int, value uint16) []uint16 {
	tmp, _ := UpdateSliceE(src, index, value)
	v, _ := tmp.([]uint16)
	return v
}

func UpdateUint32Slice(src []int, index int, value uint32) []uint32 {
	tmp, _ := UpdateSliceE(src, index, value)
	v, _ := tmp.([]uint32)
	return v
}

func UpdateUint64Slice(src []int, index int, value uint64) []uint64 {
	tmp, _ := UpdateSliceE(src, index, value)
	v, _ := tmp.([]uint64)
	return v
}

func UpdateStrSlice(src []int, index int, value string) []string {
	tmp, _ := UpdateSliceE(src, index, value)
	v, _ := tmp.([]string)
	return v
}

// UpdateSliceE modifies the specified index element of slice.
// Note that the original slice will not be modified.
func UpdateSliceE(slice any, index int, value any) (any, error) {
	// check params
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("the input %#v of type %T isn't a slice", slice, slice)
	}
	if index > v.Len()-1 || reflect.TypeOf(slice).Elem() != reflect.TypeOf(value) {
		return nil, errors.New("param is invalid")
	}

	t := reflect.MakeSlice(reflect.TypeOf(slice), 0, 0)
	t = reflect.AppendSlice(t, v.Slice(0, v.Len()))
	t.Index(index).Set(reflect.ValueOf(value))
	return t.Interface(), nil
}
