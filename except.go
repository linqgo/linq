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
	return ExceptByI(a, b, indexify(key))
}

// ExceptByI returns all elements of a except those whose key is found in b. The
// key function takes the index and value of each element.
func ExceptByI[T, K comparable](
	a Query[T],
	b Query[K],
	key func(i int, t T) K,
) Query[T] {
	return NewQuery(func() Enumerator[T] {
		s := setFrom(b.Enumerator())
		return a.WhereI(func(i int, t T) bool {
			return !s.Has(key(i, t))
		}).Enumerator()
	})
}
