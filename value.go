package linq

func valueEnumerator[T any](t T) Enumerator[T] {
	ok := true
	return func() Maybe[T] {
		valid := ok
		ok = false
		return NewMaybe(t, valid)
	}
}

func sliceEnumerator[T any](s []T) Enumerator[T] {
	i := 0
	return func() Maybe[T] {
		if i == len(s) {
			return No[T]()
		}
		t := s[i]
		i++
		return Some(t)
	}
}

// From returns a query containing the specified parameters.
func From[T any](t ...T) Query[T] {
	if len(t) == 0 {
		return None[T]()
	}

	return NewQuery(
		func() Enumerator[T] {
			return sliceEnumerator(t)
		},
		FastCountOption[T](len(t)),
		FastGetOption(LenGetGetter(len(t), func(i int) T { return t[i] })),
	)
}
