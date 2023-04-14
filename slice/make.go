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
