package linq

// Longer returns true if and only if q has more elements than r.
func (q Query[T]) Longer(r Query[T]) bool {
	return Shorter(r, q)
}

// Shorter returns true if and only if q has fewer elements than r.
func (q Query[T]) Shorter(r Query[T]) bool {
	return Shorter(q, r)
}

// FastLonger returns true if and only if a has more elements than b and this
// can be determined in O(1) time, otherwise returns ok = false.
func (a Query[T]) FastLonger(b Query[T]) (longer, ok bool) {
	return FastLonger(a, b)
}

// FastShorter returns true if and only if a has fewer elements than b and this
// can be determined in O(1) time, otherwise returns ok = false.
func (a Query[T]) FastShorter(b Query[T]) (shorter, ok bool) {
	return FastShorter(a, b)
}

// MustFastLonger returns true if and only if a has more elements than b and
// this can be determined in O(1) time, otherwise panics.
func (a Query[T]) MustFastLonger(b Query[T]) bool {
	return MustFastLonger(a, b)
}

// MustFastShorter returns true if and only if a has fewer elements than b and
// this can be determined in O(1) time, otherwise panics.
func (a Query[T]) MustFastShorter(b Query[T]) bool {
	return MustFastShorter(a, b)
}

// FastLonger returns true if and only if a has more elements than b and this
// can be determined in O(1) time, otherwise returns ok = false.
func FastLonger[A, B any](a Query[A], b Query[B]) (longer, ok bool) {
	return FastShorter(b, a)
}

// FastShorter returns true if and only if a has fewer elements than b and this
// can be determined in O(1) time, otherwise returns ok = false.
func FastShorter[A, B any](a Query[A], b Query[B]) (shorter, ok bool) {
	diff, ok := fastLenDiff(a, b)
	return diff < 0, ok
}

// Longer returns true if and only if a has more elements than b.
func Longer[A, B any](a Query[A], b Query[B]) bool {
	return Shorter(b, a)
}

// MustFastLonger returns true if and only if a has more elements than b and
// this can be determined in O(1) time, otherwise panics.
func MustFastLonger[A, B any](a Query[A], b Query[B]) bool {
	return MustFastShorter(b, a)
}

// MustFastShorter returns true if and only if a has fewer elements than b and
// this can be determined in O(1) time, otherwise panics.
func MustFastShorter[A, B any](a Query[A], b Query[B]) bool {
	diff, ok := fastLenDiff(a, b)
	return valueOrPanicNoFastCount(diff < 0, ok)
}

// Shorter returns true if and only if a has fewer elements than b.
func Shorter[A, B any](a Query[A], b Query[B]) bool {
	if shorter, ok := FastShorter(a, b); ok {
		return shorter
	}

	var aok, bok bool
	Drain(zipEnumerator(a.Enumerator(), b.Enumerator(), &aok, &bok))
	return !aok && bok
}

func fastLenDiff[A, B any](a Query[A], b Query[B]) (int, bool) {
	alen, alenok := a.FastCount()
	blen, blenok := b.FastCount()
	return alen - blen, alenok && blenok
}
