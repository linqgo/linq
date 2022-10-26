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

func countWhenTaking[T any](q Query[T], take int) (count int, all bool) {
	if take == 0 {
		return 0, false
	}
	count = q.fastCount()
	if take < count {
		return take, false
	}
	return count, count >= 0
}

// Take returns a query with the first n elements of q.
func Take[T any](q Query[T], take int) Query[T] {
	count, all := countWhenTaking(q, take)
	if all {
		return q
	}
	if count == 0 {
		return None[T]()
	}
	var get Getter[T]
	if qget := q.getter(); qget != nil {
		get = func(i int) Maybe[T] {
			if 0 <= i && i < take {
				return qget(i)
			}
			return No[T]()
		}
	}
	return Pipe(q,
		func(next Enumerator[T]) Enumerator[T] {
			i := 0
			return func() Maybe[T] {
				if i < take {
					i++
					return next()
				}
				return No[T]()
			}
		},
		FastCountOption[T](count),
		FastGetOption(get),
	)
}

// TakeLast returns a query with the last n elements of q.
func TakeLast[T any](q Query[T], take int) Query[T] {
	if q.count >= 0 {
		return Skip(q, q.count-take)
	}
	count, _ := countWhenTaking(q, take)
	return Pipe(q,
		func(next Enumerator[T]) Enumerator[T] {
			buf := newBuffer(next, take)
			Drain(buf.Next)
			return buf.Enumerator()
		},
		FastCountOption[T](count),
	)
}

// TakeWhile returns a query that takes elements of q while pred returns true.
func TakeWhile[T any](q Query[T], pred func(t T) bool) Query[T] {
	return Pipe(q, func(next Enumerator[T]) Enumerator[T] {
		return func() Maybe[T] {
			if t, ok := next().Get(); ok && pred(t) {
				return Some(t)
			}
			return No[T]()
		}
	}, FastCountIfEmptyOption[T](q.fastCount()))
}
