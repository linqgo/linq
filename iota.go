package linq

import "golang.org/x/exp/constraints"

// Iota returns a query with all integers from 0 up.
func Iota[I constraints.Integer]() Query[I] {
	return NewQuery(func() Enumerator[I] {
		var i I
		i--
		return func() (I, bool) {
			i++
			return i, true
		}
	})
}

// Iota1 returns a query with all integers in the range [0, stop).
func Iota1[I constraints.Integer](stop I) Query[I] {
	return Iota3(0, stop, 1)
}

// Iota2 returns a query with all integers in the range [start, stop).
func Iota2[I constraints.Integer](start, stop I) Query[I] {
	return Iota3(start, stop, 1)
}

// Iota3 returns a query with every step-th integer in the range [start, stop).
func Iota3[I constraints.Integer](start, stop, step I) Query[I] {
	return NewQuery(func() Enumerator[I] {
		i := start - step
		return func() (I, bool) {
			i += step
			return i, i < stop
		}
	})
}
