package slice

import (
	"errors"
	"reflect"
)

//
// Note that since Go 1.18, standard exp lib function https://pkg.go.dev/golang.org/x/exp/slices#Delete
// implemented by generics should be prefered.
//

// DeleteStrSlice deletes string slice elements by indexes.
func DeleteStrSlice(src []string, indexes ...int) []string {
	return Delete(src, indexes...).([]string)
}

// DeleteIntSlice deletes int slice elements by indexes.
func DeleteIntSlice(src []int, indexes ...int) []int {
	return Delete(src, indexes...).([]int)
}

// DeleteInt8Slice deletes int8 slice elements by indexes.
func DeleteInt8Slice(src []int8, indexes ...int) []int8 {
	return Delete(src, indexes...).([]int8)
}

// DeleteInt16Slice deletes int16 slice elements by indexes.
func DeleteInt16Slice(src []int16, indexes ...int) []int16 {
	return Delete(src, indexes...).([]int16)
}

// DeleteInt32Slice deletes int32 slice elements by indexes.
func DeleteInt32Slice(src []int32, indexes ...int) []int32 {
	return Delete(src, indexes...).([]int32)
}

// DeleteInt64Slice deletes int64 slice elements by indexes.
func DeleteInt64Slice(src []int64, indexes ...int) []int64 {
	return Delete(src, indexes...).([]int64)
}

// DeleteUintSlice deletes uint slice elements by indexes.
func DeleteUintSlice(src []uint, indexes ...int) []uint {
	return Delete(src, indexes...).([]uint)
}

// DeleteUint8Slice deletes uint8 slice elements by indexes.
func DeleteUint8Slice(src []uint8, indexes ...int) []uint8 {
	return Delete(src, indexes...).([]uint8)
}

// DeleteUint16Slice deletes uint16 slice elements by indexes.
func DeleteUint16Slice(src []uint16, indexes ...int) []uint16 {
	return Delete(src, indexes...).([]uint16)
}

// DeleteUint32Slice deletes uint32 slice elements by indexes.
func DeleteUint32Slice(src []uint32, indexes ...int) []uint32 {
	return Delete(src, indexes...).([]uint32)
}

// DeleteUint64Slice deletes uint64 slice elements by indexes.
func DeleteUint64Slice(src []uint64, indexes ...int) []uint64 {
	return Delete(src, indexes...).([]uint64)
}

// DeleteElems deletes the specified elements from the slice.
// Note that the original slice will not be modified.
func DeleteElems(i any, elms ...any) any {
	r, _ := DeleteElemsE(i, elms...)
	return r
}

// DeleteElemsE deletes the specified elements from the slice.
// Note that the original slice will not be modified.
func DeleteElemsE(i any, elms ...any) (any, error) {
	// Check params.
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
	// Convert the elements to map set.
	m := make(map[any]struct{})
	for _, v := range elms {
		m[v] = struct{}{}
	}
	// Filter out specified elements.
	t := reflect.MakeSlice(reflect.TypeOf(i), 0, v.Len())
	for i := 0; i < v.Len(); i++ {
		if _, ok := m[v.Index(i).Interface()]; !ok {
			t = reflect.Append(t, v.Index(i))
		}
	}
	return t.Interface(), nil
}

// Delete deletes the specified index element from the slice.
// Note that the original slice will not be modified.
func Delete(slice any, indexes ...int) any {
	res, _ := DeleteE(slice, indexes...)
	return res
}

// DeleteE deletes the element specified by index from the slice with returned error.
// Note that the original slice will not be modified.
func DeleteE(slice any, indexes ...int) (any, error) {
	// Check params.
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, errors.New("the input isn't a slice")
	}
	if v.Len() == 0 || len(indexes) == 0 {
		return slice, nil
	}
	// Convert the indexes to map set.
	m := make(map[int]struct{})
	for _, i := range indexes {
		m[i] = struct{}{}
	}
	// Delete.
	t := reflect.MakeSlice(reflect.TypeOf(slice), 0, v.Len())
	for i := 0; i < v.Len(); i++ {
		if _, ok := m[i]; !ok {
			t = reflect.Append(t, v.Index(i))
		}
	}
	return t.Interface(), nil
}
