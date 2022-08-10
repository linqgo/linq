package linq

func (q Query[T]) Indexed() Query[KV[int, T]] {
	return Indexed(q)
}

func (q Query[T]) IndexedFrom(start int) Query[KV[int, T]] {
	return IndexedFrom(q, 0)
}

func Indexed[T any](q Query[T]) Query[KV[int, T]] {
	return IndexedFrom(q, 0)
}

func IndexedFrom[T any](q Query[T], start int) Query[KV[int, T]] {
	i := counter(start)
	return Select(q, func(t T) KV[int, T] {
		return NewKV(i(), t)
	})
}
