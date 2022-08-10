package linq

func Flatten[T any](q Query[Query[T]]) Query[T] {
	return SelectMany(q, Identity[Query[T]])
}

func FlattenSlices[T any](q Query[[]T]) Query[T] {
	return SelectMany(q, func(t []T) Query[T] { return From(t...) })
}
