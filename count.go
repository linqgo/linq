package linq

// Count returns the number of elements in q.
func (q Query[T]) Count() int {
	return Count(q)
}

// CountLimit returns the number of elements in q.
func (q Query[T]) CountLimit(limit int) int {
	return CountLimit(q, limit)
}

// Count returns the number of elements in q.
func (q Query[T]) FastCount() Maybe[int] {
	return FastCount(q)
}

// Count returns the number of elements in q.
func Count[T any](q Query[T]) int {
	if c, ok := q.FastCount().Get(); ok {
		return c
	}
	next := q.Enumerator()
	n := 0
	for t := next(); t.Valid(); t = next() {
		n++
	}
	return n
}

// CountLimit returns the lower of limit and q.Count(). This is useful when you
// need a sense of how big the input is, but only need to know up to a point.
// For example, you may need to specify pagination controls for a collection
// that has at least 10 elements.
func CountLimit[T any](q Query[T], limit int) int {
	next := q.Enumerator()
	n := 0
	for t := next(); t.Valid() && n < limit; t = next() {
		n++
	}
	return n
}

func FastCount[T any](q Query[T]) Maybe[int] {
	count := q.fastCount()
	return NewMaybe(count, count >= 0)
}
