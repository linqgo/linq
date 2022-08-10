package linq

func Index[T any](q Query[T]) Query[KV[int, T]] {
	return IndexFrom(q, 0)
}

func IndexFrom[T any](q Query[T], start int) Query[KV[int, T]] {
	i := counter(start)
	return Select(q, func(t T) KV[int, T] {
		return NewKV(i(), t)
	})
}
