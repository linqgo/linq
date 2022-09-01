package linq

// OfType returns a Query that contains all the elements of q that have type U.
func OfType[U, T any](q Query[T]) Query[U] {
	return Pipe(q, func(next Enumerator[T]) Enumerator[U] {
		return func() (U, bool) {
			for t, ok := next(); ok; t, ok = next() {
				var i any = t
				if u, is := i.(U); is {
					return u, true
				}
			}
			var u U
			return u, false
		}
	})
}
