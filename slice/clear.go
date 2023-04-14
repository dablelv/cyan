package slice

import (
	"reflect"

	"github.com/dablelv/go-huge-util/cond"
)

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
// ClearZeroRef is implemented base on reflection.
func ClearZeroRef[S ~[]E, E any](s S) S {
	r := make([]E, 0, len(s))
	for i := range s {
		if !reflect.ValueOf(s[i]).IsZero() {
			r = append(r, s[i])
		}
	}
	return r
}
