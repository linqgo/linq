package linq

// Concat returns the concatenation of q and r. Enumerating it enumerates the
// elements of each Query in turn.
func (q Query[T]) Concat(r Query[T]) Query[T] {
	return Concat(q, r)
}

// Concat returns the concatenation of queries. Enumerating it enumerates the
// elements of each Query in turn.
func Concat[T any](queries ...Query[T]) Query[T] {
	return NewQuery(func() Enumerator[T] {
		enumerators := make([]Enumerator[T], 0, len(queries))
		for _, q := range queries {
			enumerators = append(enumerators, q.Enumerator())
		}
		return concatEnumerators(enumerators...)
	})
}

func concatEnumerators[T any](nexts ...Enumerator[T]) Enumerator[T] {
	next := noneEnumerator[T]
	return func() (t T, ok bool) {
		if t, ok = next(); ok {
			return
		}
		if len(nexts) > 0 {
			next, nexts = nexts[0], nexts[1:]
			return next()
		}
		return
	}
}
