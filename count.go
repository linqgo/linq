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
func (q Query[T]) FastCount() (int, bool) {
	return FastCount(q)
}

func (q Query[T]) MustFastCount() int {
	return MustFastCount(q)
}

// Count returns the number of elements in q.
func Count[T any](q Query[T]) int {
	if c, ok := q.FastCount(); ok {
		return c
	}
	next := q.Enumerator()
	n := 0
	for _, ok := next(); ok; _, ok = next() {
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
	for _, ok := next(); ok && n < limit; _, ok = next() {
		n++
	}
	return n
}

func FastCount[T any](q Query[T]) (int, bool) {
	count := q.fastCount()
	return count, count >= 0
}

func MustFastCount[T any](q Query[T]) int {
	if c, ok := FastCount(q); ok {
		return c
	}
	panic(noFastCountError)
}
