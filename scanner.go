package linq

func (q Query[T]) Scanner() func(p *T) bool {
	return Scanner(q)
}

// Scanner returns a function that scans entries from q, setting the passed
// pointer parameter to each element and returning true until it runs out of
// elements. Then it returns false.
func Scanner[T any](q Query[T]) func(p *T) bool {
	next := q.Enumerator()
	return func(p *T) bool {
		t, ok := next().Get()
		if ok {
			*p = t
		}
		return ok
	}
}
