package linq

func (q Query[T]) Scan() func(p *T) bool {
	return Scan(q)
}

// Scan returns a function that scans entries from q, setting the passed
// pointer parameter to each element and returning true until it runs out of
// elements. Then it returns false.
func Scan[T any](q Query[T]) func(p *T) bool {
	next := q.Enumerator()
	return func(p *T) bool {
		t, ok := next().Get()
		if ok {
			*p = t
		}
		return ok
	}
}
