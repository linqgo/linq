package linq

// Select returns a query with the elements of q transformed by sel.
func Select[T, U any](q Query[T], sel func(t T) U) Query[U] {
	return NewQuery(func() Enumerator[U] {
		next := q.Enumerator()
		return func() (U, bool) {
			if t, ok := next(); ok {
				return sel(t), true
			}
			var u U
			return u, false
		}
	})
}

// SelectMany projects each element of q to a subquery and flattens the
// subqueries into a single query.
func SelectMany[T, U any](q Query[T], project func(t T) Query[U]) Query[U] {
	return NewQuery(func() Enumerator[U] {
		next := q.Enumerator()
		var t *T
		var tNext Enumerator[U]
		return func() (u U, ok bool) {
			for !ok {
				if t == nil {
					var v T
					t = &v
					if *t, ok = next(); !ok {
						return u, ok
					}
					tNext = project(*t).Enumerator()
				}

				u, ok = tNext()
				if !ok {
					t = nil
				}
			}
			return u, ok
		}
	})
}
