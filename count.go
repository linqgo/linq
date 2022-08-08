package linq

// Count returns the number of elements in q.
func (q Query[T]) Count() int {
	return Count(q)
}

// Count returns the number of elements in q.
func Count[T any](q Query[T]) int {
	next := q.Enumerator()
	n := 0
	for _, ok := next(); ok; _, ok = next() {
		n++
	}
	return n
}
