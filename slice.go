package linq

// ToSlice returns a slice containing the elements of q.
func (q Query[T]) ToSlice() []T {
	return ToSlice(q)
}

// ToSlice returns a slice containing the elements of q.
func ToSlice[T any](q Query[T]) []T {
	next := q.Enumerator()
	var ret []T
	for t, ok := next(); ok; t, ok = next() {
		ret = append(ret, t)
	}
	return ret
}
