// Package util contains utility functions.
package util

// Ptr returns a pointer to any arbitrary variable.
func Ptr[T any](x T) *T {
	return &x
}
