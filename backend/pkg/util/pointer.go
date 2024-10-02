package util

// Ptr returns a pointer to t
func Ptr[T any](t T) *T {
	return &t
}
