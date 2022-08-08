package linq

// Skip returns a query all elements of q except the first n.
func (q Query[T]) Skip(n int) Query[T] {
	return Skip(q, n)
}

// SkipLast returns a query all elements of q except the last n.
func (q Query[T]) SkipLast(n int) Query[T] {
	return SkipLast(q, n)
}

// SkipWhile returns a query that skips elements of q while pred returns true.
func (q Query[T]) SkipWhile(pred func(t T) bool) Query[T] {
	return SkipWhile(q, pred)
}

// SkipWhileI returns a query that skips elements of q while pred returns true.
// The pred function takes the index and value of each element.
func (q Query[T]) SkipWhileI(pred func(i int, t T) bool) Query[T] {
	return SkipWhileI(q, pred)
}

// Skip returns a query all elements of q except the first n.
func Skip[T any](q Query[T], n int) Query[T] {
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
func SkipLast[T any](q Query[T], n int) Query[T] {
	if n == 0 {
		return q
	}
	return NewQuery(func() Enumerator[T] {
		return newBuffer(q.Enumerator(), n).Next
	})
}

// SkipWhile returns a query that skips elements of q while pred returns true.
func SkipWhile[T any](q Query[T], pred func(t T) bool) Query[T] {
	return SkipWhileI(q, indexify(pred))
}

// SkipWhileI returns a query that skips elements of q while pred returns true.
// The pred function takes the index and value of each element.
func SkipWhileI[T any](q Query[T], pred func(i int, t T) bool) Query[T] {
	return NewQuery(func() Enumerator[T] {
		next := q.Enumerator()
		i := counter(0)
		for t, ok := next(); ok; t, ok = next() {
			if !pred(i(), t) {
				return concatEnumerators(valueEnumerator(t), next)
			}
		}
		return noneEnumerator[T]
	})
}
