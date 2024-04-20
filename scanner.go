package linq

import "iter"

func (q Query[T]) Scanner() (pull func(p *T) bool, stop func()) {
	return Scanner(q)
}

// Scanner returns a function that scans entries from q, setting the passed
// pointer parameter to each element and returning true until it runs out of
// elements. Then it returns false.
func Scanner[T any](q Query[T]) (scan func(p *T) bool, stop func()) {
	next, stop := iter.Pull(q.Range())
	return func(p *T) bool {
		if t, ok := next(); ok {
			*p = t
			return true
		}
		return false
	}, stop
}
