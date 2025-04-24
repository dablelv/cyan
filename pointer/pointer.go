package pointer

// Pointer returns a pointer to the provided value of type T.
// It's useful in situations where you need a pointer to a value,
// particularly when dealing with literals or temporary values.
//
// Example usage:
// intPtr := Pointer(42)
// strPtr := Pointer("hello")
func Pointer[T any](t T) *T {
	return &t
}

// Value returns the value pointed to by the given pointer of type T.
// If the pointer is nil, it returns the zero value of type T.
// It's useful for safely dereferencing pointers, especially when
// dealing with potentially nil pointers.
//
// Example usage:
// intPtr := Pointer(42)
// intValue := Value(intPtr)  // intValue == 42
//
// var nilPtr *string
// strValue := Value(nilPtr)  // strValue == "" (zero value for string)
func Value[T any](ptr *T) T {
	if ptr == nil {
		var zero T
		return zero
	}

	return *ptr
}
