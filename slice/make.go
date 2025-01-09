package slice

// Make makes a slice with specified element, len and capacity.
// If there is no size argument to call Make, e.g. Make(1) will return a int slice []int{1} only including one specified element,
// The first size argument represents the slice length.
// A second size argument may be provided to specify a different capacity.
// From above description we can know that Make(e) is equal to Make(e, 1) and Make(e, 1, 1).
func Make[E any](e E, size ...int) []E {
	var slice []E
	switch l := len(size); {
	case l >= 2:
		slice = make([]E, size[0], size[1])
	case l >= 1:
		slice = make([]E, size[0])
	default:
		return []E{e}
	}
	for i := 0; i < size[0]; i++ {
		slice[i] = e
	}
	return slice
}

// Chunk divide the slice into chunks.
func Chunk[T any](slice []T, chunkSize int) [][]T {
	var chunks [][]T
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// Filter filter out elements that do not meet the conditions.
func Filter[T any](data []T, retain func(T) bool) []T {
	res := make([]T, 0, len(data))

	for _, e := range data {
		if retain(e) {
			res = append(res, e)
		}
	}
	return res
}

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

// GroupFunc groups elements of the input slice based on a key extraction function.
// T is the type of the input slice elements, and U is the type of the keys.
// The keys must be comparable, meaning they can be used as map keys.
func GroupFunc[T, U comparable](data []T, key func(T) U) map[U][]T {
	res := make(map[U][]T, len(data))

	for _, e := range data {
		res[key(e)] = append(res[key(e)], e)
	}
	return res
}

// Distinct returns a new slice containing only the unique elements from the input slice.
// T is the type of the input slice elements, and it must be comparable.
func Distinct[T comparable](data []T) []T {
	r := make([]T, 0, len(data))
	m := make(map[T]struct{}, len(data))

	for _, n := range data {
		if _, ok := m[n]; !ok {
			r = append(r, n)
		}
		m[n] = struct{}{}
	}
	return r
}

// DistinctFunc returns a new slice containing only the unique elements from the input slice,
// based on a key extraction function that determines the uniqueness of each element.
// T is the type of the input slice elements, and TK is the type of the keys used for comparison.
// TK must be comparable, meaning it can be used as a key in a map.
func DistinctFunc[T any, K comparable](data []T, key func(item T) K) []T {
	r := make([]T, 0, len(data))
	m := make(map[K]struct{}, len(data))
	for _, v := range data {
		k := key(v)
		if _, ok := m[k]; !ok {
			r = append(r, v)
		}
		m[k] = struct{}{}
	}
	return r
}

// Merge takes multiple slices of type T and combines them into a single slice.
// T can be any type, allowing this function to work with slices of different data types.
func Merge[T any](ss ...[]T) []T {
	var r []T
	for _, s := range ss {
		r = append(r, s...)
	}
	return r
}
