package linq

func valueEnumerator[T any](t T) Enumerator[T] {
	done := false
	return func() (T, bool) {
		valid := !done
		done = true
		return t, valid
	}
}

func sliceEnumerator[T any](s []T) Enumerator[T] {
	i := -1
	return func() (T, bool) {
		if i++; i == len(s) {
			i--
			var t T
			return t, false
		}
		return s[i], true
	}
}

// From returns a query containing the specified parameters.
func From[T any](t ...T) Query[T] {
	if len(t) == 0 {
		return None[T]()
	}
	if len(t) == 1 {
		v := t[0]
		return NewQuery(func() Enumerator[T] {
			return valueEnumerator(v)
		})
	}
	return NewQuery(func() Enumerator[T] {
		return sliceEnumerator(t)
	})
}
