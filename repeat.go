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
			count -= 1
			return value, true
		}
	})
}
