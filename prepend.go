package linq

// Prepend returns a query with the elements of t followed by the elements of q.
//
// Be careful! Prepend puts the tail arguments, t, in front of q. To avoid
// confusion and bugs, consider using the corresponding global Prepend function.
func (q Query[T]) Prepend(t ...T) Query[T] {
	return Prepend(t...)(q)
}

// Prepend returns a query with the elements of t followed by the elements of q.
//
// So that t... args appear before q, Prepend takes just t and returns a func
// that takes q.
func Prepend[T any](t ...T) func(q Query[T]) Query[T] {
	return func(q Query[T]) Query[T] {
		return From(t...).Concat(q)
	}
}
