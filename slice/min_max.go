package slice

import (
	"errors"

	"golang.org/x/exp/constraints"
)

//
// Desc: get the min or max element of a slice.
// Other useful functions with slices of any type can refer standard lib
// https://pkg.go.dev/golang.org/x/exp/slices.
//

// MinSlice returns the smallest element of the slice.
func MinSlice[T constraints.Ordered](s []T) T {
	v, _ := MinSliceE(s)
	return v
}

// MinSliceE returns the smallest element of the slice and an error if occurred.
func MinSliceE[T constraints.Ordered](s []T) (T, error) {
	var t T
	if len(s) == 0 {
		return t, errors.New("no element")
	}
	t = s[0]
	for i := 1; i < len(s); i++ {
		if s[i] < t {
			t = s[i]
		}
	}
	return t, nil
}

// MaxSlice returns the largest element of the slice.
func MaxSlice[T constraints.Ordered](s []T) T {
	v, _ := MaxSliceE(s)
	return v
}

// MaxSliceE returns the largest element of the slice and an error if occurred.
func MaxSliceE[T constraints.Ordered](s []T) (T, error) {
	var t T
	if len(s) == 0 {
		return t, errors.New("no element")
	}
	t = s[0]
	for i := 1; i < len(s); i++ {
		if s[i] > t {
			t = s[i]
		}
	}
	return t, nil
}
