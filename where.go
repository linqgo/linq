package linq

// Where returns a query with elements from q for which pred returns true.
func (q Query[T]) Where(pred func(t T) bool) Query[T] {
	return Where(q, pred)
}

// Where returns a query with elements from q for which pred returns true.
func Where[T any](q Query[T], pred func(t T) bool) Query[T] {
	return Pipe(q, func(next Enumerator[T]) Enumerator[T] {
		return func() (t T, ok bool) {
			for t, ok := next(); ok; t, ok = next() {
				if pred(t) {
					return t, true
				}
			}
			return t, ok
		}
	}, FastCountIfEmptyOption[T](q.fastCount()))
}
