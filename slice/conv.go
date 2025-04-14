package slice

// ToMap converts a slice of type T into a map, using a specified key extraction function.
// T can be any type, and K is the type of the keys used in the resulting map, which must be comparable.
func ToMap[T any, K comparable](items []T, getKey func(e T) K) map[K]T {
	m := make(map[K]T, len(items))

	for _, item := range items {
		m[getKey(item)] = item
	}
	return m
}

func GroupBy[T any, K comparable](items []T, getKey func(T) K) map[K][]T {
	m := make(map[K][]T, len(items))

	for _, item := range items {
		k := getKey(item)
		m[k] = append(m[k], item)
	}
	return m
}
