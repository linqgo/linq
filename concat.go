package linq

// Concat returns the concatenation of q and r. Enumerating it enumerates the
// elements of each Query in turn.
func (q Query[T]) Concat(r Query[T]) Query[T] {
	return Concat(q, r)
}

// Concat returns the concatenation of queries. Enumerating it enumerates the
// elements of each Query in turn.
func Concat[T any](queries ...Query[T]) Query[T] {
	oneshot := false
	for _, q := range queries {
		if q.OneShot() {
			oneshot = true
			break
		}
	}

	nonempty := 0
	count := 0
	for i, q := range queries {
		c := q.fastCount()
		if c != 0 {
			if nonempty < i {
				queries[nonempty] = q
			}
			nonempty++
			if c < 0 {
				count = -1
			}
		}
		if count >= 0 {
			count += c
		}
	}

	// Exactly one non-empty input?
	if nonempty == 1 {
		return queries[0]
	}
	queries = queries[:nonempty]

	return NewQuery(func() Enumerator[T] {
		enumerators := make([]Enumerator[T], 0, len(queries))
		for _, q := range queries {
			enumerators = append(enumerators, q.Enumerator())
		}
		return concatEnumerators(enumerators...)
	}, OneShotOption[T](oneshot), FastCountOption[T](count))
}

func concatEnumerators[T any](nexts ...Enumerator[T]) Enumerator[T] {
	next := No[T]
	return func() Maybe[T] {
		for {
			if t := next(); t.Valid() {
				return t
			}
			if len(nexts) == 0 {
				return No[T]()
			}
			next, nexts = nexts[0], nexts[1:]
		}
	}
}
