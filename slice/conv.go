package slice

// Map applies a transformation function to each element of the input slice
// and returns a new slice containing the results.
// T is the type of the input slice elements, and U is the type of the output slice elements.
func Map[T, U any](data []T, f func(T) U) []U {
	res := make([]U, 0, len(data))

	for _, e := range data {
		res = append(res, f(e))
	}
	return res
}

// ToMap converts a slice of type T into a map, using a specified key extraction function.
// T can be any type, and K is the type of the keys used in the resulting map, which must be comparable.
func ToMap[E any, K comparable, S ~[]E](s S, getKey func(e E) K) map[K]E {
	m := make(map[K]E, len(s))
	for _, v := range s {
		m[getKey(v)] = v
	}
	return m
}
