// Package cond contains some functions for conditional judgment. E.g. If, Bool, And, Or, TernaryOperator ...
// The implementation of this package refers to the implementation of carlmjohnson's truthy package, you may find more
// useful information in truthy(https://github.com/carlmjohnson/truthy), thanks carlmjohnson.
package cond

import "reflect"

// Bool returns the truthy value of anything.
// If the value's type has a Bool() bool method, the method is called and returned.
// If the type has an IsZero() bool method, the opposite value is returned.
// Slices and maps are truthy if they have a length greater than zero.
// All other types are truthy if they are not their zero value.
func Bool[T any](value T) bool {
	switch m := any(value).(type) {
	case interface{ Bool() bool }:
		return m.Bool()
	case interface{ IsZero() bool }:
		return !m.IsZero()
	}
	return reflectValue(&value)
}

func reflectValue(v any) bool {
	switch rv := reflect.ValueOf(v).Elem(); rv.Kind() {
	case reflect.Map, reflect.Slice, reflect.Array:
		return rv.Len() != 0
	default:
		return !rv.IsZero()
	}
}

// And returns true if both a and b are truthy.
func And[T, U any](a T, b U) bool {
	return Bool(a) && Bool(b)
}

// Or returns false if neither a nor b is truthy.
func Or[T, U any](a T, b U) bool {
	return Bool(a) || Bool(b)
}

// Xor returns true if a or b but not both is truthy.
func Xor[T, U any](a T, b U) bool {
	valA := Bool(a)
	valB := Bool(b)
	return (valA || valB) && valA != valB
}

// Nor returns true if neither a nor b is truthy.
func Nor[T, U any](a T, b U) bool {
	return !(Bool(a) || Bool(b))
}

// Xnor returns true if both a and b or neither a nor b are truthy.
func Xnor[T, U any](a T, b U) bool {
	valA := Bool(a)
	valB := Bool(b)
	return (valA && valB) || (!valA && !valB)
}

// Nand returns false if both a and b are truthy.
func Nand[T, U any](a T, b U) bool {
	return !Bool(a) || !Bool(b)
}

// If simulates the conditional operator used in C, C++ and Java.
func If[T any](b bool, trueVal, falseVal T) T {
	if b {
		return trueVal
	}
	return falseVal
}
