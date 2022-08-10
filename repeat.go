package linq

import "golang.org/x/exp/constraints"

// Repeat returns a query with value repeated count times.
func Repeat[T any, I constraints.Integer](value T, count I) Query[T] {
	if count == 0 {
		return None[T]()
	}
	return NewQuery(func() Enumerator[T] {
		return func() (T, bool) {
			if count == 0 {
				return value, false
			}
			count--
			return value, true
		}
	})
}

// RepeatForever returns a query with value repeated forever.
func RepeatForever[T any](value T) Query[T] {
	return NewQuery(func() Enumerator[T] {
		return func() (T, bool) {
			return value, true
		}
	})
}
