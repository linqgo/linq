package linq

// Take returns a query with the first n elements of q.
func (q Query[T]) Take(count int) Query[T] {
	if count == 0 {
		return None[T]()
	}
	return NewQuery(func() Enumerator[T] {
		next := q.Enumerator()
		return func() (_ T, _ bool) {
			if count > 0 {
				count -= 1
				return next()
			}
			return
		}
	})
}

// TakeLast returns a query with the last n elements of q.
func (q Query[T]) TakeLast(count int) Query[T] {
	if count == 0 {
		return None[T]()
	}
	return NewQuery(func() Enumerator[T] {
		buf := newBuffer(q.Enumerator(), count)
		drain(buf.Next)
		return buf.Enumerator()
	})
}

// TakeWhile returns a query that takes elements of q while pred returns true.
func (q Query[T]) TakeWhile(pred func(t T) bool) Query[T] {
	return NewQuery(func() Enumerator[T] {
		next := q.Enumerator()
		return func() (_ T, _ bool) {
			if t, ok := next(); ok && pred(t) {
				return t, ok
			}
			return
		}
	})
}
