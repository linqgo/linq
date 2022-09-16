package linq

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
