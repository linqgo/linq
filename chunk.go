package linq

// Chunk returns the elements of q in queries containing chunks of the specified
// size.
func Chunk[T any](q Query[T], size int) Query[Query[T]] {
	return NewQuery(func() Enumerator[Query[T]] {
		next := q.Enumerator()
		return func() (Query[T], bool) {
			chunk := make([]T, 0, size)
			for i := 0; i < size; i++ {
				t, ok := next()
				if !ok {
					next = noneEnumerator[T]
					return From(chunk...), len(chunk) > 0
				}
				chunk = append(chunk, t)
			}
			return From(chunk...), true
		}
	})
}

// ChunkSlices returns the elements of q in slices of the specified size.
func ChunkSlices[T any](q Query[T], size int) Query[[]T] {
	return Select(Chunk(q, size), func(c Query[T]) []T { return c.ToSlice() })
}