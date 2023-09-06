package slice

import (
	"reflect"

	"github.com/dablelv/cyan/cond"
)

//
// All of the functions in this file are the supplement to the standard library slices package.
// The standard library functions implemented by generics in the
// https://pkg.go.dev/golang.org/x/exp/slices package should be used first.
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

// Unique replaces repeated elements with a single copy and returns a new slice.
// Unique is like the experimental lib function https://pkg.go.dev/golang.org/x/exp/slices#Compact,
// but Unique do not modify the original slice and the original slice also doesn't need to be sorted.
// Unique implemented by generics is recommended to be used.
// E.g. input []int{1, 2, 3, 2} and output is []int{1, 2, 3}.
func Unique[E comparable, S ~[]E](s S) S {
	if s == nil {
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

// UniqueRef replaces repeated elements with a single copy and returns a new slice.
// The original slice will not be modified.
func UniqueRef(a any) any {
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
