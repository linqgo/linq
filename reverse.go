package linq

// Reverse returns a query with the elements of q in reverse.
func (q Query[T]) Reverse() Query[T] {
	return Reverse(q)
}

// Reverse returns a query with the elements of q in reverse.
func Reverse[T any](q Query[T]) Query[T] {
	return NewQuery(func() Enumerator[T] {
		data := q.ToSlice()
		return func() (T, bool) {
			var e T
			last := len(data) - 1
			if last >= 0 {
				e, data = data[last], data[:last]
			}
			return e, last >= 0
		}
	})
}
