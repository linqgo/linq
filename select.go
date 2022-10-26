package linq

// Select returns a query with the elements of q transformed by sel.
//
// Caveat: The output must be of the same type. For transforms to different
// types, use the corresponding free function.
func (q Query[T]) Select(sel func(t T) T) Query[T] {
	return Select(q, sel)
}

// Select returns a query with the elements of q transformed by sel.
func Select[T, U any](q Query[T], sel func(t T) U) Query[U] {
	var get Getter[U]
	if qget := q.getter(); qget != nil {
		get = func(i int) Maybe[U] {
			if t, ok := qget(i).Get(); ok {
				return Some(sel(t))
			}
			return No[U]()
		}
	}
	return PipeOneToOne(q, func() func(t T) U { return sel }, FastGetOption(get))
}

// SelectMany projects each element of q to a subquery and flattens the
// subqueries into a single query.
func SelectMany[T, U any](q Query[T], project func(t T) Query[U]) Query[U] {
	return Pipe(q, func(next Enumerator[T]) Enumerator[U] {
		var t *T
		var tNext Enumerator[U]
		return func() Maybe[U] {
			var u U
			ok := false
			for !ok {
				if t == nil {
					var v T
					t = &v
					if *t, ok = next().Get(); !ok {
						return No[U]()
					}
					tNext = project(*t).Enumerator()
				}

				u, ok = tNext().Get()
				if !ok {
					t = nil
				}
			}
			return NewMaybe(u, ok)
		}
	}, FastCountIfEmptyOption[U](q.fastCount()))
}
