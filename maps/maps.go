package maps

import (
	"cmp"
	"slices"
)

// Keys returns a slice of all the keys in m.
// The keys returned are in indeterminate order.
// As of Go 1.23, you can also use standard library https://pkg.go.dev/maps#Keys.
func Keys[K comparable, V any, M ~map[K]V](m M) []K {
	s := make([]K, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s
}

// SortedKeys returns the keys from a map in ascending order.
func SortedKeys[K cmp.Ordered, V any, M ~map[K]V](m M) []K {
	ks := Keys(m)
	slices.Sort(ks)
	return ks
}

// Vals returns a slice of all the values in m.
// The values returned are in indeterminate order.
// As of Go 1.23, you can also use standard library https://pkg.go.dev/maps#Values.
func Vals[K comparable, V any, M ~map[K]V](m M) []V {
	vs := make([]V, 0, len(m))
	for _, v := range m {
		vs = append(vs, v)
	}
	return vs
}

// SortedVals returns the values from a map in ascending order.
func SortedVals[K comparable, V cmp.Ordered, M ~map[K]V](m M) []V {
	vs := Vals(m)
	slices.Sort(vs)
	return vs
}

// KeyVals returns the keys and values from a map in indeterminate order.
func KeyVals[K comparable, V any, M ~map[K]V](m M) ([]K, []V) {
	ks, vs := make([]K, 0, len(m)), make([]V, 0, len(m))
	for k, v := range m {
		ks = append(ks, k)
		vs = append(vs, v)
	}
	return ks, vs
}
