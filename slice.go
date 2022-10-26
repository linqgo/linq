package linq

// ToSlice returns a slice containing the elements of q.
func (q Query[T]) ToSlice() []T {
	return ToSlice(q)
}

// ToSlice returns a slice containing the elements of q.
func ToSlice[T any](q Query[T]) []T {
	next := q.Enumerator()
	var ret []T
	for t, ok := next().Get(); ok; t, ok = next().Get() {
		ret = append(ret, t)
	}
	return ret
}
