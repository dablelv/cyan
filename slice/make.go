package slice

// MakeSlice make a slice wiht specified element.
// The first integer parameter represents the slice length.
// A second integer argument may be provided to specify a different capacity.
// If there is no integer parameter, MakeSlice will return a slice only include one specified element.
// Which means that MakeSlice(val) is equal to MakeSlice(val, 1).
func MakeSlice[T any](val T, size ...int) []T {
	var slice []T
	switch l := len(size); {
	case l >= 2:
		slice = make([]T, size[0], size[1])
	case l >= 1:
		slice = make([]T, size[0])
	default:
		return []T{val}
	}
	for i := 0; i < size[0]; i++ {
		slice[i] = val
	}
	return slice
}
