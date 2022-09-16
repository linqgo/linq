package linq

// DefaultIfEmpty returns q if not empty, otherwise it returns a query
// containing alt.
func (q Query[T]) DefaultIfEmpty(alt T) Query[T] {
	return DefaultIfEmpty(q, alt)
}

// DefaultIfEmpty returns q if not empty, otherwise it returns a query
// containing alt.
func DefaultIfEmpty[T any](q Query[T], alt T) Query[T] {
	count := q.fastCount()
	switch count {
	case -1:
	case 0:
		return From(alt)
	default:
		return q
	}
	return Pipe(q, func(next Enumerator[T]) Enumerator[T] {
		delivered := false
		return func() (T, bool) {
			if next != nil {
				if t, ok := next(); ok {
					delivered = true
					return t, ok
				}
				next = nil
			}
			if !delivered {
				delivered = true
				return alt, true
			}
			var t T
			return t, false
		}
	}, FastCountOption[T](count))
}
