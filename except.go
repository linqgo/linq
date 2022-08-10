package linq

// Except returns all elements of a except those also found in b.
func Except[T comparable](a, b Query[T]) Query[T] {
	return ExceptBy(a, b, Identity[T])
}

// ExceptBy returns all elements of a except those whose key is found in b.
func ExceptBy[T, K comparable](
	a Query[T],
	b Query[K],
	key func(t T) K,
) Query[T] {
	return NewQuery(func() Enumerator[T] {
		s := setFrom(b.Enumerator())
		return a.Where(func(t T) bool { return !s.Has(key(t)) }).Enumerator()
	})
}
