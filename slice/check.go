package slice

import "reflect"

// Contains reports whether slice contains any one of target elements.
func Contains[E comparable, S ~[]E](s S, targets ...E) bool {
	// Convert the targets to map set.
	m := make(map[E]struct{}, len(s))
	for i := range targets {
		m[targets[i]] = struct{}{}
	}

	for i := range s {
		if _, ok := m[s[i]]; ok {
			return true
		}
	}
	return false
}

// ContainsRef reports whether slice contains any one of target elements.
// Note that if the target elements are a numeric literal, please specify their type explicitly,
// otherwise type defaults to int.
// E.g. you might call like ContainsRef([]int32{1,2,3}, int32(1)).
// ContainsRef will panic if argument a isn't slice.
func ContainsRef(a any, targets ...any) bool {
	v := reflect.ValueOf(a)
	if v.IsNil() {
		return false
	}

	// Convert the target slice to map set.
	m := make(map[any]struct{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		m[v.Index(i).Interface()] = struct{}{}
	}

	for i := range targets {
		if _, ok := m[targets[i]]; ok {
			return true
		}
	}
	return false
}

// ContainsAll checks whether all given elements exist in the target slice.
func ContainsAll[E comparable, S ~[]E](s, elems S) bool {
	m := make(map[E]struct{}, len(s))
	for _, item := range s {
		m[item] = struct{}{}
	}

	// Check if all required elements exist in the map.
	for _, elem := range elems {
		if _, exists := m[elem]; !exists {
			return false
		}
	}
	return true
}
