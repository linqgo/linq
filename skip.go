package linq

// Skip returns a query all elements of q except the first n.
func (q Query[T]) Skip(n int) Query[T] {
	if n == 0 {
		return q
	}
	return NewQuery(func() Enumerator[T] {
		next := q.Enumerator()
		for i := 0; i < n; i++ {
			if _, ok := next(); !ok {
				return noneEnumerator[T]
			}
		}
		return next
	})
}

// SkipLast returns a query all elements of q except the last n.
func (q Query[T]) SkipLast(n int) Query[T] {
	if n == 0 {
		return q
	}
	return NewQuery(func() Enumerator[T] {
		return newBuffer(q.Enumerator(), n).Next
	})
}

// Skip returns a query that skips elements of q while pred returns true.
func (q Query[T]) SkipWhile(pred func(t T) bool) Query[T] {
	return NewQuery(func() Enumerator[T] {
		next := q.Enumerator()
		for t, ok := next(); ok; t, ok = next() {
			if !pred(t) {
				return concatEnumerators(valueEnumerator(t), next)
			}
		}
		return noneEnumerator[T]
	})
}
