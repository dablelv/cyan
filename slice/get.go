package slice

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/dablelv/cyan/math"
	"golang.org/x/exp/constraints"
)

type numerical interface {
	constraints.Integer | constraints.Float
}

// RandomElem returns a random element from a Slice or Array.
// If the length of Slice or Array is zero it will panic.
func RandomElem(a any) any {
	v := reflect.ValueOf(a)
	i := math.RandIntn(v.Len())
	return v.Index(i).Interface()
}

// Min returns the smallest element of the slice and an error if occurred.
// Min implemented by generics is recomended to be used.
func Min[T constraints.Ordered](s []T) T {
	var t T
	if len(s) == 0 {
		return t
	}
	t = s[0]
	for i := 1; i < len(s); i++ {
		if s[i] < t {
			t = s[i]
		}
	}
	return t
}

// Max returns the largest element of the slice.
// Max implemented by generics is recomended to be used.
func Max[T constraints.Ordered](s []T) T {
	var t T
	if len(s) == 0 {
		return t
	}
	t = s[0]
	for i := 1; i < len(s); i++ {
		if s[i] > t {
			t = s[i]
		}
	}
	return t
}

// MinRef returns the smallest element of the Slice or Array.
func MinRef(a any) any {
	r, _ := MinRefE(a)
	return r
}

// MinRefE returns the smallest element of the Slice or Array.
// MinRefE will panic if argument a isn't Slice or Array.
func MinRefE(a any) (any, error) {
	val := reflect.ValueOf(a)
	if val.Len() == 0 {
		return nil, errors.New("No elements")
	}
	min := val.Index(0).Interface()
	for i := 1; i < val.Len(); i++ {
		switch v := val.Index(i).Interface().(type) {
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
			return nil, fmt.Errorf("The element %#v of type %T isn't numerical type", v, v)
		}
	}
	return min, nil
}

// MaxRef returns the largest element of the Slice or Array.
func MaxRef(a any) any {
	r, _ := MaxRefE(a)
	return r
}

// MaxRefE returns the largest element of the Slice or Array.
// MaxRefE will panic if argument a isn't Slice or Array.
func MaxRefE(a any) (any, error) {
	v := reflect.ValueOf(a)
	if v.Len() == 0 {
		return nil, errors.New("No elements")
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
			return nil, fmt.Errorf("The element %#v of type %T isn't numerical type", v, v)
		}
	}
	return max, nil
}

// Sum calculates the sum of slice elements.
// Sum implemented by generics is recommended to be used.
func Sum[S ~[]E, E numerical](s S) float64 {
	var sum float64
	for _, v := range s {
		sum += float64(v)
	}
	return sum
}

// SumRef sum slice or array elements.
func SumRef(a any) float64 {
	v, _ := SumRefE(a)
	return v
}

// SumRefE returns the sum of slice or array elements implemented by reflect.
// E.g. input []int32{1, 2, 3} and output is 6.
// SumRefE will panic if argument a's Kind isn't Slice, Array or String.
// If a's Kind is String, SumRefE will sum characters encoding in byte.
func SumRefE(a any) (float64, error) {
	var sum float64
	v := reflect.ValueOf(a)
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
			return 0.0, fmt.Errorf("The element %#v of slice type %T isn't numerical type", v, v)
		}
	}
	return sum, nil
}
