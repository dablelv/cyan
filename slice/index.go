package slice

import (
	"fmt"
	"reflect"

	"github.com/dablelv/go-huge-util/math"
)

//
// All of the functions in this file are the supplement to the standard library slices package.
// You should also know the Index and IndexFunc function in the https://pkg.go.dev/golang.org/x/exp/slices package.
// The standard library functions implemented by generics should be used first.
//

// Indexes returns the specified element all indexes from a slice or array.
func Indexes(a any, value any) []int {
	indexes, _ := IndexesE(a, value)
	return indexes
}

// IndexesE returns the specified element all indexes from a slice or array with returned error.
func IndexesE(a any, value any) ([]int, error) {
	// Check arguments.
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return nil, fmt.Errorf("The input %#v of type %T isn't a slice and array", a, a)
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

// Indices returns the specified element all indexes.
// Indices implemented by generics has a better performance than Indexes implemented by generics
// and be recommended to use.
func Indices[E comparable](s []E, v E) []int {
	var indices []int
	for i := range s {
		if v == s[i] {
			indices = append(indices, i)
		}
	}
	return indices
}

// IndicesFunc returns the specified element all indexes satisfying f(s[i]).
func IndicesFunc[E any](s []E, f func(E) bool) []int {
	var indices []int
	for i := range s {
		if f(s[i]) {
			indices = append(indices, i)
		}
	}
	return indices
}

// RandomElem returns a random element from a slice or array.
// If the length of slice or array is zero it will panic.
func RandomElem(a any) any {
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return a
	}
	i := math.RandIntn(v.Len())
	return v.Index(i).Interface()
}
