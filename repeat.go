package linq

import "golang.org/x/exp/constraints"

// Repeat returns a query with value repeated count times.
func Repeat[T any, I constraints.Integer](value T, count I) Query[T] {
	if count == 0 {
		return None[T]()
	}
	return NewQuery(func() Enumerator[T] {
		var i I = 0
		return func() (T, bool) {
			if i < count {
				i++
				return value, true
			}
			var t T
			return t, false
		}
	}, FastCountOption[T](int(count)))
}

// RepeatForever returns a query with value repeated forever.
func RepeatForever[T any](value T) Query[T] {
	return NewQuery(func() Enumerator[T] {
		return func() (T, bool) {
			return value, true
		}
	})
}
