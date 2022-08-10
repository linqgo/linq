package linq

// Join returns the join of q1 and q2. selKey1 and selKey2 produce keys from
// elements of q1 and q2, respectively. Element pairs with the same key are
// passed to selResult to produce output elements.
func Join[A, B, R any, K comparable]( //nolint:revive
	a Query[A],
	b Query[B],
	selKeyA func(t1 A) K,
	selKeyB func(t2 B) K,
	selResult func(t1 A, t2 B) R,
) Query[R] {
	return NewQuery(func() Enumerator[R] {
		nextA := a.Enumerator()
		nextB := b.Enumerator()
		lupA := map[K][]A{}
		lupB := map[K][]B{}

		// Scan both inputs till one runs out. The exhausted input's map will be
		// used for lookups. The other side will be repackaged into a new query
		// for full traversal. This includes the entries already loaded into the
		// now unneeded lookup and the values remaining in the enumerator.
		for {
			tA, okA := nextA()
			if okA {
				k := selKeyA(tA)
				lupA[k] = append(lupA[k], tA)
			}

			tB, okB := nextB()
			if okB {
				k := selKeyB(tB)
				lupB[k] = append(lupB[k], tB)
			}

			switch {
			case !okA:
				return SelectMany(repackLup(lupB, nextB), func(t2 B) Query[R] {
					return Select(From(lupA[selKeyB(t2)]...), func(t1 A) R {
						return selResult(t1, t2)
					})
				}).Enumerator()
			case !okB:
				return SelectMany(repackLup(lupA, nextA), func(t1 A) Query[R] {
					return Select(From(lupB[selKeyA(t1)]...), func(t2 B) R {
						return selResult(t1, t2)
					})
				}).Enumerator()
			}
		}
	})
}

func repackLup[K comparable, T any](lup map[K][]T, next Enumerator[T]) Query[T] {
	return Concat(
		SelectMany(FromMap(lup), func(kv KV[K, []T]) Query[T] {
			return From(kv.Value...)
		}),
		newQueryFromEnumerator(next),
	)
}
