package slice

import (
	"errors"
	"fmt"
	"reflect"

	"golang.org/x/exp/constraints"
)

//
// Desc: get the min or max element of a slice.
// Other useful functions for any type slices can refer to the standard lib
// https://pkg.go.dev/golang.org/x/exp/slices.
//

// Min returns the smallest element of the slice.
// Min implemented by generics is recomended to be used.
func Min[T constraints.Ordered](s []T) T {
	v, _ := MinE(s)
	return v
}

// MinSliceE returns the smallest element of the slice and an error if occurred.
// MinSliceE implemented by generics is recomended to be used.
func MinE[T constraints.Ordered](s []T) (T, error) {
	var t T
	if len(s) == 0 {
		return t, errors.New("no element")
	}
	t = s[0]
	for i := 1; i < len(s); i++ {
		if s[i] < t {
			t = s[i]
		}
	}
	return t, nil
}

// Max returns the largest element of the slice.
// Max implemented by generics is recomended to be used.
func Max[T constraints.Ordered](s []T) T {
	v, _ := MaxE(s)
	return v
}

// MaxE returns the largest element of the slice and an error if error occurred.
// MaxE implemented by generics is recomended to be used.
func MaxE[T constraints.Ordered](s []T) (T, error) {
	var t T
	if len(s) == 0 {
		return t, errors.New("no element")
	}
	t = s[0]
	for i := 1; i < len(s); i++ {
		if s[i] > t {
			t = s[i]
		}
	}
	return t, nil
}

func MinIntSlice(s []int) int {
	min, _ := MinSliceE(s)
	v, _ := min.(int)
	return v
}

func MinInt8Slice(s []int8) int8 {
	min, _ := MinSliceE(s)
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

func MaxUintSlice(sl []uint) uint {
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

// MinSliceE returns the smallest element of the slice and an error if error occurred.
// If slice length is zero return the zero value of the element type.
func MinSliceE(slice any) (any, error) {
	// Check param.
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("the input %#v of type %T isn't a slice", slice, slice)
	}
	if v.Len() == 0 {
		return nil, errors.New("no element")
	}
	// Get the min element.
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

// MaxSliceE returns the largest element of the slice and an error if error occurred.
// If slice length is zero return the zero value of the element type.
func MaxSliceE(slice any) (any, error) {
	// Check param.
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("the input %#v of type %T isn't a slice", slice, slice)
	}
	if v.Len() == 0 {
		return nil, errors.New("no element")
	}
	// Get the max element.
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
