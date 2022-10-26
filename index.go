package linq

func Index[T any](q Query[T]) Query[KV[int, T]] {
	return IndexFrom(q, 0)
}

func IndexFrom[T any](q Query[T], start int) Query[KV[int, T]] {
	var get Getter[KV[int, T]]
	if qget := q.getter(); qget != nil {
		get = func(i int) Maybe[KV[int, T]] {
			if t, ok := qget(i).Get(); ok {
				return Some(NewKV(start+i, t))
			}
			return No[KV[int, T]]()
		}
	}
	return PipeOneToOne(q,
		func() func(t T) KV[int, T] {
			i := start - 1
			return func(t T) KV[int, T] {
				i++
				return NewKV(i, t)
			}
		},
		FastGetOption(get),
	)
}
