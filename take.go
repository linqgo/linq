package linq

// Take returns a query with the first n elements of q.
func (q Query[T]) Take(count int) Query[T] {
	return Take(q, count)
}

// TakeLast returns a query with the last n elements of q.
func (q Query[T]) TakeLast(count int) Query[T] {
	return TakeLast(q, count)
}

// TakeWhile returns a query that takes elements of q while pred returns true.
func (q Query[T]) TakeWhile(pred func(t T) bool) Query[T] {
	return TakeWhile(q, pred)
}

// Take returns a query with the first n elements of q.
func Take[T any](q Query[T], n int) Query[T] {
	count, all := takeCount(q, n)
	if all {
		return q
	}
	return Pipe(q, func(next Enumerator[T]) Enumerator[T] {
		i := 0
		return func() (T, bool) {
			if i < n {
				i++
				return next()
			}
			var t T
			return t, false
		}
	}, FastCountOption[T](count))
}

// TakeLast returns a query with the last n elements of q.
func TakeLast[T any](q Query[T], n int) Query[T] {
	count, all := takeCount(q, n)
	if all {
		return q
	}
	return Pipe(q, func(next Enumerator[T]) Enumerator[T] {
		buf := newBuffer(next, n)
		Drain(buf.Next)
		return buf.Enumerator()
	}, FastCountOption[T](count))
}

// TakeWhile returns a query that takes elements of q while pred returns true.
func TakeWhile[T any](q Query[T], pred func(t T) bool) Query[T] {
	return Pipe(q, func(next Enumerator[T]) Enumerator[T] {
		return func() (t T, ok bool) {
			if t, ok := next(); ok && pred(t) {
				return t, ok
			}
			return t, ok
		}
	}, FastCountIfEmptyOption[T](q.fastCount()))
}

func takeCount[T any](q Query[T], n int) (count int, all bool) {
	if n == 0 {
		return 0, false
	}
	count = q.fastCount()
	if n < count {
		return n, false
	}
	return count, count >= 0
}
