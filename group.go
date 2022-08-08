package linq

// GroupBy returns a Query[KV[K, Query[T]]] with elements from q grouped using
// the specified key function.
func GroupBy[T any, K comparable](
	q Query[T],
	key func(t T) K,
) Query[KV[K, Query[T]]] {
	return GroupByI(q, indexify(key))
}

// GroupByI returns a Query[KV[K, Query[T]]] with elements from q grouped using
// the specified key function. The key function takes the index and value of
// each element.
func GroupByI[T any, K comparable](
	q Query[T],
	key func(i int, t T) K,
) Query[KV[K, Query[T]]] {
	return GroupBySelectI(q, keyIdentityI(key))
}

// GroupBySlices returns a Query[KV[K, []T]] with elements from q grouped using
// the specified key function.
func GroupBySlices[T any, K comparable](
	q Query[T],
	key func(t T) K,
) Query[KV[K, []T]] {
	return GroupBySlicesI(q, indexify(key))
}

// GroupBySlicesI returns a Query[KV[K, []T]] with elements from q grouped using
// the specified key function. The sel function takes the index and value of
// each element.
func GroupBySlicesI[T any, K comparable](
	q Query[T],
	key func(i int, t T) K,
) Query[KV[K, []T]] {
	return GroupBySelectSlicesI(q, keyIdentityI(key))
}

// GroupBySelect returns a Query[KV[K, Query[T]]] with elements from q grouped
// using the specified sel function, which produces a key/value pair for each
// source element.
func GroupBySelect[T, U any, K comparable](
	q Query[T],
	sel func(t T) KV[K, U],
) Query[KV[K, Query[U]]] {
	return GroupBySelectI(q, indexify(sel))
}

// GroupBySelectI returns a Query[KV[K, Query[T]]] with elements from q grouped
// using the specified sel function, which produces a key/value pair for each
// source element. The sel function takes the index and value of each element.
func GroupBySelectI[T, U any, K comparable](
	q Query[T],
	sel func(i int, t T) KV[K, U],
) Query[KV[K, Query[U]]] {
	return Select(
		GroupBySelectSlicesI(q, sel),
		func(kv KV[K, []U]) KV[K, Query[U]] {
			return NewKV(kv.Key, From(kv.Value...))
		},
	)
}

// GroupBySelectSlices returns a Query[KV[K, []T]] with elements from q grouped
// using the specified sel function, which produces a key/value pair for each
// source element.
func GroupBySelectSlices[T, U any, K comparable](
	q Query[T],
	sel func(t T) KV[K, U],
) Query[KV[K, []U]] {
	return GroupBySelectSlicesI(q, indexify(sel))
}

// GroupBySelectSlicesI returns a Query[KV[K, []T]] with elements from q grouped
// using the specified sel function, which produces a key/value pair for each
// source element. The sel function takes the index and value of each element.
func GroupBySelectSlicesI[T, U any, K comparable](
	q Query[T],
	sel func(i int, t T) KV[K, U],
) Query[KV[K, []U]] {
	return NewQuery(func() Enumerator[KV[K, []U]] {
		next := q.Enumerator()
		m := map[K][]U{}
		i := counter(0)
		for t, ok := next(); ok; t, ok = next() {
			kv := sel(i(), t)
			m[kv.Key] = append(m[kv.Key], kv.Value)
		}
		return FromMap(m).Enumerator()
	})
}

func keyIdentityI[T any, K comparable](key func(i int, t T) K) func(i int, t T) KV[K, T] {
	return func(i int, t T) KV[K, T] {
		return NewKV(key(i, t), t)
	}
}
