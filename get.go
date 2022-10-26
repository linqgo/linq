package linq

type Getter[T any] func(i int) Maybe[T]

// ArrayGetter returns a Getter for an Array.
func ArrayGetter[T any](a Array[T]) Getter[T] {
	return LenGetGetter(a.Len(), a.Get)
}

// LenGetGetter returns a Getter for a len/get pair.
func LenGetGetter[T any](n int, get func(i int) T) Getter[T] {
	return func(i int) Maybe[T] {
		if 0 <= i && i < n {
			return Some(get(i))
		}
		return No[T]()
	}
}

func FromGetter[T any](get Getter[T]) Query[T] {
	return NewQuery(
		func() Enumerator[T] {
			i := 0
			return func() Maybe[T] {
				t := get(i)
				if t.Valid() {
					i++
				}
				return t
			}
		},
		FastGetOption(get),
	)
}

// ToGetter returns a Getter providing access to the elements of q.
func (q Query[T]) ToGetter() Getter[T] {
	return ToGetter(q)
}

// ToGetter returns a Getter providing access to the elements of q.
func ToGetter[T any](q Query[T]) Getter[T] {
	return q.ElementAt
}
