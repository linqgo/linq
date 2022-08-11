package linq

// Query represents a query that can be enumerated. This is the main Linq
// object, with many methods defined against it. Most Linq functions take and
// return instances of this type.
type Query[T any] struct {
	enumerator func() Enumerator[T]
	lesser     lesserFunc[T]
}

// NewQuery returns a new query based on a function that returns enumerators.
func NewQuery[T any](i func() Enumerator[T]) Query[T] {
	return Query[T]{enumerator: i}
}

// Enumerator returns an enumerator for q.
func (q Query[T]) Enumerator() Enumerator[T] {
	return q.enumerator()
}

type lesserFunc[T any] func([]T) func(i, j int) bool

func newQueryFromEnumerator[T any](e Enumerator[T]) Query[T] {
	return Query[T]{enumerator: func() Enumerator[T] { return e }}
}
