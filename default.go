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
		return Pipe(q, func(next Enumerator[T]) Enumerator[T] {
			delivered := false
			return func() Maybe[T] {
				if next != nil {
					if t := next(); t.Valid() {
						delivered = true
						return t
					}
					next = nil
				}
				if !delivered {
					delivered = true
					return Some(alt)
				}
				return No[T]()
			}
		})
	case 0:
		return From(alt)
	default:
		return q
	}
}
