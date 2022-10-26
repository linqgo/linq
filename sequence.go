// Copyright 2022 Marcelo Cantos
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package linq

import (
	"golang.org/x/exp/constraints"
)

// SequenceEqualEq returns true if q and r contain the same number of elements
// and each sequential element from a equals the corresponding sequential
// element from b. The eq function is called to determine equality.
func (q Query[T]) SequenceEqualEq(r Query[T], eq func(a, b T) bool) bool {
	return SequenceEqualEq(q, r, eq)
}

// SequenceGreaterComp compares elements pairwise from q and r in sequence order
// and returns true if and only if one of the following occurs:
//
//  1. less(rElem, qElem) returns true. (Note the reversed parameters.)
//  2. Query r runs out of elements before q.
//
// This is known as lexicographical sort and is equivalent to the < operator on
// strings.
func (q Query[T]) SequenceGreaterComp(r Query[T], less func(a, b T) bool) bool {
	return SequenceGreaterComp(q, r, less)
}

// SequenceLessComp compares elements pairwise from q and r in sequence order
// and returns true if and only if one of the following occurs:
//
//  1. less(qElem, rElem) returns true.
//  2. Query q runs out of elements before r.
//
// This is known as lexicographical sort and is equivalent to the < operator on
// strings.
func (q Query[T]) SequenceLessComp(r Query[T], less func(a, b T) bool) bool {
	return SequenceLessComp(q, r, less)
}

// SequenceEqual returns true if a and b contain the same number of elements and
// each sequential element from a equals the corresponding sequential element
// from b.
func SequenceEqual[T comparable](a, b Query[T]) bool {
	return SequenceEqualEq(a, b, Equal[T])
}

// SequenceEqualEq returns true if a and b contain the same number of elements
// and each sequential element from a equals the corresponding sequential
// element from b. The eq function is called to determine equality.
func SequenceEqualEq[T any](a, b Query[T], eq func(a, b T) bool) bool {
	if lenDiff, ok := fastLenDiff(a, b).Get(); ok && lenDiff != 0 {
		return false
	}

	var aok, bok bool
	next := zipEnumerator(a.Enumerator(), b.Enumerator(), &aok, &bok)
	for ab, ok := next().Get(); ok; ab, ok = next().Get() {
		x, y := ab.KV()
		if !eq(x, y) {
			return false
		}
	}
	return aok == bok
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

// SequenceGreaterComp compares elements pairwise from a and b in sequence order
// and returns true if and only if one of the following occurs:
//
//  1. less(bElem, aElem) returns true. (Note the order of parameters.)
//  2. Query b runs out of elements before a.
//
// This is known as lexicographical sort and is equivalent to the < operator on
// strings.
func SequenceGreaterComp[T any](a, b Query[T], less func(a, b T) bool) bool {
	return SequenceLessComp(b, a, less)
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
	return SequenceLessComp(a, b, Less[T])
}

// SequenceLessComp compares elements pairwise from a and b in sequence order
// and returns true if and only if one of the following occurs:
//
//  1. less(aElem, bElem) returns true.
//  2. Query a runs out of elements before b.
//
// This is known as lexicographical sort and is equivalent to the < operator on
// strings.
func SequenceLessComp[T any](a, b Query[T], less func(a, b T) bool) bool {
	var aok, bok bool
	next := zipEnumerator(a.Enumerator(), b.Enumerator(), &aok, &bok)
	for ab, ok := next().Get(); ok; ab, ok = next().Get() {
		x, y := ab.KV()
		if less(x, y) {
			return true
		}
		if less(y, x) {
			return false
		}
	}
	return !aok && bok
}
