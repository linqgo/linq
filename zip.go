package linq

func Zip[A, B, R any](a Query[A], b Query[B], zip func(a A, b B) R) Query[R] {
	return NewQuery(func() Enumerator[R] {
		a := a.Enumerator()
		b := b.Enumerator()
		return func() (r R, ok bool) {
			x, ok := a()
			if !ok {
				return r, ok
			}
			y, ok := b()
			if !ok {
				return r, ok
			}
			return zip(x, y), true
		}
	})
}
