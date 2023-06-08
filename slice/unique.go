package slice

import (
	"fmt"
	"reflect"
)

//
// Unique a slice, e.g. input []int{1, 2, 3, 2} and output is []int{1, 2, 3}.
//

// Unique replaces repeated elements with a single copy and returns a new slice.
// Unique is like the experimental lib function https://pkg.go.dev/golang.org/x/exp/slices#Compact,
// but Unique do not modify the original slice and the original slice also doesn't need to be sorted.
// Unique implemented by generics is recommended to be used.
func Unique[E comparable, S ~[]E](s S) S {
	if len(s) == 0 {
		return s
	}
	r := make(S, 0, len(s))
	m := make(map[E]struct{})
	for i := range s {
		if _, ok := m[s[i]]; !ok {
			r = append(r, s[i])
			m[s[i]] = struct{}{}
		}
	}
	return r
}

func UniqueInt(src []int) []int {
	r, _ := UniqueE(src)
	v, _ := r.([]int)
	return v
}

func UniqueInt8(src []int8) []int8 {
	r, _ := UniqueE(src)
	v, _ := r.([]int8)
	return v
}

func UniqueInt16(src []int16) []int16 {
	r, _ := UniqueE(src)
	v, _ := r.([]int16)
	return v
}

func UniqueInt32(src []int32) []int32 {
	r, _ := UniqueE(src)
	v, _ := r.([]int32)
	return v
}

func UniqueInt64(src []int64) []int64 {
	r, _ := UniqueE(src)
	v, _ := r.([]int64)
	return v
}

func UniqueUint(src []uint) []uint {
	r, _ := UniqueE(src)
	v, _ := r.([]uint)
	return v
}

func UniqueUint8(src []uint8) []uint8 {
	r, _ := UniqueE(src)
	v, _ := r.([]uint8)
	return v
}

func UniqueUint16(src []uint16) []uint16 {
	r, _ := UniqueE(src)
	v, _ := r.([]uint16)
	return v
}

func UniqueUint32(src []uint32) []uint32 {
	r, _ := UniqueE(src)
	v, _ := r.([]uint32)
	return v
}

func UniqueUint64(src []uint64) []uint64 {
	r, _ := UniqueE(src)
	v, _ := r.([]uint64)
	return v
}

func UniqueFloat32(src []float32) []float32 {
	r, _ := UniqueE(src)
	v, _ := r.([]float32)
	return v
}

func UniqueFloat64(src []float64) []float64 {
	r, _ := UniqueE(src)
	v, _ := r.([]float64)
	return v
}

func UniqueStr(src []string) []string {
	r, _ := UniqueE(src)
	v, _ := r.([]string)
	return v
}

// UniqueE replaces repeated elements with a single copy and returns a new slice.
// Note that the original slice will not be modified.
func UniqueE(slice any) (any, error) {
	// Check params.
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("the input %#v of type %T isn't a slice", slice, slice)
	}
	// Unique slice.
	r := reflect.MakeSlice(reflect.TypeOf(slice), 0, v.Len())
	m := make(map[any]struct{})
	for i := 0; i < v.Len(); i++ {
		if _, ok := m[v.Index(i).Interface()]; !ok {
			r = reflect.Append(r, v.Index(i))
			m[v.Index(i).Interface()] = struct{}{}
		}
	}
	return r.Interface(), nil
}
