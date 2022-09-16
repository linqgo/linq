package linq

// Every returns a query that contains every nth element from q.
func (q Query[T]) Every(n int) Query[T] {
	return Every(q, n)
}

// Every returns a query that contains every nth element from q, starting at the
// start-th element.
func (q Query[T]) EveryFrom(start, n int) Query[T] {
	return EveryFrom(q, start, n)
}

// Every returns a query that contains every nth element from q.
func Every[T any](q Query[T], n int) Query[T] {
	return EveryFrom(q, 0, n)
}

// Every returns a query that contains every nth element from q, starting at the
// start-th element.
func EveryFrom[T any](q Query[T], start, n int) Query[T] {
	return Pipe(q, func(next Enumerator[T]) Enumerator[T] {
		skip := start
		return func() (T, bool) {
			for t, ok := next(); ok; t, ok = next() {
				if skip == 0 {
					skip = n - 1
					return t, true
				}
				skip--
			}
			var t T
			return t, false
		}
	}, ComputedFastCountOption[T](q.fastCount(), func(count int) int {
		return (count-start-1)/n + 1
	}))
}
