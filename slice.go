package util

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/spf13/cast"
)

//
// Part 1: unique a slice, e.g. input []int32{1, 2, 2, 3} and output is []int32{1, 2, 3}.
//

func UniqueIntSlice(src []int) []int {
	dst, _ := UniqueSliceE(src)
	v, _ := dst.([]int)
	return v
}

func UniqueInt8Slice(src []int8) []int8 {
	dst, _ := UniqueSliceE(src)
	v, _ := dst.([]int8)
	return v
}

func UniqueInt16Slice(src []int16) []int16 {
	dst, _ := UniqueSliceE(src)
	v, _ := dst.([]int16)
	return v
}

func UniqueInt32Slice(src []int32) []int32 {
	dst, _ := UniqueSliceE(src)
	v, _ := dst.([]int32)
	return v
}

func UniqueInt64Slice(src []int64) []int64 {
	dst, _ := UniqueSliceE(src)
	v, _ := dst.([]int64)
	return v
}

func UniqueUintSlice(src []uint) []uint {
	dst, _ := UniqueSliceE(src)
	v, _ := dst.([]uint)
	return v
}

func UniqueUint8Slice(src []uint8) []uint8 {
	dst, _ := UniqueSliceE(src)
	v, _ := dst.([]uint8)
	return v
}

func UniqueUint16Slice(src []uint16) []uint16 {
	dst, _ := UniqueSliceE(src)
	v, _ := dst.([]uint16)
	return v
}

func UniqueUint32Slice(src []uint32) []uint32 {
	dst, _ := UniqueSliceE(src)
	v, _ := dst.([]uint32)
	return v
}

func UniqueUint64Slice(src []uint64) []uint64 {
	dst, _ := UniqueSliceE(src)
	v, _ := dst.([]uint64)
	return v
}

func UniqueFloat32Slice(src []float32) []float32 {
	dst, _ := UniqueSliceE(src)
	v, _ := dst.([]float32)
	return v
}

func UniqueFloat64Slice(src []float64) []float64 {
	dst, _ := UniqueSliceE(src)
	v, _ := dst.([]float64)
	return v
}

func UniqueStrSlice(src []string) []string {
	dst, _ := UniqueSliceE(src)
	v, _ := dst.([]string)
	return v
}

//
// Part 2: reverse a slice, e.g. input []int32{1, 2, 3} and output is []int32{3, 2, 1}.
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

//
// Part 3: sum a slice, e.g. input []int32{1, 2, 3} and output is 6.
//

// SumSlice calculates the sum of slice elements
func SumSlice(slice interface{}) float64 {
	v, _ := SumSliceE(slice)
	return v
}

//
// Part 4: determine whether the slice contains an element.
//

// IsContains checks whether slice or array contains the target element.
// Note that if the target element is a numeric literal, please specify its type explicitly, otherwise it defaults to int.
// For example you might call like IsContains([]int32{1,2,3}, int32(1)).
func IsContains(i interface{}, target interface{}) bool {
	if i == nil {
		return false
	}
	t := reflect.TypeOf(i)
	if t.Kind() != reflect.Slice && t.Kind() != reflect.Array {
		return false
	}
	v := reflect.ValueOf(i)
	for i := 0; i < v.Len(); i++ {
		if target == v.Index(i).Interface() {
			return true
		}
	}
	return false
}

//
// Part 5: join slice with a separator. e.g. input []int32{1, 2, 3} and separator "#", output is a string "1#2#3"
//

// JoinSliceWithSep joins all elements in slice with a separator
func JoinSliceWithSep(slice interface{}, sep string) string {
	s, _ := JoinSliceWithSepE(slice, sep)
	return s
}

//
// Part 6: CRUD(Create Read Update Delete) on slice by index.
//

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

func GetEleIndexesSlice(slice interface{}, value interface{}) []int {
	indexes, _ := GetEleIndexesSliceE(slice, value)
	return indexes
}

//
// Part 7: get the min or max element of a slice.
//

func MinIntSlice(sl []int) int {
	min, _ := MinSliceE(sl)
	v, _ := min.(int)
	return v
}

func MinInt8Slice(sl []int8) int8 {
	min, _ := MinSliceE(sl)
	v, _ := min.(int8)
	return v
}

func MinInt16Slice(sl []int16) int16 {
	min, _ := MinSliceE(sl)
	v, _ := min.(int16)
	return v
}

func MinInt32Slice(sl []int32) int32 {
	min, _ := MinSliceE(sl)
	v, _ := min.(int32)
	return v
}

func MinInt64Slice(sl []int64) int64 {
	min, _ := MinSliceE(sl)
	v, _ := min.(int64)
	return v
}

func MinUintSlice(sl []uint) uint {
	min, _ := MinSliceE(sl)
	v, _ := min.(uint)
	return v
}

func MinUint8Slice(sl []uint8) uint8 {
	min, _ := MinSliceE(sl)
	v, _ := min.(uint8)
	return v
}

func MinUint16Slice(sl []uint16) uint16 {
	min, _ := MinSliceE(sl)
	v, _ := min.(uint16)
	return v
}

func MinUint32Slice(sl []uint32) uint32 {
	min, _ := MinSliceE(sl)
	v, _ := min.(uint32)
	return v
}

func MinUint64Slice(sl []uint64) uint64 {
	min, _ := MinSliceE(sl)
	v, _ := min.(uint64)
	return v
}

func MinFloat32Slice(sl []float32) float32 {
	min, _ := MinSliceE(sl)
	v, _ := min.(float32)
	return v
}

func MinFloat64Slice(sl []float64) float64 {
	min, _ := MinSliceE(sl)
	v, _ := min.(float64)
	return v
}

func MaxIntSlice(sl []int) int {
	max, _ := MaxSliceE(sl)
	v, _ := max.(int)
	return v
}

func MaxInt8Slice(sl []int8) int8 {
	max, _ := MaxSliceE(sl)
	v, _ := max.(int8)
	return v
}

func MaxInt16Slice(sl []int16) int16 {
	max, _ := MaxSliceE(sl)
	v, _ := max.(int16)
	return v
}

func MaxInt32Slice(sl []int32) int32 {
	max, _ := MaxSliceE(sl)
	v, _ := max.(int32)
	return v
}

func MaxInt64Slice(sl []int64) int64 {
	max, _ := MaxSliceE(sl)
	v, _ := max.(int64)
	return v
}

func MaxUintSl(sl []uint) uint {
	max, _ := MaxSliceE(sl)
	v, _ := max.(uint)
	return v
}

func MaxUint8Slice(sl []uint8) uint8 {
	max, _ := MaxSliceE(sl)
	v, _ := max.(uint8)
	return v
}

func MaxUint16Slice(sl []uint16) uint16 {
	max, _ := MaxSliceE(sl)
	v, _ := max.(uint16)
	return v
}

func MaxUint32Slice(sl []uint32) uint32 {
	max, _ := MaxSliceE(sl)
	v, _ := max.(uint32)
	return v
}

func MaxUint64Slice(sl []uint64) uint64 {
	max, _ := MaxSliceE(sl)
	v, _ := max.(uint64)
	return v
}

func MaxFloat32Slice(sl []float32) float32 {
	max, _ := MaxSliceE(sl)
	v, _ := max.(float32)
	return v
}

func MaxFloat64Slice(sl []float64) float64 {
	max, _ := MaxSliceE(sl)
	v, _ := max.(float64)
	return v
}

//
// Part 8: get a random element from a slice or array.
//

// GetRandomSliceElem get a random element from a slice or array.
// If the length of slice or array is zero it will panic.
func GetRandomSliceElem(i interface{}) interface{} {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return i
	}
	idx := RandIntn(v.Len())
	return v.Index(idx).Interface()
}

//
// Part x: basic operating functions of slice.
//

// UniqueSliceE deletes repeated elements in a slice with error.
// Note that the original slice will not be modified.
func UniqueSliceE(slice interface{}) (interface{}, error) {
	// check params
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("the input %#v of type %T isn't a slice", slice, slice)
	}
	// unique the slice
	dst := reflect.MakeSlice(reflect.TypeOf(slice), 0, v.Len())
	m := make(map[interface{}]struct{})
	for i := 0; i < v.Len(); i++ {
		if _, ok := m[v.Index(i).Interface()]; !ok {
			dst = reflect.Append(dst, v.Index(i))
			m[v.Index(i).Interface()] = struct{}{}
		}
	}
	return dst.Interface(), nil
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

// SumSliceE returns the sum of slice elements and an error if occurred.
func SumSliceE(slice interface{}) (float64, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return 0.0, fmt.Errorf("the input %#v of type %T isn't a slice", slice, slice)
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
		case int64:
			sum += float64(v)
		case uint:
			sum += float64(v)
		case uint8:
			sum += float64(v)
		case uint16:
			sum += float64(v)
		case uint32:
			sum += float64(v)
		case uint64:
			sum += float64(v)
		case float32:
			sum += float64(v)
		case float64:
			sum += v
		default:
			return 0.0, fmt.Errorf("the element %#v of slice type %T isn't numerical type", v, v)
		}
	}
	return sum, nil
}

// MinSliceE returns the smallest element of the slice and an error if occurred.
// If slice length is zero return the zero value of the element type.
func MinSliceE(slice interface{}) (interface{}, error) {
	// check params
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("the input %#v of type %T isn't a slice", slice, slice)
	}
	if v.Len() == 0 {
		return nil, nil
	}
	// get the min element
	min := v.Index(0).Interface()
	for i := 1; i < v.Len(); i++ {
		switch v := v.Index(i).Interface().(type) {
		case int:
			if v < min.(int) {
				min = v
			}
		case int8:
			if v < min.(int8) {
				min = v
			}
		case int16:
			if v < min.(int16) {
				min = v
			}
		case int32:
			if v < min.(int32) {
				min = v
			}
		case int64:
			if v < min.(int64) {
				min = v
			}
		case uint:
			if v < min.(uint) {
				min = v
			}
		case uint8:
			if v < min.(uint8) {
				min = v
			}
		case uint16:
			if v < min.(uint16) {
				min = v
			}
		case uint32:
			if v < min.(uint32) {
				min = v
			}
		case uint64:
			if v < min.(uint64) {
				min = v
			}
		case float32:
			if v < min.(float32) {
				min = v
			}
		case float64:
			if v < min.(float64) {
				min = v
			}
		default:
			return nil, fmt.Errorf("the element %#v of slice type %T isn't numerical type", v, v)
		}
	}
	return min, nil
}

// MaxSliceE returns the largest element of the slice and an error if occurred.
// If slice length is zero return the zero value of the element type.
func MaxSliceE(slice interface{}) (interface{}, error) {
	// check params
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("the input %#v of type %T isn't a slice", slice, slice)
	}
	if v.Len() == 0 {
		return nil, nil
	}
	// get the max element.
	max := v.Index(0).Interface()
	for i := 1; i < v.Len(); i++ {
		switch v := v.Index(i).Interface().(type) {
		case int:
			if v > max.(int) {
				max = v
			}
		case int8:
			if v > max.(int8) {
				max = v
			}
		case int16:
			if v > max.(int16) {
				max = v
			}
		case int32:
			if v > max.(int32) {
				max = v
			}
		case int64:
			if v > max.(int64) {
				max = v
			}
		case uint:
			if v > max.(uint) {
				max = v
			}
		case uint8:
			if v > max.(uint8) {
				max = v
			}
		case uint16:
			if v > max.(uint16) {
				max = v
			}
		case uint32:
			if v > max.(uint32) {
				max = v
			}
		case uint64:
			if v > max.(uint64) {
				max = v
			}
		case float32:
			if v > max.(float32) {
				max = v
			}
		case float64:
			if v > max.(float64) {
				max = v
			}
		default:
			return nil, fmt.Errorf("the element %#v of slice type %T isn't numerical type", v, v)
		}
	}
	return max, nil
}

// JoinSliceWithSepE joins all elements in slice or array with separator and return an error if occurred.
func JoinSliceWithSepE(slice interface{}, sep string) (string, error) {
	// check params
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return "", fmt.Errorf("the input %#v of type %T isn't a slice or array", slice, slice)
	}
	// join the slice or array
	var s string
	for i := 0; i < v.Len(); i++ {
		if len(s) > 0 {
			s += sep
		}
		str, err := cast.ToStringE(v.Index(i).Interface())
		if err != nil {
			return "", err
		}
		s += str
	}
	return s, nil
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

// UpdateSliceE modifies the specified index element of slice.
// Note that the original slice will not be modified.
func UpdateSliceE(slice interface{}, index int, value interface{}) (interface{}, error) {
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

// GetEleIndexesSliceE finds all indexes of the specified element in a slice.
func GetEleIndexesSliceE(slice interface{}, value interface{}) ([]int, error) {
	// check params
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("the input %#v of type %T isn't a slice", slice, slice)
	}
	// get indexes
	var indexes []int
	for i := 0; i < v.Len(); i++ {
		if v.Index(i).Interface() == value {
			indexes = append(indexes, i)
		}
	}
	return indexes, nil
}
