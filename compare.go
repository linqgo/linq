package linq

import "golang.org/x/exp/constraints"

// SequenceEqual returns true if a and b contain the same number of elements and
// each sequential element from a equals the corresponding sequential element
// from b.
func SequenceEqual[T comparable](a, b Query[T]) bool {
	anext := a.Enumerator()
	bnext := b.Enumerator()
	for {
		x, aok := anext()
		y, bok := bnext()
		if aok != bok {
			return false
		}
		if !aok {
			return true
		}
		if x != y {
			return false
		}
	}
}

// SequenceLess compares elements pairwise from a and b in sequence order and
// returns true if and only if one of the following occurs:
//
//  1. Two elements differ and the element from a is less than the one from b.
//  2. Query a runs out of elements before b.
//
// This is known as lexicographical sort and is equivalent to the < operator on
// strings.
func SequenceLess[T constraints.Ordered](a, b Query[T]) bool {
	anext := a.Enumerator()
	bnext := b.Enumerator()
	for {
		x, aok := anext()
		y, bok := bnext()
		if !aok {
			return bok
		}
		if !bok {
			return false
		}
		if x != y {
			return x < y
		}
	}
}

// SequenceGreater compares elements pairwise from a and b in sequence order and
// returns true if and only if one of the following occurs:
//
//  1. Two elements differ and the element from a is greater than the one from b.
//  2. Query b runs out of elements before a.
//
// This is known as lexicographical sort and is equivalent to the > operator on
// strings.
func SequenceGreater[T constraints.Ordered](a, b Query[T]) bool {
	return SequenceLess(b, a)
}

// Shorter returns true if and only if a has fewer elements than b.
func Shorter[T any](a, b Query[T]) bool {
	anext := a.Enumerator()
	bnext := b.Enumerator()
	for {
		_, aok := anext()
		_, bok := bnext()
		if !aok {
			return bok
		}
		if !bok {
			return false
		}
	}
}

// Longer returns true if and only if a has more elements than b.
func Longer[T any](a, b Query[T]) bool {
	return Longer(b, a)
}
