package linq

// Select returns a query with the elements of q transformed by sel.
func Select[T, U any](q Query[T], sel func(t T) U) Query[U] {
	return SelectI(q, indexify(sel))
}

// SelectI returns a query with the elements of q transformed by sel. The sel
// function takes the index and value of each element.
func SelectI[T, U any](q Query[T], sel func(i int, t T) U) Query[U] {
	return NewQuery(func() Enumerator[U] {
		next := q.Enumerator()
		i := counter(0)
		return func() (U, bool) {
			if t, ok := next(); ok {
				return sel(i(), t), true
			}
			var u U
			return u, false
		}
	})
}

// SelectMany projects each element of q to a subquery and flattens the
// subqueries into a single query.
func SelectMany[T, U any](q Query[T], project func(t T) Query[U]) Query[U] {
	return SelectManyI(q, indexify(project))
}

// SelectManyI projects each element of q to a subquery and flattens the
// subqueries into a single query. The project function takes the index and
// value of each element.
func SelectManyI[T, U any](q Query[T], project func(i int, t T) Query[U]) Query[U] {
	return NewQuery(func() Enumerator[U] {
		next := q.Enumerator()
		var t *T
		var tNext Enumerator[U]
		i := counter(0)
		return func() (u U, ok bool) {
			for !ok {
				if t == nil {
					var v T
					t = &v
					if *t, ok = next(); !ok {
						return u, ok
					}
					tNext = project(i(), *t).Enumerator()
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
