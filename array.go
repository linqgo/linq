package linq

type Array[T any] interface {
	Len() int
	Get(i int) T
}

func ArrayFromLenGet[T any](n int, get func(i int) T) Array[T] {
	return lenGetArray[T]{n: n, get: get}
}

func FromArray[T any](a Array[T]) Query[T] {
	return NewQuery(
		func() Enumerator[T] {
			n := a.Len()
			i := 0
			return func() Maybe[T] {
				if i == n {
					return No[T]()
				}
				t := a.Get(i)
				i++
				return Some(t)
			}
		},
		FastCountOption[T](a.Len()),
		FastGetOption(ArrayGetter(a)),
	)
}

// ToArray returns an Array interface containing the elements of q.
func (q Query[T]) ToArray() Array[T] {
	return ToArray(q)
}

// ToSlice returns a slice containing the elements of q.
func ToArray[T any](q Query[T]) Array[T] {
	return queryArray[T]{q: q}
}

type lenGetArray[T any] struct {
	n   int
	get func(i int) T
}

func (a lenGetArray[T]) Len() int {
	return a.n
}

func (a lenGetArray[T]) Get(i int) T {
	return a.get(i)
}

type queryArray[T any] struct {
	q Query[T]
}

func (a queryArray[T]) Len() int {
	return a.q.Count()
}

func (a queryArray[T]) Get(i int) T {
	return a.q.ElementAt(i).Must()
}
