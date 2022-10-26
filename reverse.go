package linq

// Reverse returns a query with the elements of q in reverse.
func (q Query[T]) Reverse() Query[T] {
	return Reverse(q)
}

// Reverse returns a query with the elements of q in reverse.
func Reverse[T any](q Query[T]) Query[T] {
	var get Getter[T]
	if q.count >= 0 {
		if qget := q.getter(); qget != nil {
			last := q.count - 1
			return FromArray(ArrayFromLenGet(q.count, func(i int) T {
				return qget(last - i).Must()
			}))
		}
	}
	return NewQuery(
		func() Enumerator[T] {
			data := q.ToSlice()
			return func() Maybe[T] {
				var e T
				last := len(data) - 1
				if last >= 0 {
					e, data = data[last], data[:last]
				}
				return NewMaybe(e, last >= 0)
			}
		},
		OneShotOption[T](q.OneShot()),
		FastCountOption[T](q.fastCount()),
		FastGetOption(get),
	)
}
