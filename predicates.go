package linq

// All returns true if pred returns true for all elements in q, including if q
// is empty.
func (q Query[T]) All(pred func(t T) bool) bool {
	return All(q, pred)
}

// AllI returns true if pred returns true for all elements in q, including if q
// is empty. The pred function takes the index and value of each element.
func (q Query[T]) AllI(pred func(i int, t T) bool) bool {
	return AllI(q, pred)
}

// Any returns true if pred returns true for at least one element in q.
func (q Query[T]) Any(pred func(t T) bool) bool {
	return Any(q, pred)
}

// AnyI returns true if pred returns true for at least one element in q. The
// pred function takes the index and value of each element.
func (q Query[T]) AnyI(pred func(i int, t T) bool) bool {
	return AnyI(q, pred)
}

// Empty returns true if q has no elements.
func (q Query[T]) Empty() bool {
	return Empty(q)
}

// All returns true if pred returns true for all elements in q, including if q
// is empty.
func All[T any](q Query[T], pred func(t T) bool) bool {
	return AllI(q, indexify(pred))
}

// AllI returns true if pred returns true for all elements in q, including if q
// is empty. The pred function takes the index and value of each element.
func AllI[T any](q Query[T], pred func(i int, t T) bool) bool {
	next := q.Enumerator()
	i := counter(0)
	for t, ok := next(); ok; t, ok = next() {
		if !pred(i(), t) {
			return false
		}
	}
	return true
}

// Any returns true if pred returns true for at least one element in q.
func Any[T any](q Query[T], pred func(t T) bool) bool {
	return AnyI(q, indexify(pred))
}

// AnyI returns true if pred returns true for at least one element in q. The
// pred function takes the index and value of each element.
func AnyI[T any](q Query[T], pred func(i int, t T) bool) bool {
	next := q.Enumerator()
	i := counter(0)
	for t, ok := next(); ok; t, ok = next() {
		if pred(i(), t) {
			return true
		}
	}
	return false
}

// Empty returns true if q has no elements.
func Empty[T any](q Query[T]) bool {
	_, ok := q.Enumerator()()
	return !ok
}
