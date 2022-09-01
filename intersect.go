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
	return NewQuery(func() Enumerator[T] {
		s := setFrom(b.Enumerator())
		return a.Where(func(t T) bool { return s.Has(key(t)) }).Enumerator()
	}).withOneShot(a.OneShot() || b.OneShot())
}
