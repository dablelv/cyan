package slice

import (
	"fmt"
	"reflect"

	"github.com/dablelv/go-huge-util/math"
)

//
// Note that since Go 1.18, this file is deprecated.
// Please use the standard lib function https://pkg.go.dev/golang.org/x/exp/slices#Index and IndexFunc
// implemented by generics.
//

// GetElemIndexes finds all indexes for the specified element in a slice or array.
func GetElemIndexes(a any, value any) []int {
	indexes, _ := GetElemIndexesE(a, value)
	return indexes
}

// GetElemIndexesE finds all indexes for the specified element in a slice with returned error.
func GetElemIndexesE(a any, value any) ([]int, error) {
	// Check params.
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return nil, fmt.Errorf("the input %#v of type %T isn't a slice and array", a, a)
	}
	// Get indexes.
	var indexes []int
	for i := 0; i < v.Len(); i++ {
		if v.Index(i).Interface() == value {
			indexes = append(indexes, i)
		}
	}
	return indexes, nil
}

// GetRandomElem get a random element from a slice or array.
// If the length of slice or array is zero it will panic.
func GetRandomElem(i any) any {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return i
	}
	idx := math.RandIntn(v.Len())
	return v.Index(idx).Interface()
}
