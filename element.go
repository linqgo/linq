package linq

// ElementAt returns the element at position i or !ok if there is no element i.
func (q Query[T]) ElementAt(i int) Maybe[T] {
	return ElementAt(q, i)
}

// FastElementAt returns the element at position i or !ok if there is no element
// i or the element cannot be accessed in O(1) time.
func (q Query[T]) FastElementAt(i int) Maybe[T] {
	return FastElementAt(q, i)
}

// First returns the first element or !ok if q is empty.
func (q Query[T]) First() Maybe[T] {
	return First(q)
}

// Last returns the last element or !ok if q is empty.
func (q Query[T]) Last() Maybe[T] {
	return Last(q)
}

// ElementAt returns the element at position i or !ok if there is no element i.
func ElementAt[T any](q Query[T], i int) Maybe[T] {
	if i < 0 {
		return No[T]()
	}
	if t := FastElementAt(q, i); t.Valid() {
		return t
	}
	next := q.Enumerator()
	for ; i > 0; i-- {
		if t := next(); !t.Valid() {
			return No[T]()
		}
	}
	return next()
}

// FastElementAt returns the element at position i or !ok if there is no element
// i or element i cannot be accessed in O(1) time.
func FastElementAt[T any](q Query[T], i int) Maybe[T] {
	if i < 0 {
		return No[T]()
	}
	if q.fastCount() != 0 {
		if get := q.getter(); get != nil {
			return get(i)
		}
	}
	return No[T]()
}

// FastLast returns the last element or !ok if q is empty or the last element
// cannot be accessed in O(1) time.
func FastLast[T any](q Query[T]) Maybe[T] {
	if q.count > 0 {
		if get := q.getter(); get != nil {
			return get(q.count - 1)
		}
	}
	return No[T]()
}

// First returns the first element or !ok if q is empty.
func First[T any](q Query[T]) Maybe[T] {
	return q.Enumerator()()
}

// Last returns the last element or !ok if q is empty.
func Last[T any](q Query[T]) Maybe[T] {
	if t := FastLast(q); t.Valid() {
		return t
	}
	next := q.Enumerator()
	e, ok := next().Get()
	if !ok {
		return No[T]()
	}
	for {
		if e2, ok := next().Get(); ok {
			e = e2
		} else {
			break
		}
	}
	return Some(e)
}
