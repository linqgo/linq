package linq

// ElementAt returns the element at position i or !ok if there is no element i.
func (q Query[T]) ElementAt(i int) (t T, ok bool) {
	return ElementAt(q, i)
}

// ElementAtElse returns the element at position i or alt if there is no element
// i.
func (q Query[T]) ElementAtElse(i int, alt T) T {
	return ElementAtElse(q, i, alt)
}

// First returns the first element or !ok if q is empty.
func (q Query[T]) First() (_ T, ok bool) {
	return First(q)
}

// FirstElse returns the first element or alt if q is empty.
func (q Query[T]) FirstElse(alt T) T {
	return FirstElse(q, alt)
}

// Last returns the last element or !ok if q is empty.
func (q Query[T]) Last() (_ T, ok bool) {
	return Last(q)
}

// LastElse returns the last element or alt if q is empty.
func (q Query[T]) LastElse(alt T) T {
	return LastElse(q, alt)
}

// MustElementAt returns the element at position i or panics if there is no
// element i.
func (q Query[T]) MustElementAt(i int) T {
	return MustElementAt(q, i)
}

// MustFirst returns the first element or panics if q is empty.
func (q Query[T]) MustFirst() T {
	return MustFirst(q)
}

// MustLast returns the last element or panics if q is empty.
func (q Query[T]) MustLast() T {
	return MustLast(q)
}

// ElementAt returns the element at position i or !ok if there is no element i.
func ElementAt[T any](q Query[T], i int) (t T, ok bool) {
	next := q.Enumerator()
	for ; i > 0; i-- {
		if _, ok = next(); !ok {
			return t, ok
		}
	}
	return next()
}

// ElementAtElse returns the element at position i or alt if there is no element
// i.
func ElementAtElse[T any](q Query[T], i int, alt T) T {
	e, ok := q.ElementAt(i)
	return valueElse(e, ok, alt)
}

// First returns the first element or !ok if q is empty.
func First[T any](q Query[T]) (e T, ok bool) {
	return q.Enumerator()()
}

// FirstElse returns the first element or alt if q is empty.
func FirstElse[T any](q Query[T], alt T) T {
	e, ok := First(q)
	return valueElse(e, ok, alt)
}

// Last returns the last element or !ok if q is empty.
func Last[T any](q Query[T]) (t T, ok bool) {
	next := q.Enumerator()
	e, ok := next()
	if !ok {
		return t, ok
	}
	for {
		if e2, ok := next(); ok {
			e = e2
		} else {
			break
		}
	}
	return e, true
}

// LastElse returns the last element or alt if q is empty.
func LastElse[T any](q Query[T], alt T) T {
	e, ok := q.Last()
	return valueElse(e, ok, alt)
}

// MustElementAt returns the element at position i or panics if there is no
// element i.
func MustElementAt[T any](q Query[T], i int) T {
	e, ok := ElementAt(q, i)
	return valueOrPanicf(e, ok, "no element at %d", i)
}

// MustFirst returns the first element or panics if q is empty.
func MustFirst[T any](q Query[T]) T {
	e, ok := First(q)
	return valueOrPanic(e, ok, emptySourceError)
}

// MustLast returns the last element or panics if q is empty.
func MustLast[T any](q Query[T]) T {
	e, ok := Last(q)
	return valueOrPanic(e, ok, emptySourceError)
}
