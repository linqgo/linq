package linq

// OfType returns a Query that contains all the elements of q that have type U.
func OfType[U, T any](q Query[T]) Query[U] {
	return Pipe(q, func(next Enumerator[T]) Enumerator[U] {
		return func() Maybe[U] {
			for t, ok := next().Get(); ok; t, ok = next().Get() {
				var i any = t
				if u, is := i.(U); is {
					return Some(u)
				}
			}
			return No[U]()
		}
	}, FastCountIfEmptyOption[U](q.fastCount()))
}
