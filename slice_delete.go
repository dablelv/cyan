package util

import (
	"errors"
	"reflect"
)

// DeleteStrSliceElms deletes the specified elements from string slice.
func DeleteStrSliceElms(i []string, elms ...interface{}) []string {
	res, err := DeleteSliceElmsE(i, elms...)
	if err != nil {
		return nil
	}
	return res.([]string)
}

// DeleteStrSliceElmsE deletes the specified elements from string slice with error.
func DeleteStrSliceElmsE(i []string, elms ...interface{}) ([]string, error) {
	res, err := DeleteSliceElmsE(i, elms...)
	if err != nil {
		return nil, err
	}
	return res.([]string), nil
}

// DeleteIntSliceElms deletes the specified elements from int slice.
func DeleteIntSliceElms(i []int, elms ...interface{}) []int {
	res, err := DeleteSliceElmsE(i, elms...)
	if err != nil {
		return nil
	}
	return res.([]int)
}

// DeleteIntSliceElmsE deletes the specified elements from int slice with error.
func DeleteIntSliceElmsE(i []int, elms ...interface{}) ([]int, error) {
	res, err := DeleteSliceElmsE(i, elms...)
	if err != nil {
		return nil, err
	}
	return res.([]int), nil
}

// DeleteInt8SliceElms deletes the specified elements from int8 slice.
func DeleteInt8SliceElms(i []int8, elms ...interface{}) []int8 {
	res, err := DeleteSliceElmsE(i, elms...)
	if err != nil {
		return nil
	}
	return res.([]int8)
}

// DeleteInt8SliceElmsE deletes the specified elements from int8 slice with error.
func DeleteInt8SliceElmsE(i []int8, elms ...interface{}) ([]int8, error) {
	res, err := DeleteSliceElmsE(i, elms...)
	if err != nil {
		return nil, err
	}
	return res.([]int8), nil
}

// DeleteInt16SliceElms deletes the specified elements from int16 slice.
func DeleteInt16SliceElms(i []int16, elms ...interface{}) []int16 {
	res, err := DeleteSliceElmsE(i, elms...)
	if err != nil {
		return nil
	}
	return res.([]int16)
}

// DeleteInt16SliceElmsE deletes the specified elements from int16 slice with error.
func DeleteInt16SliceElmsE(i []int16, elms ...interface{}) ([]int16, error) {
	res, err := DeleteSliceElmsE(i, elms...)
	if err != nil {
		return nil, err
	}
	return res.([]int16), nil
}

// DeleteInt32SliceElms deletes the specified elements from int32 slice.
func DeleteInt32SliceElms(i []int32, elms ...interface{}) []int32 {
	res, err := DeleteSliceElmsE(i, elms...)
	if err != nil {
		return nil
	}
	return res.([]int32)
}

// DeleteInt32SliceElmsE deletes the specified elements from int32 slice with error.
func DeleteInt32SliceElmsE(i []int32, elms ...interface{}) ([]int32, error) {
	res, err := DeleteSliceElmsE(i, elms...)
	if err != nil {
		return nil, err
	}
	return res.([]int32), nil
}

// DeleteInt64SliceElms deletes the specified elements from int64 slice.
func DeleteInt64SliceElms(i []int64, elms ...interface{}) []int64 {
	res, err := DeleteSliceElmsE(i, elms...)
	if err != nil {
		return nil
	}
	return res.([]int64)
}

// DeleteInt64SliceElmsE deletes the specified elements from int64 slice with error.
func DeleteInt64SliceElmsE(i []int64, elms ...interface{}) ([]int64, error) {
	res, err := DeleteSliceElmsE(i, elms...)
	if err != nil {
		return nil, err
	}
	return res.([]int64), nil
}

// DeleteUintSliceElms deletes the specified elements from uint slice.
func DeleteUintSliceElms(i []uint, elms ...interface{}) []uint {
	res, err := DeleteSliceElmsE(i, elms...)
	if err != nil {
		return nil
	}
	return res.([]uint)
}

// DeleteUintSliceElmsE deletes the specified elements from int slice with error.
func DeleteUintSliceElmsE(i []uint, elms ...interface{}) ([]uint, error) {
	res, err := DeleteSliceElmsE(i, elms...)
	if err != nil {
		return nil, err
	}
	return res.([]uint), nil
}

// DeleteUint8SliceElms deletes the specified elements from uint8 slice.
func DeleteUint8SliceElms(i []uint8, elms ...interface{}) []uint8 {
	res, err := DeleteSliceElmsE(i, elms...)
	if err != nil {
		return nil
	}
	return res.([]uint8)
}

// DeleteUint8SliceElmsE deletes the specified elements from uint8 slice with error.
func DeleteUint8SliceElmsE(i []uint8, elms ...interface{}) ([]uint8, error) {
	res, err := DeleteSliceElmsE(i, elms...)
	if err != nil {
		return nil, err
	}
	return res.([]uint8), nil
}

// DeleteUint16SliceElms deletes the specified elements from uint16 slice.
func DeleteUint16SliceElms(i []uint16, elms ...interface{}) []uint16 {
	res, err := DeleteSliceElmsE(i, elms...)
	if err != nil {
		return nil
	}
	return res.([]uint16)
}

// DeleteUint16SliceElmsE deletes the specified elements from uint16 slice with error.
func DeleteUint16SliceElmsE(i []uint16, elms ...interface{}) ([]uint16, error) {
	res, err := DeleteSliceElmsE(i, elms...)
	if err != nil {
		return nil, err
	}
	return res.([]uint16), nil
}

// DeleteUint32SliceElms deletes the specified elements from uint32 slice.
func DeleteUint32SliceElms(i []uint16, elms ...interface{}) []uint32 {
	res, err := DeleteSliceElmsE(i, elms...)
	if err != nil {
		return nil
	}
	return res.([]uint32)
}

// DeleteUint32SliceElmsE deletes the specified elements from uint32 slice with error.
func DeleteUint32SliceElmsE(i []uint16, elms ...interface{}) ([]uint32, error) {
	res, err := DeleteSliceElmsE(i, elms...)
	if err != nil {
		return nil, err
	}
	return res.([]uint32), nil
}

// DeleteUint64SliceElms deletes the specified elements from uint64 slice.
func DeleteUint64SliceElms(i []uint16, elms ...interface{}) []uint64 {
	res, err := DeleteSliceElmsE(i, elms...)
	if err != nil {
		return nil
	}
	return res.([]uint64)
}

// DeleteUint64SliceElmsE deletes the specified elements from uint64 slice with error.
func DeleteUint64SliceElmsE(i []uint16, elms ...interface{}) ([]uint64, error) {
	res, err := DeleteSliceElmsE(i, elms...)
	if err != nil {
		return nil, err
	}
	return res.([]uint64), nil
}

// DeleteSliceElmsE deletes the specified elements from the slice.
// Note that the original slice will not be modified.
func DeleteSliceElmsE(i interface{}, elms ...interface{})(interface{}, error) {
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
