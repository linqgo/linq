package linq

// Where returns a query with elements from q for which pred returns true.
func (q Query[T]) Where(pred func(t T) bool) Query[T] {
	return Where(q, pred)
}

// WhereI returns a query with elements from q for which pred returns true. The
// pred function takes the index and value of each element.
func (q Query[T]) WhereI(pred func(i int, t T) bool) Query[T] {
	return WhereI(q, pred)
}

// Where returns a query with elements from q for which pred returns true.
func Where[T any](q Query[T], pred func(t T) bool) Query[T] {
	return WhereI(q, indexify(pred))
}

// WhereI returns a query with elements from q for which pred returns true. The
// pred function takes the index and value of each element.
func WhereI[T any](q Query[T], pred func(i int, t T) bool) Query[T] {
	return NewQuery(func() Enumerator[T] {
		next := q.Enumerator()
		i := counter(0)
		return func() (t T, ok bool) {
			for t, ok := next(); ok; t, ok = next() {
				if pred(i(), t) {
					return t, true
				}
			}
			return t, ok
		}
	})
}
