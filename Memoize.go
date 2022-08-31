package linq

import "sync"

func (q Query[T]) Memoize() Query[T] {
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
					var t T
					return t, false
				}
				mux.Lock()
				defer mux.Unlock()
				if i == len(cache) {
					if t, ok := next(); ok {
						cache = append(cache, t)
						c = cache
					} else {
						done = true
						return t, false
					}
				}
			}
			return c[i], true
		}
	})
}
