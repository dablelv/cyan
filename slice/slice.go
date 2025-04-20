package slice

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/dablelv/cyan/cond"
)

//
// All of the functions in this file are the supplement to the standard library slices package.
// The standard library functions implemented by generics in the
// https://pkg.go.dev/slices package should be used first.
//

// Contains reports whether slice contains any one of target elements.
func Contains[E comparable](s []E, targets ...E) bool {
	// Convert the targets to map set.
	m := make(map[E]struct{})
	for i := range targets {
		m[targets[i]] = struct{}{}
	}

	for i := range s {
		if _, ok := m[s[i]]; ok {
			return true
		}
	}
	return false
}

// ContainsRef reports whether slice contains any one of target elements.
// Note that if the target elements are a numeric literal, please specify their type explicitly,
// otherwise type defaults to int.
// E.g. you might call like ContainsRef([]int32{1,2,3}, int32(1)).
// ContainsRef will panic if argument a isn't slice.
func ContainsRef(a any, targets ...any) bool {
	v := reflect.ValueOf(a)
	if v.IsNil() {
		return false
	}

	// Convert the targets to map set.
	m := make(map[any]struct{})
	for i := range targets {
		m[targets[i]] = struct{}{}
	}

	for i := 0; i < v.Len(); i++ {
		if _, ok := m[v.Index(i).Interface()]; ok {
			return true
		}
	}
	return false
}

// ClearZero creates a slice with all zero values removed.
func ClearZero[S ~[]E, E comparable](s S) S {
	r := make([]E, 0, len(s))
	for i := range s {
		if !cond.IsZero(s[i]) {
			r = append(r, s[i])
		}
	}
	return r
}

// ClearZeroRef creates a slice with all zero values removed.
// ClearZeroRef will panic is argument a isn't a slice.
func ClearZeroRef(a any) any {
	v := reflect.ValueOf(a)
	r := reflect.MakeSlice(reflect.TypeOf(a), 0, v.Len())
	for i := 0; i < v.Len(); i++ {
		if !v.Index(i).IsZero() {
			r = reflect.Append(r, v.Index(i))
		}
	}
	return r.Interface()
}

// Reverse reverses the specified slice without modifying the original slice.
// Reverse implemented by generics is recommended to be used.
func Reverse[E any, S ~[]E](s S) S {
	if s == nil {
		return s
	}
	r := make([]E, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return r
}

// ReverseRef implemented by reflect reverses the specified slice without modifying the original slice.
// ReverseRef will panic if argument a isn't Slice.
func ReverseRef(a any) any {
	v := reflect.ValueOf(a)
	if v.IsNil() {
		return a
	}
	r := reflect.MakeSlice(reflect.TypeOf(a), 0, v.Len())
	for i := v.Len() - 1; i >= 0; i-- {
		r = reflect.Append(r, v.Index(i))
	}
	return r.Interface()
}

// Chunk divides the slice into chunks.
func Chunk[T any](slice []T, chunkSize int) [][]T {
	var chunks [][]T
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// FilterFunc filters out elements that do not meet the conditions.
func FilterFunc[T any](data []T, retain func(T) bool) []T {
	r := make([]T, 0, len(data))
	for _, e := range data {
		if retain(e) {
			r = append(r, e)
		}
	}
	return r
}

// Distinct returns a new slice containing only the unique elements from the input slice.
// T is the type of the input slice elements, and it must be comparable.
func Distinct[E comparable, S ~[]E](s S) S {
	if s == nil {
		return s
	}

	r := make(S, 0, len(s))
	m := make(map[E]struct{}, len(s))

	for _, v := range s {
		if _, ok := m[v]; !ok {
			r = append(r, v)
			m[v] = struct{}{}
		}
	}
	return r
}

// DistinctRef returns a new slice containing only the unique elements from the input slice.
// T is the type of the input slice elements, and it must be comparable.
// Deprecated: plz use generic func Distinct first.
func DistinctRef(a any) any {
	v := reflect.ValueOf(a)
	r := reflect.MakeSlice(reflect.TypeOf(a), 0, v.Len())
	m := make(map[any]struct{})
	for i := 0; i < v.Len(); i++ {
		if _, ok := m[v.Index(i).Interface()]; !ok {
			r = reflect.Append(r, v.Index(i))
			m[v.Index(i).Interface()] = struct{}{}
		}
	}
	return r.Interface()
}

// DistinctFunc returns a new slice containing only the unique elements from the input slice,
// based on a key extraction function that determines the uniqueness of each element.
func DistinctFunc[E any, K comparable, S ~[]E](s S, getKey func(item E) K) S {
	if s == nil {
		return s
	}

	r := make(S, 0, len(s))
	m := make(map[K]struct{}, len(s))

	for _, v := range s {
		key := getKey(v)
		if _, ok := m[key]; !ok {
			r = append(r, v)
			m[key] = struct{}{}
		}
	}
	return r
}

// Merge takes multiple slices of type T and combines them into a single slice.
// T can be any type, allowing this function to work with slices of different data types.
func Merge[T any](ss ...[]T) []T {
	var r []T
	for _, s := range ss {
		r = append(r, s...)
	}
	return r
}

// GroupFunc groups elements of the slice based on a key extraction function.
// T is the type of the input slice elements, and U is the type of the keys.
// The keys must be comparable, meaning they can be used as map keys.
func GroupFunc[E any, K comparable, S ~[]E](s S, getKey func(E) K) map[K]S {
	m := make(map[K]S, len(s))

	for _, item := range s {
		k := getKey(item)
		m[k] = append(m[k], item)
	}
	return m
}

// ExtractField takes a slice of structs and a field name, and returns a slice of values for that field.
func ExtractField[S, E any](s []S, fieldName string) ([]E, error) {
	if len(s) == 0 {
		return nil, nil
	}

	rv := reflect.Indirect(reflect.ValueOf(s))
	if reflect.Indirect(rv.Index(0)).Kind() != reflect.Struct {
		return nil, errors.New("the type of slices element is not struct")
	}

	vals := make([]E, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		field := reflect.Indirect(rv.Index(i)).FieldByName(fieldName)
		if !field.IsValid() {
			return nil, fmt.Errorf("no such field: %s in item %#v", fieldName, rv.Index(i))
		}
		vals[i] = field.Interface().(E)
	}
	return vals, nil
}
