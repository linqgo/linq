package linq

// Query represents a query that can be enumerated. This is the main Linq
// object, with many methods defined against it. Most Linq functions take and
// return instances of this type.
type Query[T any] struct {
	enumerator func() Enumerator[T]
	extra      *queryExtra[T]
}

// NewQuery returns a new query based on a function that returns enumerators.
func NewQuery[T any](i func() Enumerator[T]) Query[T] {
	return Query[T]{enumerator: i}
}

// Enumerator returns an enumerator for q.
func (q Query[T]) Enumerator() Enumerator[T] {
	return q.enumerator()
}

func (q Query[T]) lesser() lesserFunc[T] {
	if q.extra == nil {
		return nil
	}
	return q.extra.lesser
}

// func (q Query[T]) replayable() bool {
// 	if q.extra == nil {
// 		return true
// 	}
// 	return !q.extra.nonReplayable
// }

type queryExtra[T any] struct {
	lesser        lesserFunc[T]
	nonReplayable bool
}

type lesserFunc[T any] func([]T) func(i, j int) bool

func newQueryFromEnumerator[T any](e Enumerator[T]) Query[T] {
	return Query[T]{enumerator: func() Enumerator[T] { return e }}
}
