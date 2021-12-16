package util

import (
	"errors"
	"reflect"

	"github.com/spf13/cast"
)

//
// part 0: unique a slice, e.g. input []int32{1, 2, 2, 3} and output is []int32{1, 2, 3}
//

func UniqueIntSlice(src []int) []int {
	tmp, _ := UniqueSliceE(src)
	var dst []int
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, cast.ToInt(v))
	}
	return dst
}

func UniqueInt8Slice(src []int8) []int8 {
	tmp, _ := UniqueSliceE(src)
	var dst []int8
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, cast.ToInt8(v))
	}
	return dst
}

func UniqueInt16Slice(src []int16) []int16 {
	tmp, _ := UniqueSliceE(src)
	var dst []int16
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, cast.ToInt16(v))
	}
	return dst
}

func UniqueInt32Slice(src []int32) []int32 {
	tmp, _ := UniqueSliceE(src)
	var dst []int32
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, cast.ToInt32(v))
	}
	return dst
}

func UniqueInt64Slice(src []int64) []int64 {
	tmp, _ := UniqueSliceE(src)
	var dst []int64
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, cast.ToInt64(v))
	}
	return dst
}

func UniqueUintSlice(src []uint) []uint {
	tmp, _ := UniqueSliceE(src)
	var dst []uint
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, cast.ToUint(v))
	}
	return dst
}

func UniqueUint8Slice(src []uint8) []uint8 {
	tmp, _ := UniqueSliceE(src)
	var dst []uint8
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, cast.ToUint8(v))
	}
	return dst
}

func UniqueUint16Slice(src []uint16) []uint16 {
	tmp, _ := UniqueSliceE(src)
	var dst []uint16
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, cast.ToUint16(v))
	}
	return dst
}

func UniqueUint32Slice(src []uint32) []uint32 {
	tmp, _ := UniqueSliceE(src)
	var dst []uint32
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, cast.ToUint32(v))
	}
	return dst
}

func UniqueUint64Slice(src []uint64) []uint64 {
	tmp, _ := UniqueSliceE(src)
	var dst []uint64
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, cast.ToUint64(v))
	}
	return dst
}

func UniqueFloat32Slice(src []uint32) []float32 {
	tmp, _ := UniqueSliceE(src)
	var dst []float32
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, cast.ToFloat32(v))
	}
	return dst
}

func UniqueFloat64Slice(src []uint64) []float64 {
	tmp, _ := UniqueSliceE(src)
	var dst []float64
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, cast.ToFloat64(v))
	}
	return dst
}

func UniqueStrSlice(src []string) []string {
	tmp, _ := UniqueSliceE(src)
	var dst []string
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, cast.ToString(v))
	}
	return dst
}

//
// part 1: reverse a slice, e.g. input []int32{1, 2, 3} and output is []int32{3, 2, 1}
//

func ReverseIntSlice(src []int) []int {
	tmp, _ := ReverseSliceE(src)
	dst := []int{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, cast.ToInt(v))
	}
	return dst
}

func ReverseInt8Slice(src []int8) []int8 {
	tmp, _ := ReverseSliceE(src)
	dst := []int8{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, cast.ToInt8(v))
	}
	return dst
}

func ReverseInt16Slice(src []int16) []int16 {
	tmp, _ := ReverseSliceE(src)
	dst := []int16{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, cast.ToInt16(v))
	}
	return dst
}

func ReverseInt32Slice(src []int32) []int32 {
	tmp, _ := ReverseSliceE(src)
	dst := []int32{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, cast.ToInt32(v))
	}
	return dst
}

func ReverseInt64Slice(src []int64) []int64 {
	tmp, _ := ReverseSliceE(src)
	dst := []int64{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, cast.ToInt64(v))
	}
	return dst
}

func ReverseUintSlice(src []uint) []uint {
	tmp, _ := ReverseSliceE(src)
	dst := []uint{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, cast.ToUint(v))
	}
	return dst
}

func ReverseUint8Slice(src []uint8) []uint8 {
	tmp, _ := ReverseSliceE(src)
	dst := []uint8{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, cast.ToUint8(v))
	}
	return dst
}

func ReverseUint16Slice(src []uint16) []uint16 {
	tmp, _ := ReverseSliceE(src)
	dst := []uint16{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, cast.ToUint16(v))
	}
	return dst
}

func ReverseUint32Slice(src []uint32) []uint32 {
	tmp, _ := ReverseSliceE(src)
	dst := []uint32{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, cast.ToUint32(v))
	}
	return dst
}

func ReverseUint64Slice(src []uint64) []uint64 {
	tmp, _ := ReverseSliceE(src)
	dst := []uint64{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, cast.ToUint64(v))
	}
	return dst
}

func ReverseStrSlice(src []string) []string {
	tmp, _ := ReverseSliceE(src)
	dst := []string{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, cast.ToString(v))
	}
	return dst
}

//
// part 2: sum a slice, e.g. input []int32{1, 2, 3} and output is 6.
//

// SumSlice calculates the sum of slice elements
func SumSlice(slice interface{}) float64 {
	v, _ := SumSliceE(slice)
	return v
}

//
// part3: determine whether the slice contains an element.
//

// IsContains check slice whether contain target element.
// note: if the target element is a numeric literal, please specify its type explicitly, otherwise it defaults to int.
// You might call IsContains like IsContains([]int32{1,2,3}, int32(1)).
func IsContains(slice interface{}, target interface{}) bool {
	if reflect.TypeOf(slice).Kind() == reflect.Slice {
		v := reflect.ValueOf(slice)
		for i := 0; i < v.Len(); i++ {
			if target == v.Index(i).Interface() {
				return true
			}
		}
	}
	return false
}

//
// part 4: join slice with a separator. e.g. input []int32{1, 2, 3} and separator "#", output is a string "1#2#3"
//

// JoinSliceWithSep joins all elements in slice with a separator
func JoinSliceWithSep(slice interface{}, sep string) string {
	s, _ := JoinSliceWithSepE(slice, sep)
	return s
}

//
// part 5: CRUD(Create Read Update Delete) on slice by index.
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

func DeleteIntSlice(src []int, index int) []int {
	tmp, _ := DeleteSliceE(src, index)
	v, _ := tmp.([]int)
	return v
}

func DeleteInt8Slice(src []int8, index int) []int8 {
	tmp, _ := DeleteSliceE(src, index)
	v, _ := tmp.([]int8)
	return v
}

func DeleteInt16Slice(src []int, index int) []int16 {
	tmp, _ := DeleteSliceE(src, index)
	v, _ := tmp.([]int16)
	return v
}

func DeleteInt32Slice(src []int, index int) []int32 {
	tmp, _ := DeleteSliceE(src, index)
	v, _ := tmp.([]int32)
	return v
}

func DeleteInt64Slice(src []int, index int) []int64 {
	tmp, _ := DeleteSliceE(src, index)
	v, _ := tmp.([]int64)
	return v
}

func DeleteUintSlice(src []int, index int) []uint {
	tmp, _ := DeleteSliceE(src, index)
	v, _ := tmp.([]uint)
	return v
}

func DeleteUint8Slice(src []int8, index int) []uint8 {
	tmp, _ := DeleteSliceE(src, index)
	v, _ := tmp.([]uint8)
	return v
}

func DeleteUint16Slice(src []int, index int) []uint16 {
	tmp, _ := DeleteSliceE(src, index)
	v, _ := tmp.([]uint16)
	return v
}

func DeleteUint32Slice(src []int, index int) []uint32 {
	tmp, _ := DeleteSliceE(src, index)
	v, _ := tmp.([]uint32)
	return v
}

func DeleteUint64Slice(src []int, index int) []uint64 {
	tmp, _ := DeleteSliceE(src, index)
	v, _ := tmp.([]uint64)
	return v
}

func DeleteStrSlice(src []int, index int) []string {
	tmp, _ := DeleteSliceE(src, index)
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
// part 6: get the min or max element of a slice.
//

func MinIntSl(sl []int) int {
	min, _ := MinSliceE(sl)
	return cast.ToInt(min)
}

func MinInt8Sl(sl []int8) int8 {
	min, _ := MinSliceE(sl)
	return cast.ToInt8(min)
}

func MinInt16Sl(sl []int16) int16 {
	min, _ := MinSliceE(sl)
	return cast.ToInt16(min)
}

func MinInt32Sl(sl []int32) int32 {
	min, _ := MinSliceE(sl)
	return cast.ToInt32(min)
}

func MinInt64Sl(sl []int64) int64 {
	min, _ := MinSliceE(sl)
	return cast.ToInt64(min)
}

func MinUintSl(sl []uint) uint {
	min, _ := MinSliceE(sl)
	return cast.ToUint(min)
}

func MinUint8Sl(sl []uint8) uint8 {
	min, _ := MinSliceE(sl)
	return cast.ToUint8(min)
}

func MinUint16Sl(sl []uint16) uint16 {
	min, _ := MinSliceE(sl)
	return cast.ToUint16(min)
}

func MinUint32Sl(sl []uint32) uint32 {
	min, _ := MinSliceE(sl)
	return cast.ToUint32(min)
}

func MinUint64Sl(sl []uint64) uint64 {
	min, _ := MinSliceE(sl)
	return cast.ToUint64(min)
}

func MinFloat32Sl(sl []float32) float32 {
	min, _ := MinSliceE(sl)
	return cast.ToFloat32(min)
}

func MinFloat64Sl(sl []float64) float64 {
	min, _ := MinSliceE(sl)
	return cast.ToFloat64(min)
}

func MaxIntSl(sl []int) int {
	max, _ := MaxSliceE(sl)
	return cast.ToInt(max)
}

func MaxInt8Sl(sl []int8) int8 {
	max, _ := MaxSliceE(sl)
	return cast.ToInt8(max)
}

func MaxInt16Sl(sl []int16) int16 {
	min, _ := MaxSliceE(sl)
	return cast.ToInt16(min)
}

func MaxInt32Sl(sl []int32) int32 {
	max, _ := MaxSliceE(sl)
	return cast.ToInt32(max)
}

func MaxInt64Sl(sl []int64) int64 {
	max, _ := MaxSliceE(sl)
	return cast.ToInt64(max)
}

func MaxUintSl(sl []uint) uint {
	max, _ := MaxSliceE(sl)
	return cast.ToUint(max)
}

func MaxUint8Sl(sl []uint8) uint8 {
	max, _ := MaxSliceE(sl)
	return cast.ToUint8(max)
}

func MaxUint16Sl(sl []uint16) uint16 {
	max, _ := MaxSliceE(sl)
	return cast.ToUint16(max)
}

func MaxUint32Sl(sl []uint32) uint32 {
	max, _ := MaxSliceE(sl)
	return cast.ToUint32(max)
}

func MaxUint64Sl(sl []uint64) uint64 {
	max, _ := MaxSliceE(sl)
	return cast.ToUint64(max)
}

func MaxFloat32Sl(sl []float32) float32 {
	max, _ := MaxSliceE(sl)
	return cast.ToFloat32(max)
}

func MaxFloat64Sl(sl []float64) float64 {
	max, _ := MaxSliceE(sl)
	return cast.ToFloat64(max)
}

//
// part x: basic operating functions of slice
//

// UniqueSliceE delete repeated elements in a slice with error
func UniqueSliceE(slice interface{}) (interface{}, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, errors.New("input isn't a slice")
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

// ReverseSliceE reverses the specified slice without modifying the original slice
func ReverseSliceE(slice interface{}) (interface{}, error) {
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

// SumSliceE returns the sum of slice elements and an error if occurred.
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
			return 0.0, errors.New("elements in slice aren't numerical type")
		}
	}
	return sum, nil
}

// MinSliceE returns the smallest element of the slice and an error if occurred.
// If slice length is zero return the zero value of the element type.
func MinSliceE(slice interface{}) (interface{}, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, errors.New("param isn't a slice")
	}
	if v.Len() == 0 {
		return nil, nil
	}

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
			return nil, errors.New("elements in slice aren't numerical type")
		}
	}
	return min, nil
}

// MaxSliceE returns the largest element of the slice and an error if occurred.
// If slice length is zero return the zero value of the element type.
func MaxSliceE(slice interface{}) (interface{}, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, errors.New("param isn't a slice")
	}
	if v.Len() == 0 {
		return nil, nil
	}

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
			return nil, errors.New("elements in slice aren't numerical type")
		}
	}
	return max, nil
}

// JoinSliceWithSepE joins all elements in slice with separator and return an error if occurred.
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
		return nil, errors.New("target isn't a slice")
	}
	if index < 0 || index > v.Len() || reflect.TypeOf(slice).Elem() != reflect.TypeOf(value) {
		return nil, errors.New("param is invalid")
	}

	t := reflect.MakeSlice(reflect.TypeOf(slice), 0, 0)

	// add the element to the end of slice
	if index == v.Len() {
		t = reflect.AppendSlice(t, v.Slice(0, v.Len()))
		t = reflect.Append(t, reflect.ValueOf(value))
		return t.Interface(), nil
	}

	t = reflect.AppendSlice(t, v.Slice(0, index+1))
	t = reflect.AppendSlice(t, v.Slice(index, v.Len()))
	t.Index(index).Set(reflect.ValueOf(value))
	return t.Interface(), nil
}

// DeleteSliceE deletes the specified index element from the slice.
// Note that the original slice will not be modified.
func DeleteSliceE(slice interface{}, index int) (interface{}, error) {
	// check params
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, errors.New("the input isn't a slice")
	}
	if v.Len() == 0 || index < 0 || index > v.Len()-1 {
		return nil, errors.New("index is invalid")
	}
	// delete
	t := reflect.MakeSlice(reflect.TypeOf(slice), 0, 0)
	t = reflect.AppendSlice(t, v.Slice(0, index))
	t = reflect.AppendSlice(t, v.Slice(index+1, v.Len()))
	return t.Interface(), nil
}

// UpdateSliceE modifies the specified index element of slice.
// Note that the original slice will not be modified.
func UpdateSliceE(slice interface{}, index int, value interface{}) (interface{}, error) {
	// check params
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, errors.New("target isn't a slice")
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
		return nil, errors.New("target isn't a slice")
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