package linq

func (q Query[T]) Single() Maybe[T] {
	return Single(q)
}

func Single[T any](q Query[T]) Maybe[T] {
	next := q.Enumerator()
	if t, ok := next().Get(); ok {
		if _, tooMany := next().Get(); !tooMany {
			return Some(t)
		}
	}
	return No[T]()
}
