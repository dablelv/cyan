package util

import (
	"errors"
	"reflect"
)

func UniqueIntSlice(src []int) []int {
	tmp, _ := UniqueSlice(src)
	dst := []int{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToInt(v))
	}
	return dst
}

func UniqueInt8Slice(src []int8) []int8 {
	tmp, _ := UniqueSlice(src)
	dst := []int8{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToInt8(v))
	}
	return dst
}

func UniqueInt16Slice(src []int16) []int16 {
	tmp, _ := UniqueSlice(src)
	dst := []int16{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToInt16(v))
	}
	return dst
}

func UniqueInt32Slice(src []int32) []int32 {
	tmp, _ := UniqueSlice(src)
	dst := []int32{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToInt32(v))
	}
	return dst
}

func UniqueInt64Slice(src []int64) []int64 {
	tmp, _ := UniqueSlice(src)
	dst := []int64{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToInt64(v))
	}
	return dst
}

func UniqueUintSlice(src []uint) []uint {
	tmp, _ := UniqueSlice(src)
	dst := []uint{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToUint(v))
	}
	return dst
}

func UniqueUint8Slice(src []uint8) []uint8 {
	tmp, _ := UniqueSlice(src)
	dst := []uint8{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToUint8(v))
	}
	return dst
}

func UniqueUint16Slice(src []uint16) []uint16 {
	tmp, _ := UniqueSlice(src)
	dst := []uint16{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToUint16(v))
	}
	return dst
}

func UniqueUint32Slice(src []uint32) []uint32 {
	tmp, _ := UniqueSlice(src)
	dst := []uint32{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToUint32(v))
	}
	return dst
}

func UniqueUint64Slice(src []uint64) []uint64 {
	tmp, _ := UniqueSlice(src)
	dst := []uint64{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToUint64(v))
	}
	return dst
}

func UniqueStrSlice(src []string) []string {
	tmp, _ := UniqueSlice(src)
	dst := []string{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToString(v))
	}
	return dst
}

func ReverseIntSlice(src []int) []int {
	tmp, _ := ReverseSlice(src)
	dst := []int{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToInt(v))
	}
	return dst
}

func ReverseInt8Slice(src []int8) []int8 {
	tmp, _ := ReverseSlice(src)
	dst := []int8{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToInt8(v))
	}
	return dst
}

func ReverseInt16Slice(src []int16) []int16 {
	tmp, _ := ReverseSlice(src)
	dst := []int16{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToInt16(v))
	}
	return dst
}

func ReverseInt32Slice(src []int32) []int32 {
	tmp, _ := ReverseSlice(src)
	dst := []int32{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToInt32(v))
	}
	return dst
}

func ReverseInt64Slice(src []int64) []int64 {
	tmp, _ := ReverseSlice(src)
	dst := []int64{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToInt64(v))
	}
	return dst
}

func ReverseUintSlice(src []uint) []uint {
	tmp, _ := ReverseSlice(src)
	dst := []uint{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToUint(v))
	}
	return dst
}

func ReverseUint8Slice(src []uint8) []uint8 {
	tmp, _ := ReverseSlice(src)
	dst := []uint8{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToUint8(v))
	}
	return dst
}

func ReverseUint16Slice(src []uint16) []uint16 {
	tmp, _ := ReverseSlice(src)
	dst := []uint16{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToUint16(v))
	}
	return dst
}

func ReverseUint32Slice(src []uint32) []uint32 {
	tmp, _ := ReverseSlice(src)
	dst := []uint32{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToUint32(v))
	}
	return dst
}

func ReverseUint64Slice(src []uint64) []uint64 {
	tmp, _ := ReverseSlice(src)
	dst := []uint64{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToUint64(v))
	}
	return dst
}

func ReverseByteSlice(src []byte) []byte {
	tmp, _ := ReverseSlice(src)
	dst := []byte{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, byte(ToUint8(v)))
	}
	return dst
}

func ReverseStrSlice(src []string) []string {
	tmp, _ := ReverseSlice(src)
	dst := []string{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToString(v))
	}
	return dst
}

// IsContains check slice whether contain target element
// note: if the target element is a numeric literal, please specify its type explicitly, otherwise it defaults to int.
// You might call IsContains like IsContains([]int32{1,2,3}, int32(1))
func IsContains(list interface{}, target interface{}) bool {
	if reflect.TypeOf(list).Kind() == reflect.Slice {
		v := reflect.ValueOf(list)
		for i := 0; i < v.Len(); i++ {
			if target == v.Index(i).Interface() {
				return true
			}
		}
	}
	return false
}

// UniqueSlice delete repeated elements in a slice
func UniqueSlice(slice interface{}) (interface{}, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, errors.New("param isn't a slice")
	}

	var dst []interface{}
	m := make(map[interface{}]bool)
	for i := 0; i < v.Len(); i++ {
		if _, ok := m[v.Index(i).Interface()]; !ok {
			dst = append(dst, v.Index(i).Interface())
			m[v.Index(i).Interface()] = true
		}
	}
	return dst, nil
}

// ReverseSlice reverses the specified slice without modifying the original slice
func ReverseSlice(slice interface{}) (interface{}, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, errors.New("param isn't a slice")
	}

	var dst []interface{}
	for i := v.Len() - 1; i >= 0; i-- {
		dst = append(dst, v.Index(i).Interface())
	}
	return dst, nil
}

// SumSlice calculates the sum of slice elements
func SumSlice(slice interface{}) float64 {
	v, _ := SumSliceE(slice)
	return v
}

// SumSliceE calculates the sum of slice elements and return has an error param
func SumSliceE(slice interface{}) (float64, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return 0.0, errors.New("param isn't a slice")
	}

	var sum float64
	for i := 0; i < v.Len(); i++ {
		switch v := v.Index(i).Interface().(type) {
		case int:
			sum += float64(v)
		case int8:
			sum += float64(v)
		case int16:
			sum += float64(v)
		case int32:
			sum += float64(v)
		case uint:
			sum += float64(v)
		case uint8:
			sum += float64(v)
		case uint16:
			sum += float64(v)
		case uint32:
			sum += float64(v)
		case float32:
			sum += float64(v)
		case float64:
			sum += v
		default:
			return 0.0, errors.New("elements in slice aren't numerical type")
		}
	}
	return sum, nil
}

// JoinSliceWithSep joins all elements in slice with separator
func JoinSliceWithSep(slice interface{}, sep string) string {
	s, _ := JoinSliceWithSepE(slice, sep)
	return s
}

// JoinSliceWithSepE joins all elements in slice with separator and return has an error param
func JoinSliceWithSepE(slice interface{}, sep string) (string, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return "", errors.New("param isn't a slice")
	}

	var s string
	for i := 0; i < v.Len(); i++ {
		if len(s) > 0 {
			s += sep
		}
		str, err := ToStringE(v.Index(i).Interface())
		if err != nil {
			return "", err
		}
		s += str
	}
	return s, nil
}
