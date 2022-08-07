package linq

// Append returns a query with the elements of q followed by the elements of t.
func (q Query[T]) Append(t ...T) Query[T] {
	return Append(q, t...)
}

// Append returns a query with the elements of q followed by the elements of t.
func Append[T any](q Query[T], t ...T) Query[T] {
	return q.Concat(From(t...))
}
