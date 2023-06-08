package slice

import (
	"errors"
	"fmt"
	"reflect"
)

//
// Note that since Go 1.18, the standard exp lib function https://pkg.go.dev/golang.org/x/exp/slices#Replace
// implemented by generics should be prefered.
//

func UpdateInt(src []int, index, value int) []int {
	tmp, _ := UpdateE(src, index, value)
	v, _ := tmp.([]int)
	return v
}

func UpdateInt8(src []int8, index int, value int8) []int8 {
	tmp, _ := UpdateE(src, index, value)
	v, _ := tmp.([]int8)
	return v
}

func UpdateInt16(src []int16, index int, value int16) []int16 {
	tmp, _ := UpdateE(src, index, value)
	v, _ := tmp.([]int16)
	return v
}

func UpdateInt32(src []int32, index int, value int32) []int32 {
	tmp, _ := UpdateE(src, index, value)
	v, _ := tmp.([]int32)
	return v
}

func UpdateInt64(src []int64, index int, value int64) []int64 {
	tmp, _ := UpdateE(src, index, value)
	v, _ := tmp.([]int64)
	return v
}

func UpdateUint(src []uint, index int, value uint) []uint {
	tmp, _ := UpdateE(src, index, value)
	v, _ := tmp.([]uint)
	return v
}

func UpdateUint8(src []uint8, index int, value uint8) []uint8 {
	tmp, _ := UpdateE(src, index, value)
	v, _ := tmp.([]uint8)
	return v
}

func UpdateUint16(src []uint16, index int, value uint16) []uint16 {
	tmp, _ := UpdateE(src, index, value)
	v, _ := tmp.([]uint16)
	return v
}

func UpdateUint32(src []uint32, index int, value uint32) []uint32 {
	tmp, _ := UpdateE(src, index, value)
	v, _ := tmp.([]uint32)
	return v
}

func UpdateUint64(src []uint64, index int, value uint64) []uint64 {
	tmp, _ := UpdateE(src, index, value)
	v, _ := tmp.([]uint64)
	return v
}

func UpdateStr(src []string, index int, value string) []string {
	tmp, _ := UpdateE(src, index, value)
	v, _ := tmp.([]string)
	return v
}

// UpdateE modifies the specified index element of slice with returned error.
// Note that the original slice will not be modified.
func UpdateE(a any, index int, value any) (any, error) {
	// Check params.
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("the input %#v of type %T isn't a slice", a, a)
	}
	if index > v.Len()-1 {
		return nil, errors.New("index is out of range")
	}
	if reflect.TypeOf(a).Elem() != reflect.TypeOf(value) {
		return nil, errors.New("value type is inconsistent with the slice element type")
	}

	// Update slice.
	r := reflect.MakeSlice(reflect.TypeOf(a), 0, 0)
	r = reflect.AppendSlice(r, v.Slice(0, v.Len()))
	r.Index(index).Set(reflect.ValueOf(value))
	return r.Interface(), nil
}
