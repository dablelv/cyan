package slice

import (
	"reflect"
)

//
// All of the functions in this file are the supplement to the standard library slices package.
// The standard library functions Insert, Delete, DeleteFunc, Replace, Index, IndexFunc
// implemented by generics in the https://pkg.go.dev/slices package should be used first.
//

// Insert inserts the values v... into s at index i and returns the result slice.
// Unlike the standard library, Insert won't modify the original slice.
func Insert[S ~[]E, E any](s S, i int, v ...E) S {
	r := make(S, len(s)+len(v))
	copy(r, s[:i])
	copy(r[i:], v)
	copy(r[i+len(v):], s[i:])
	return r
}

// InsertRef inserts elements to slice in the specified index implemented by reflect.
// Note that the original slice will not be modified.
func InsertRef(a any, i int, v ...any) any {
	val := reflect.ValueOf(a)
	t := reflect.TypeOf(a)

	r := reflect.MakeSlice(t, 0, val.Len()+len(v))

	r = reflect.AppendSlice(r, val.Slice(0, i))
	for i := range v {
		r = reflect.Append(r, reflect.ValueOf(v[i]))
	}
	r = reflect.AppendSlice(r, val.Slice(i, val.Len()))
	return r.Interface()
}

// Delete removes the specified indexes elements from the slice.
// Unlike the standard library function Delete, Delete won't modify the original slice.
func Delete[S ~[]E, E any](s S, indexes ...int) S {
	// Convert the indexes to map set.
	m := make(map[int]struct{})
	for _, i := range indexes {
		m[i] = struct{}{}
	}

	// Filter out specified index elements.
	r := make(S, 0, len(s))
	for i := range s {
		if _, ok := m[i]; !ok {
			r = append(r, s[i])
		}
	}
	return r
}

// DeleteRef removes the elements specified by indexes from the slice.
// Note that the original slice will not be modified.
func DeleteRef(a any, indexes ...int) any {
	// Convert the indexes to map set.
	m := make(map[int]struct{})
	for _, i := range indexes {
		m[i] = struct{}{}
	}

	// Filter out specified index elements.
	v := reflect.ValueOf(a)
	t := reflect.MakeSlice(reflect.TypeOf(a), 0, v.Len())
	for i := 0; i < v.Len(); i++ {
		if _, ok := m[i]; !ok {
			t = reflect.Append(t, v.Index(i))
		}
	}
	return t.Interface()
}

// DeleteElems removes the specified elements from slice.
// Note that the original slice will not be modified.
func DeleteElems[S ~[]E, E comparable](s S, elms ...E) S {
	// Convert the elements to map set.
	m := make(map[E]struct{})
	for i := range elms {
		m[elms[i]] = struct{}{}
	}

	// Filter out specified elements.
	r := make(S, 0, len(s))
	for i := range s {
		if _, ok := m[s[i]]; !ok {
			r = append(r, s[i])
		}
	}
	return r
}

// DeleteElemsRef removes the specified elements from slice implemented by reflect.
// Note that the original slice will not be modified.
func DeleteElemsRef(a any, elms ...any) any {
	// Convert the elements to map set.
	m := make(map[any]struct{})
	for i := range elms {
		m[elms[i]] = struct{}{}
	}

	// Filter out specified elements.
	v := reflect.ValueOf(a)
	r := reflect.MakeSlice(reflect.TypeOf(a), 0, v.Len())
	for i := 0; i < v.Len(); i++ {
		if _, ok := m[v.Index(i).Interface()]; !ok {
			r = reflect.Append(r, v.Index(i))
		}
	}
	return r.Interface()
}

// Indexes returns the specified element all indexes.
// Indexes implemented by generics has a better performance than IndexesRef implemented by reflect,
// so Indexes is recommended to be used while not IndexesRef.
func Indexes[E comparable](s []E, v E) []int {
	var indexes []int
	for i := range s {
		if v == s[i] {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

// IndexesFunc returns the specified element all indexes satisfying f(s[i]).
func IndexesFunc[E any](s []E, f func(E) bool) []int {
	var indices []int
	for i := range s {
		if f(s[i]) {
			indices = append(indices, i)
		}
	}
	return indices
}

// IndexesRef returns the specified element all indexes from a slice or array with returned error.
func IndexesRef(a any, value any) []int {
	v := reflect.ValueOf(a)
	var indexes []int
	for i := 0; i < v.Len(); i++ {
		if v.Index(i).Interface() == value {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

// Replace replaces the elements s[i:j] by the given v.
// The resulting slice len is equal to the original slice.
// Replace panics if s[i:j] is not a valid slice of s.
// Note that the original slice will not be modified.
func Replace[S ~[]E, E any](s S, i, j int, v ...E) S {
	_ = s[i:j] // verify that i:j is a valid subslice
	r := make(S, len(s))
	copy(r, s[:i])
	for k := i; k < j; k++ {
		r[k] = v[k-i]
	}
	copy(r[j:], s[j:])
	return r
}

// ReplaceRef modifies the specified index elements of slice implemented by reflect.
// The resulting slice len is equal to the original slice.
// ReplaceRef panics if s[i:j] is not a valid slice of s.
// Note that the original slice will not be modified.
func ReplaceRef(a any, i, j int, v ...any) any {
	val := reflect.ValueOf(a)
	r := reflect.MakeSlice(reflect.TypeOf(a), val.Len(), val.Len())
	reflect.Copy(r.Slice(0, i), val.Slice(0, i))
	for k := i; k < j; k++ {
		r.Index(k).Set(reflect.ValueOf(v[k-i]))
	}
	reflect.Copy(r.Slice(j, val.Len()), val.Slice(j, val.Len()))
	return r.Interface()
}
