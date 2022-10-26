package linq

// Chunk returns the elements of q in queries containing chunks of the specified
// size.
func Chunk[T any](q Query[T], size int) Query[Query[T]] {
	var get Getter[Query[T]]
	if qget := q.getter(); qget != nil {
		get = func(i int) Maybe[Query[T]] {
			start := size * i
			if _, ok := qget(start).Get(); ok {
				return Some(q.Skip(start).Take(size))
			}
			return No[Query[T]]()
		}
	}
	return Pipe(q,
		func(next Enumerator[T]) Enumerator[Query[T]] {
			return func() Maybe[Query[T]] {
				chunk := make([]T, 0, size)
				for i := 0; i < size; i++ {
					t, ok := next().Get()
					if !ok {
						next = No[T]
						return NewMaybe(From(chunk...), len(chunk) > 0)
					}
					chunk = append(chunk, t)
				}
				return Some(From(chunk...))
			}
		},
		ComputedFastCountOption[Query[T]](q.fastCount(), func(count int) int {
			return (count-1)/size + 1
		}),
		FastGetOption(get),
	)
}

// ChunkSlices returns the elements of q in slices of the specified size.
func ChunkSlices[T any](q Query[T], size int) Query[[]T] {
	return Select(Chunk(q, size), func(c Query[T]) []T { return c.ToSlice() })
}
