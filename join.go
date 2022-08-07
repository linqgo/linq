package linq

// Join returns the join of q1 and q2. selKey1 and selKey2 produce keys from
// elements of q1 and q2, respectively. Element pairs with the same key are
// passed to selResult to produce output elements.
func Join[T1, T2, U any, K comparable](
	q1 Query[T1],
	q2 Query[T2],
	selKey1 func(t1 T1) K,
	selKey2 func(t2 T2) K,
	selResult func(t1 T1, t2 T2) U,
) Query[U] {
	next2 := q2.Enumerator()
	lup2 := map[K][]T2{}
	for t2, ok := next2(); ok; t2, ok = next2() {
		k := selKey2(t2)
		lup2[k] = append(lup2[k], t2)
	}

	return SelectMany(q1, func(t1 T1) Query[U] {
		k := selKey1(t1)
		t2 := lup2[k]
		return Select(From(t2...), func(t2 T2) U {
			return selResult(t1, t2)
		})
	})
}
