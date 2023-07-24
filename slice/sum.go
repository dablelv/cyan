package slice

import (
	"fmt"
	"reflect"

	"golang.org/x/exp/constraints"
)

type numerical interface {
	constraints.Integer | constraints.Float
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
