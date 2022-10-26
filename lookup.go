package linq

type lookupBuilder[T any, K comparable] struct {
	next Enumerator[T]
	key  func(T) K
	lup  map[K][]T
}

func newLookupBuilder[T any, K comparable](q Query[T], key func(T) K) *lookupBuilder[T, K] {
	return &lookupBuilder[T, K]{
		next: q.Enumerator(),
		key:  key,
		lup:  map[K][]T{},
	}
}

func buildLookup[T any, K comparable](q Query[T], key func(T) K) map[K][]T {
	b := newLookupBuilder(q, key)
	for b.Next() { //nolint:revive
	}
	return b.Lookup()
}

func (b *lookupBuilder[T, K]) Next() bool {
	if t, ok := b.next().Get(); ok {
		k := b.key(t)
		b.lup[k] = append(b.lup[k], t)
		return true
	}
	return false
}

func (b *lookupBuilder[T, K]) Lookup() map[K][]T {
	return b.lup
}

func (b *lookupBuilder[T, K]) Requery() Query[T] {
	return Concat(
		SelectMany(FromMap(b.lup), func(kv KV[K, []T]) Query[T] {
			return From(kv.Value...)
		}),
		newQueryFromEnumerator(b.next),
	)
}
