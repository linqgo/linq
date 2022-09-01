package linq

func Index[T any](q Query[T]) Query[KV[int, T]] {
	return IndexFrom(q, 0)
}

func IndexFrom[T any](q Query[T], start int) Query[KV[int, T]] {
	return Pipe(q, func(next Enumerator[T]) Enumerator[KV[int, T]] {
		i := start - 1
		return func() (kv KV[int, T], ok bool) {
			if t, ok := next(); ok {
				i++
				return NewKV(i, t), true
			}
			return kv, false
		}
	})
}
