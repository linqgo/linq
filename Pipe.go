package linq

// Pipe returns a Query that transforms an input query by transforming its
// enumerator. If q is one-shot then the returned Query is assumed to be
// one-shot.
func Pipe[T, U any](
	q Query[T],
	enum func(next Enumerator[T]) Enumerator[U],
	options ...QueryOption[U],
) Query[U] {
	options = append([]QueryOption[U]{OneShotOption[U](q.OneShot())}, options...)
	return NewQuery(func() Enumerator[U] {
		return enum(q.enumerator())
	}, options...)
}

// PipeOneToOne returns a Pipe with a bijection from input to output elements.
func PipeOneToOne[T, U any](
	q Query[T],
	selfunc func() func(t T) U,
	options ...QueryOption[U],
) Query[U] {
	options = append([]QueryOption[U]{
		FastCountOption[U](q.fastCount()),
	}, options...)
	return Pipe(q,
		func(next Enumerator[T]) Enumerator[U] {
			sel := selfunc()
			return func() Maybe[U] {
				if t, ok := next().Get(); ok {
					return Some(sel(t))
				}
				return No[U]()
			}
		},
		options...)
}
