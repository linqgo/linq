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

// NewNonReplayableQuery returns a new query based on a function that returns
// enumerators. It is tagged as non-replayable to indicate that its values can
// only be enumerated once. This applies to consumable sources such as channels
// and io.Reader.
func NewNonReplayableQuery[T any](
	i func() Enumerator[T],
	nonReplayable bool,
) Query[T] {
	q := Query[T]{enumerator: i}
	if nonReplayable {
		q.extra = &queryExtra[T]{oneShot: true}
	}
	return q
}

// Enumerator returns an enumerator for q.
func (q Query[T]) Enumerator() Enumerator[T] {
	return q.enumerator()
}

func (q Query[T]) OneShot() bool {
	if q.extra == nil {
		return false
	}
	return q.extra.oneShot
}

func (q Query[T]) lesser() lesserFunc[T] {
	if q.extra == nil {
		return nil
	}
	return q.extra.lesser
}

// SetOneShot is for internal use when implementing a third-party Query. Some
// queries are consumed during enumeration and will thus be empty the second
// time you try to enumerate the (e.g., a query that reads a channel). You must
// tag such queries by passing true to this method. This also applies when
// consuming other queries, for example:
//
//	func Exclamate(q Query[string]) Query[string] {
//	    result := NewQuery(func() Enumerator[string] {
//	        next := q.Enumerator()
//	        return func() (string, bool) {
//	            if t, ok := next(); ok {
//	                return t + "!", true
//	            }
//	            return "", false
//	        }
//	    })
//	    result.SetOneShot(q.OneShot())
//	}
//
// Note that a non one-shot query doesn't guarantee that it won't consume its
// inputs. It only says that this query can be enumerated multiple times. For
// example, q.Memoize() is never one-shot, even though it may consume a one-shot
// q.
//
// The Pipe convenience function offers a simpler mechanism to implement Queries
// and will automatically tag the output query as one-shot if the input query
// is. Prefer it over NewQuery when it makes sense.
func (q *Query[T]) SetOneShot(oneShot bool) {
	q.extra = q.extra.withOneShot(oneShot)
}

func (q Query[T]) withOneShot(oneShot bool) Query[T] {
	q.SetOneShot(oneShot)
	return q
}

func (q Query[T]) withLesser(lesser lesserFunc[T]) Query[T] {
	q.extra = q.extra.withLesser(lesser)
	return q
}

type queryExtra[T any] struct {
	lesser  lesserFunc[T]
	oneShot bool
}

func (qe *queryExtra[T]) ifNonZero() *queryExtra[T] {
	if qe.lesser != nil || qe.oneShot {
		return qe
	}
	return nil
}

func (qe *queryExtra[T]) ifNeeded(needed bool) *queryExtra[T] {
	switch {
	case !needed:
		return qe
	case qe == nil:
		return &queryExtra[T]{}
	default:
		return qe
	}
}

func (qe *queryExtra[T]) withLesser(lesser lesserFunc[T]) *queryExtra[T] {
	if qe = qe.ifNeeded(lesser != nil); qe != nil {
		qe.lesser = lesser
		return qe.ifNonZero()
	}
	return nil
}

func (qe *queryExtra[T]) withOneShot(oneShot bool) *queryExtra[T] {
	if qe = qe.ifNeeded(oneShot); qe != nil {
		qe.oneShot = oneShot
		return qe.ifNonZero()
	}
	return nil
}

type lesserFunc[T any] func([]T) func(i, j int) bool

func newQueryFromEnumerator[T any](e Enumerator[T]) Query[T] {
	return Query[T]{enumerator: func() Enumerator[T] { return e }}
}
