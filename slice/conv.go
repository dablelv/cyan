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
func ToMap[T any, K comparable](items []T, getKey func(e T) K) map[K]T {
	m := make(map[K]T, len(items))

	for _, item := range items {
		m[getKey(item)] = item
	}
	return m
}

// GroupFunc groups elements of the input slice based on a key extraction function.
// T is the type of the input slice elements, and U is the type of the keys.
// The keys must be comparable, meaning they can be used as map keys.
func GroupFunc[T any, K comparable](items []T, getKey func(T) K) map[K][]T {
	m := make(map[K][]T, len(items))

	for _, item := range items {
		k := getKey(item)
		m[k] = append(m[k], item)
	}
	return m
}
