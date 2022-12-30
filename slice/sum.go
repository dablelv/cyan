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
func Sum[T numerical](s []T) float64 {
	var sum float64
	for _, v := range s {
		sum += float64(v)
	}
	return sum
}

// SumSlice calculates the sum of slice elements.
// E.g. input []int32{1, 2, 3} and output is 6.
func SumSlice(slice any) float64 {
	v, _ := SumSliceE(slice)
	return v
}

// SumSliceE returns the sum of slice elements and an error if occurred.
func SumSliceE(slice any) (float64, error) {
	// Check param.
	if slice == nil {
		return 0.0, nil
	}
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
