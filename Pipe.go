package linq

func Pipe[T, U any](q Query[T], enum func(next Enumerator[T]) Enumerator[U]) Query[U] {
	return NewQuery(func() Enumerator[U] {
		return enum(q.enumerator())
	}).withOneShot(q.OneShot())
}
