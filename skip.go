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

func skipCount[T any](q Query[T], n int) int {
	count := q.fastCount()
	switch {
	case count > n:
		return count - n
	case count >= 0:
		return 0
	default:
		return count
	}
}

// Skip returns a query all elements of q except the first n.
func Skip[T any](q Query[T], n int) Query[T] {
	if n == 0 {
		return q
	}
	return Pipe(q, func(next Enumerator[T]) Enumerator[T] {
		for i := 0; i < n; i++ {
			if _, ok := next(); !ok {
				return noneEnumerator[T]
			}
		}
		return next
	}, FastCountOption[T](skipCount(q, n)))
}

// SkipLast returns a query all elements of q except the last n.
func SkipLast[T any](q Query[T], n int) Query[T] {
	if n == 0 {
		return q
	}
	return Pipe(q, func(next Enumerator[T]) Enumerator[T] {
		return newBuffer(next, n).Next
	}, FastCountOption[T](skipCount(q, n)))
}

// SkipWhile returns a query that skips elements of q while pred returns true.
func SkipWhile[T any](q Query[T], pred func(t T) bool) Query[T] {
	return Pipe(q, func(next Enumerator[T]) Enumerator[T] {
		for t, ok := next(); ok; t, ok = next() {
			if !pred(t) {
				return concatEnumerators(valueEnumerator(t), next)
			}
		}
		return noneEnumerator[T]
	}, FastCountIfEmptyOption[T](q.fastCount()))
}
