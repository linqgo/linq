package linq

type QueryOption[T any] func(q *queryExtra[T], count *int)

// Query represents a query that can be enumerated. This is the main Linq
// object, with many methods defined against it. Most Linq functions take and
// return instances of this type.
type Query[T any] struct {
	enumerator func() Enumerator[T]
	count      int
	extra      *queryExtra[T]
}

// NewQuery returns a new query based on a function that returns enumerators.
func NewQuery[T any](i func() Enumerator[T], options ...QueryOption[T]) Query[T] {
	q := Query[T]{enumerator: i, count: -1}

	for _, option := range options {
		if option != nil {
			q.extra = &queryExtra[T]{}
			for _, option := range options {
				if option != nil {
					option(q.extra, &q.count)
				}
			}
			break
		}
	}
	if q.extra != nil && q.count == 0 {
		return None[T]()
	}
	return q
}

// Enumerator returns an enumerator for q.
func (q Query[T]) Enumerator() Enumerator[T] {
	return q.enumerator()
}

func (q Query[T]) OneShot() bool {
	return q.extra != nil && q.extra.oneShot
}

func (q Query[T]) lesser() lesserFunc[T] {
	if q.extra != nil {
		return q.extra.lesser
	}
	return nil
}

func (q Query[T]) getter() Getter[T] {
	if q.extra != nil {
		return q.extra.get
	}
	return nil
}

func (q *Query[T]) fastCount() int {
	return q.count
}

// FastCountOption is used when implementing a third-party Query.
// If a query's count can be determined in O(1) time, the value may be supplied as a FastCountOption to NewQuery.
// This will be used by (Query).FastCount &co.
func FastCountOption[T any](count int) QueryOption[T] {
	if count >= 0 {
		return func(e *queryExtra[T], c *int) {
			*c = count
		}
	}
	return nil
}

// FastCountIfEmptyOption is used when implementing a third-party Query.
// If an empty input query produces an empty output query, use FastCountIfEmptyOption to report a FastCount of 0 for the output query.
// This will be used by (Query).FastCount &co.
func FastCountIfEmptyOption[T any](count int) QueryOption[T] {
	if count != 0 {
		return nil
	}
	return FastCountOption[T](0)
}

// FastGetOption is used when implement a third-party Query.
// If any element can be accessed by index in O(1) time, the accessor may be
// supplied via this option.
func FastGetOption[T any](get Getter[T]) QueryOption[T] {
	if get != nil {
		return func(q *queryExtra[T], count *int) {
			q.get = get
		}
	}
	return nil
}

// OneShotOption is used when implementing a third-party Query.
// Some queries are consumed during enumeration and will thus be empty the second time you try to enumerate them (e.g., a query that reads a channel).
// You must tag such queries by passing true to this method.
// This also applies when consuming other queries, for example:
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
//	    result.SetOneShot(q.OneShotOption())
//	}
//
// Note that a non one-shot query doesn't guarantee that it won't consume its inputs.
// It only says that this query can be enumerated multiple times.
// For example, q.Memoize() is never one-shot, even though it may consume a one-shot
// q.
//
// The Pipe convenience function offers a simpler mechanism to implement Queries and will automatically tag the output query as one-shot if the input query is.
// Prefer it over NewQuery when it makes sense.
func OneShotOption[T any](oneShot bool) QueryOption[T] {
	if oneShot {
		return func(e *queryExtra[T], _ *int) {
			e.oneShot = oneShot
		}
	}
	return nil
}

func ComputedFastCountOption[T any](
	count int,
	compute func(count int) int,
) QueryOption[T] {
	if count >= 0 {
		n := compute(count)
		return func(e *queryExtra[T], c *int) { *c = n }
	}
	return nil
}

func LesserOption[T any](lesser lesserFunc[T]) QueryOption[T] {
	return func(e *queryExtra[T], _ *int) {
		e.lesser = lesser
	}
}

type queryExtra[T any] struct {
	lesser  lesserFunc[T]
	get     Getter[T]
	oneShot bool
}

type lesserFunc[T any] func([]T) func(i, j int) bool

func newQueryFromEnumerator[T any](e Enumerator[T]) Query[T] {
	return NewQuery(func() Enumerator[T] { return e })
}
