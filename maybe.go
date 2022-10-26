package linq

type Maybe[T any] struct {
	t     T
	valid bool
}

func NewMaybe[T any](t T, valid bool) Maybe[T] {
	return Maybe[T]{t: t, valid: valid}
}

func No[T any]() Maybe[T] {
	return Maybe[T]{}
}

func Some[T any](t T) Maybe[T] {
	return Maybe[T]{t: t, valid: true}
}

func (m Maybe[T]) Else(alt T) T {
	if m.valid {
		return m.t
	}
	return alt
}

func (m Maybe[T]) Get() (T, bool) {
	return m.t, m.valid
}

func (m Maybe[T]) Must() T {
	if m.valid {
		return m.t
	}
	panic(NoValueError)
}

func (m Maybe[T]) Valid() bool {
	return m.valid
}

func ElseNaN[R realNumber](r Maybe[R]) R {
	return r.Else(R(nan))
}
