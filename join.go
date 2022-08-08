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
	return JoinI(
		q1, q2,
		indexify(selKey1),
		indexify(selKey2),
		func(_ int, t1 T1, _ int, t2 T2) U {
			return selResult(t1, t2)
		},
	)
}

// JoinI returns the join of q1 and q2. selKey1 and selKey2 produce keys from
// elements of q1 and q2, respectively. Element pairs with the same key are
// passed to selResult to produce output elements. The sel... functions take the
// index and value of each element.
func JoinI[T1, T2, U any, K comparable](
	q1 Query[T1],
	q2 Query[T2],
	selKey1 func(i int, t1 T1) K,
	selKey2 func(i int, t2 T2) K,
	selResult func(i1 int, t1 T1, i2 int, t2 T2) U,
) Query[U] {
	next2 := q2.Enumerator()
	lup2 := map[K][]KV[int, T2]{}
	i2 := counter(0)
	for t2, ok := next2(); ok; t2, ok = next2() {
		i2 := i2()
		k := selKey2(i2, t2)
		lup2[k] = append(lup2[k], NewKV(i2, t2))
	}

	return SelectManyI(q1, func(i1 int, t1 T1) Query[U] {
		k := selKey1(i1, t1)
		it2 := lup2[k]
		return Select(From(it2...), func(it2 KV[int, T2]) U {
			return selResult(i1, t1, it2.Key, it2.Value)
		})
	})
}
