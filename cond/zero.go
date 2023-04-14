package cond

import "reflect"

// IsZero reports whether v is the zero value for its type.
func IsZero[T comparable](v T) bool {
	var zero T
	return v == zero
}

// IsZeroRef reports whether v is the zero value for its type.
// IsZeroRef is implemented base on reflection.
func IsZeroRef[T any](v T) bool {
	return reflect.ValueOf(&v).Elem().IsZero()
}
