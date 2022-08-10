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

// Sample returns a query that randomly samples each element in q with
// probability p. The returned query will deterministically sample values at the
// same intervals each time an enumerator is requested. This is not the case
// across calls to Sample.
func (q Query[T]) Sample(p float64) Query[T] {
	return Sample(q, p)
}

// SampleSeed returns a query that randomly samples each element in q with
// probability p.
//
// The seed allows for deterministic results. Multiple invokations of
// SampleSeed with the same seed will return a query that samples values
// at the same intervals.
func (q Query[T]) SampleSeed(p float64, seed int64) Query[T] {
	return SampleSeed(q, p, seed)
}

// Every returns a query that contains every nth element from q, starting at the
// start-th element.
func EveryFrom[T any](q Query[T], start, n int) Query[T] {
	return NewQuery(func() Enumerator[T] {
		next := q.Enumerator()
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
	})
}
