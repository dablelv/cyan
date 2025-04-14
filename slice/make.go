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

// Chunk divides the slice into chunks.
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

// Distinct returns a new slice containing only the unique elements from the input slice.
// T is the type of the input slice elements, and it must be comparable.
func Distinct[E comparable, S ~[]E](s S) S {
	if s == nil {
		return s
	}

	r := make(S, 0, len(s))
	m := make(map[E]struct{}, len(s))

	for _, v := range s {
		if _, ok := m[v]; !ok {
			r = append(r, v)
			m[v] = struct{}{}
		}
	}
	return r
}

// DistinctFunc returns a new slice containing only the unique elements from the input slice,
// based on a key extraction function that determines the uniqueness of each element.
// T is the type of the input slice elements, and TK is the type of the keys used for comparison.
// TK must be comparable, meaning it can be used as a key in a map.
func DistinctFunc[E any, K comparable, S ~[]E](s S, getKey func(item E) K) S {
	if s == nil {
		return s
	}

	r := make(S, 0, len(s))
	m := make(map[K]struct{}, len(s))

	for _, v := range s {
		key := getKey(v)
		if _, ok := m[key]; !ok {
			r = append(r, v)
			m[key] = struct{}{}
		}
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
