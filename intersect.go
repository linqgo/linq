package linq

// Intersect returns the set intersection of a and b.
func Intersect[T comparable](a, b Query[T]) Query[T] {
	return IntersectBy(a, b, Identity[T])
}

// IntersectBy returns the set intersection of a and b.
func IntersectBy[T, K comparable](
	a Query[T],
	b Query[K],
	key func(t T) K,
) Query[T] {
	return IntersectByI(a, b, indexify(key))
}

// IntersectByI returns the set intersection of a and b. Values from a are
// transformed through key to produce comparison values. The key function takes
// the index and value of each element.
func IntersectByI[T, K comparable](
	a Query[T],
	b Query[K],
	key func(i int, t T) K,
) Query[T] {
	return NewQuery(func() Enumerator[T] {
		s := setFrom(b.Enumerator())
		return a.WhereI(func(i int, t T) bool {
			return s.Has(key(i, t))
		}).Enumerator()
	})
}
