package linq

// All returns true if pred returns true for all elements in q, including if q
// is empty.
func (q Query[T]) All(pred func(t T) bool) bool {
	return All(q, pred)
}

// Any returns true if pred returns true for at least one element in q.
func (q Query[T]) Any(pred func(t T) bool) bool {
	return Any(q, pred)
}

// Empty returns true if q has no elements.
func (q Query[T]) Empty() bool {
	return Empty(q)
}

// All returns true if pred returns true for all elements in q, including if q
// is empty.
func All[T any](q Query[T], pred func(t T) bool) bool {
	return !Any(q, Not(pred))
}

// Any returns true if pred returns true for at least one element in q.
func Any[T any](q Query[T], pred func(t T) bool) bool {
	next := q.Enumerator()
	for t, ok := next().Get(); ok; t, ok = next().Get() {
		if pred(t) {
			return true
		}
	}
	return false
}

// Empty returns true if q has no elements.
func Empty[T any](q Query[T]) bool {
	t := q.Enumerator()()
	return !t.Valid()
}
