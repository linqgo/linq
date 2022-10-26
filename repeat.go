package linq

import "golang.org/x/exp/constraints"

// Repeat returns a query with value repeated count times.
func Repeat[T any, I constraints.Integer](value T, count I) Query[T] {
	if count == 0 {
		return None[T]()
	}
	n := int(count)
	return NewQuery(
		func() Enumerator[T] {
			var i I = 0
			return func() Maybe[T] {
				if i < count {
					i++
					return Some(value)
				}
				return No[T]()
			}
		},
		FastCountOption[T](int(count)),
		FastGetOption(func(i int) Maybe[T] {
			return NewMaybe(value, 0 <= i && i < n)
		}),
	)
}

// RepeatForever returns a query with value repeated forever.
func RepeatForever[T any](value T) Query[T] {
	return NewQuery(
		func() Enumerator[T] {
			return func() Maybe[T] {
				return Some(value)
			}
		},
		FastGetOption(func(i int) Maybe[T] {
			return NewMaybe(value, 0 <= i)
		}),
	)
}
