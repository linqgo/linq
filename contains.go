package linq

// Contains returns true if and only if t is an element in q.
func Contains[T comparable](q Query[T], t T) bool {
	next := q.Enumerator()
	for e, ok := next().Get(); ok; e, ok = next().Get() {
		if e == t {
			return true
		}
	}
	return false
}
