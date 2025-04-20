package slice

import "github.com/dablelv/cyan/conv"

// Intersection get the intersection of two slice elements.
func Intersection[T comparable](s1, s2 []T) []T {
	r := make([]T, 0, len(s1))
	m := conv.SliceToSet(s1)

	for _, v := range s2 {
		if _, ok := m[v]; ok {
			r = append(r, v)
		}
	}
	return r
}

// IntersectionFunc get the intersection of two slice elements,
// based on a key extraction function that determines the equality of each element.
func IntersectionFunc[T any, K comparable](s1, s2 []T, getKey func(e T) K) []T {
	m := conv.SliceToSetFunc(s1, getKey)
	r := make([]T, 0, len(s1))
	for _, v := range s2 {
		if _, ok := m[getKey(v)]; ok {
			r = append(r, v)
		}
	}
	return r
}

// Diff computes the difference of set1 relative to set2.
func Diff[T comparable](s1 []T, s2 []T) []T {
	m := conv.SliceToSet(s2)
	r := make([]T, 0, len(s1))
	for _, v := range s1 {
		if _, ok := m[v]; !ok {
			r = append(r, v)
		}
	}
	return r
}

// DiffFunc computes the difference of set1 relative to set2 based on a specified key extraction function.
func DiffFunc[T any, K comparable](set1 []T, set2 []T, getKey func(e T) K) []T {
	m := conv.SliceToSetFunc(set2, getKey)
	r := make([]T, 0, len(set1))
	for _, item := range set1 {
		if _, ok := m[getKey(item)]; !ok {
			r = append(r, item)
		}
	}
	return r
}

// SymDiff computes the symmetric difference between two slices of type T.
func SymDiff[T comparable](s1, s2 []T) []T {
	r := make([]T, 0, len(s1)+len(s2))

	m1 := conv.SliceToSet(s1)
	m2 := conv.SliceToSet(s2)

	for _, v := range s1 {
		if _, ok := m2[v]; !ok {
			r = append(r, v)
		}
	}

	for _, v := range s2 {
		if _, ok := m1[v]; !ok {
			r = append(r, v)
		}
	}
	return r
}

// SymDiffFunc computes the symmetric difference between two slices of type T,
// based on a key extraction function that determines the equality of each element.
func SymDiffFunc[T any, K comparable](s1, s2 []T, getKey func(e T) K) []T {
	r := make([]T, 0, len(s1)+len(s2))

	m1 := conv.SliceToSetFunc(s1, getKey)
	m2 := conv.SliceToSetFunc(s2, getKey)

	for _, v := range s1 {
		if _, ok := m2[getKey(v)]; !ok {
			r = append(r, v)
		}
	}

	for _, v := range s2 {
		if _, ok := m1[getKey(v)]; !ok {
			r = append(r, v)
		}
	}
	return r
}
