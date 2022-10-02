package linq

// FirstComp returns the element in q that precedes every other element or ok =
// false if q is empty.
func (q Query[T]) FirstComp(precedes func(a, b T) bool) (_ T, ok bool) {
	return FirstComp(q, precedes)
}

// FirstCompElse returns the element in q that precedes every other element or
// ok = false if q is empty.
func (q Query[T]) FirstCompElse(precedes func(a, b T) bool, alt T) T {
	return FirstCompElse(q, precedes, alt)
}

// LastComp returns the element in q that precedes every other element or ok =
// false if q is empty.
func (q Query[T]) LastComp(precedes func(a, b T) bool) (_ T, ok bool) {
	return LastComp(q, precedes)
}

// LastCompElse returns the element in q that precedes every other element or
// ok = false if q is empty.
func (q Query[T]) LastCompElse(precedes func(a, b T) bool, alt T) T {
	return LastCompElse(q, precedes, alt)
}

// MustFirstComp returns the element in q that precedes every other element or
// panics if q is empty.
func (q Query[T]) MustFirstComp(precedes func(a, b T) bool) T {
	return MustFirstComp(q, precedes)
}

// MustLastComp returns the element in q that precedes every other element or
// panics if q is empty.
func (q Query[T]) MustLastComp(precedes func(a, b T) bool) T {
	return MustLastComp(q, precedes)
}

// FirstComp returns the element in q that precedes every other element or ok =
// false if q is empty.
func FirstComp[T any](q Query[T], precedes func(a, b T) bool) (_ T, ok bool) {
	return firstBy(q, Identity[T], precedes)
}

// FirstCompElse returns the element in q that precedes every other element or
// alt if q is empty.
func FirstCompElse[T any](q Query[T], precedes func(a, b T) bool, alt T) T {
	t, ok := firstBy(q, Identity[T], precedes)
	return valueElse(t, ok, alt)
}

// LastComp returns the element in q that precedes every other element or ok =
// false if q is empty.
func LastComp[T any](q Query[T], precedes func(a, b T) bool) (_ T, ok bool) {
	return lastBy(q, Identity[T], precedes)
}

// LastCompElse returns the element in q that precedes every other element or
// alt if q is empty.
func LastCompElse[T any](q Query[T], precedes func(a, b T) bool, alt T) T {
	t, ok := lastBy(q, Identity[T], precedes)
	return valueElse(t, ok, alt)
}

// MustFirstComp returns the element in q that precedes every other element or
// panics if q is empty.
func MustFirstComp[T any](q Query[T], precedes func(a, b T) bool) T {
	return valueOrPanicEmpty(firstBy(q, Identity[T], precedes))
}

// MustLastComp returns the element in q that precedes every other element or
// panics if q is empty.
func MustLastComp[T any](q Query[T], precedes func(a, b T) bool) T {
	return valueOrPanicEmpty(lastBy(q, Identity[T], precedes))
}

func firstBy[T, K any](q Query[T], key func(T) K, precedes func(a, b K) bool) (r T, ok bool) {
	next := q.Enumerator()
	firstValue, ok := next()
	if !ok {
		return r, ok
	}
	firstKey := key(firstValue)
	for u, ok := next(); ok; u, ok = next() {
		k := key(u)
		if precedes(k, firstKey) {
			firstValue, firstKey = u, k
		}
	}
	return firstValue, true
}

func lastBy[T, K any](q Query[T], key func(T) K, precedes func(a, b K) bool) (r T, ok bool) {
	return firstBy(q, key, SwapArgs(precedes))
}
