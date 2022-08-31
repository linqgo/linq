package linq

// FromChannel returns a query that reads values from c.
func FromChannel[T any](c <-chan T) Query[T] {
	return Query[T]{
		enumerator: func() Enumerator[T] {
			return func() (T, bool) {
				t, ok := <-c
				return t, ok
			}
		},
		extra: &queryExtra[T]{nonReplayable: true},
	}
}
