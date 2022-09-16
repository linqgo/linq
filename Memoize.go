package linq

import "sync"

// Memoize caches the elements of q. It returns a query that contains the same
// elements as q, but, in the process of enumerating it, remembers the sequence
// of values seen and ensures that every enumeration yields the same sequence.
func (q Query[T]) Memoize() Query[T] {
	return Memoize(q)
}

// Memoize caches the elements of q. It returns a query that contains the same
// elements as q, but, in the process of enumerating it, remembers the sequence
// of values seen and ensures that every enumeration yields the same sequence.
func Memoize[T any](q Query[T]) Query[T] { //nolint:revive
	var cache []T
	var mux sync.Mutex
	var next Enumerator[T]
	done := false
	return NewQuery(func() Enumerator[T] {
		mux.Lock()
		defer mux.Unlock()
		if next == nil {
			next = q.Enumerator()
		}
		i := -1
		c := cache
		return func() (T, bool) {
			// This is an unsafe check, since the access to done isn't protected
			// by a mutex, but an incorrect outcome doesn't break the logic, it
			// just misses out on the optimisation opportunity.
			if done {
				c = cache
			}
			if i++; i == len(c) {
				if done {
					i--
					var t T
					return t, false
				}
				mux.Lock()
				defer mux.Unlock()
				if i == len(cache) {
					t, ok := next()
					if !ok {
						i--
						done = true
						return t, false
					}
					cache = append(cache, t)
					c = cache
				}
			}
			return c[i], true
		}
	}, FastCountOption[T](q.fastCount()))
	// TODO: Can we switch to fast count after the input is exhausted?
}
