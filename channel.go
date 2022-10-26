package linq

// FromChannel returns a query that reads values from c.
//
// The returned query is not replayable. Use (Query).Memoize() if you need a
// replayable query.
func FromChannel[T any](c <-chan T) Query[T] {
	return NewQuery(func() Enumerator[T] {
		return func() Maybe[T] {
			t, ok := <-c
			return NewMaybe(t, ok)
		}
	}, OneShotOption[T](true))
}
