package linq

import "golang.org/x/exp/constraints"

// Identity returns t unmodified.
func Identity[T any](t T) T {
	return t
}

func Deref[T any](t *T) T {
	return *t
}

func False[T any](T) bool {
	return false
}

func Greater[T constraints.Ordered](a, b T) bool {
	return a > b
}

func Less[T constraints.Ordered](a, b T) bool {
	return a < b
}

func LongerSlice[T any](a, b []T) bool {
	return len(a) > len(b)
}

func LongerMap[K comparable, V any](a, b map[K]V) bool {
	return len(a) > len(b)
}

func Pointer[T any](t T) *T {
	return &t
}

func ShorterSlice[T any](a, b []T) bool {
	return len(a) < len(b)
}

func ShorterMap[K comparable, V any](a, b map[K]V) bool {
	return len(a) < len(b)
}

func True[T any](T) bool {
	return true
}

func Zero[T, U any](T) U {
	var u U
	return u
}

func Drain[T any](next Enumerator[T]) {
	for _, ok := next().Get(); ok; _, ok = next().Get() {
	}
}
