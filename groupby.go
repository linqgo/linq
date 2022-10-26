package linq

// GroupBy returns a Query[KV[K, Query[T]]] with elements from q grouped using
// the specified key function.
func GroupBy[T any, K comparable](
	q Query[T],
	key func(t T) K,
) Query[KV[K, Query[T]]] {
	return GroupBySelect(q, keyIdentity(key))
}

// GroupBySlices returns a Query[KV[K, []T]] with elements from q grouped using
// the specified key function.
func GroupBySlices[T any, K comparable](
	q Query[T],
	key func(t T) K,
) Query[KV[K, []T]] {
	return GroupBySelectSlices(q, keyIdentity(key))
}

// GroupBySelect returns a Query[KV[K, Query[T]]] with elements from q grouped
// using the specified sel function, which produces a key/value pair for each
// source element.
func GroupBySelect[T, U any, K comparable](
	q Query[T],
	sel func(t T) KV[K, U],
) Query[KV[K, Query[U]]] {
	return Select(
		GroupBySelectSlices(q, sel),
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
	return Pipe(q, func(next Enumerator[T]) Enumerator[KV[K, []U]] {
		m := map[K][]U{}
		for t, ok := next().Get(); ok; t, ok = next().Get() {
			kv := sel(t)
			m[kv.Key] = append(m[kv.Key], kv.Value)
		}
		return FromMap(m).Enumerator()
	}, FastCountIfEmptyOption[KV[K, []U]](q.fastCount()))
}

func keyIdentity[T any, K comparable](key func(t T) K) func(t T) KV[K, T] {
	return func(t T) KV[K, T] {
		return NewKV(key(t), t)
	}
}
