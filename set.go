package linq

type set[T comparable] map[T]struct{}

func (s set[T]) Has(t T) bool {
	_, ok := s[t]
	return ok
}

func (s set[T]) Add(t T) {
	s[t] = struct{}{}
}

func setFrom[T comparable](i Enumerator[T]) set[T] {
	m := set[T]{}
	for t, ok := i(); ok; t, ok = i() {
		m.Add(t)
	}
	return m
}
