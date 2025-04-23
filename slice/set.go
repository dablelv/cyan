package slice

import "github.com/dablelv/cyan/conv"

// Intersection get the intersection of two slice elements.
func Intersection[T comparable](ls, rs []T) []T {
	if ls == nil || rs == nil {
		return nil
	}

	m := conv.SliceToSet(ls)
	r := make([]T, 0, len(ls))
	for i := range rs {
		if _, ok := m[rs[i]]; ok {
			r = append(r, rs[i])
		}
	}
	return r
}

// IntersectionFunc get the intersection of two slice elements,
// based on a key extraction function that determines the equality of each element.
func IntersectionFunc[T any, K comparable](s1, s2 []T, getKey func(e T) K) []T {
	if s1 == nil || s2 == nil {
		return nil
	}
	m := conv.SliceToSetFunc(s1, getKey)
	r := make([]T, 0, len(s1))
	for _, v := range s2 {
		if _, ok := m[getKey(v)]; ok {
			r = append(r, v)
		}
	}
	return r
}

// Diff computes the difference of left slice relative to right slice.
func Diff[T comparable](ls, rs []T) []T {
	if len(ls) == 0 {
		return ls
	}
	m := conv.SliceToSet(rs)
	r := make([]T, 0, len(ls))
	for i := range ls {
		if _, ok := m[ls[i]]; !ok {
			r = append(r, ls[i])
		}
	}
	return r
}

// DiffFunc computes the difference of set1 relative to set2 based on a specified key extraction function.
func DiffFunc[T any, K comparable](ls, rs []T, getKey func(e T) K) []T {
	if len(ls) == 0 {
		return ls
	}
	m := conv.SliceToSetFunc(rs, getKey)
	r := make([]T, 0, len(ls))
	for i := range ls {
		if _, ok := m[getKey(ls[i])]; !ok {
			r = append(r, ls[i])
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
