package linq

// None returns an empty query.
func None[T any]() Query[T] {
	return NewQuery(noneEnumeratorer[T])
}

func noneEnumeratorer[T any]() Enumerator[T] {
	return noneEnumerator[T]
}

func noneEnumerator[T any]() (_ T, _ bool) {
	return
}