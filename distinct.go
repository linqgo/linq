package linq

// Distinct contains elements from an query with duplicates removed.
func Distinct[T comparable](q Query[T]) Query[T] {
	return DistinctBy(q, Identity[T])
}

// DistinctBy contains elements from a query with duplicates removed. A selector
// function produces values for comparison. E.g. for case-insensitive
// deduplication:
//
//	DistinctBy(names, strings.ToUpper)
func DistinctBy[T any, U comparable](q Query[T], sel func(t T) U) Query[T] {
	return NewQuery(func() Enumerator[T] {
		next := q.Enumerator()
		s := set[U]{}
		return func() (t T, ok bool) {
			for t, ok = next(); ok; t, ok = next() {
				if u := sel(t); !s.Has(u) {
					s.Add(u)
					return t, ok
				}
			}
			return t, ok
		}
	})
}
