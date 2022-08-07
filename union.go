package linq

// Union returns the set union of a and b.
func Union[T comparable](a, b Query[T]) Query[T] {
	return a.Concat(Except(b, a))
}
