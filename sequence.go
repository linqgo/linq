// Copyright 2022-2024 Marcelo Cantos
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

// SequenceCmp returns SequenceCmp(q, r, cmp).
func (q Query[T]) SequenceCmp(r Query[T], cmp CmpFn[T]) int {
	return SequenceCmp(q, r, cmp)
}

// SequenceEqualEq calls SequenceEqualEq(q, r, eq).
func (q Query[T]) SequenceEqualEq(r Query[T], eq func(a, b T) bool) bool {
	return SequenceEqualEq(q, r, eq)
}

// SequenceGreaterCmp returns SequenceGreaterCmp(q, r, cmp).
func (q Query[T]) SequenceGreaterCmp(r Query[T], cmp CmpFn[T]) bool {
	return SequenceGreaterCmp(q, r, cmp)
}

// SequenceLessCmp returns q.SequenceCmp(r) < 0.
func (q Query[T]) SequenceLessCmp(r Query[T], cmp CmpFn[T]) bool {
	return q.SequenceCmp(r, cmp) < 0
}

// SequenceEqual returns SequenceEqualEq(a, b, Equal[T]).
func SequenceEqual[T comparable](a, b Query[T]) bool {
	return SequenceEqualEq(a, b, Equal[T])
}

// SequenceEqualEq returns true if two sequences are equal, that is, a and b
// contain the same number of elements and every sequential element from a
// equals the corresponding element from b. Two elements are equal if eq(aElem,
// bElem) returns true.
func SequenceEqualEq[T any](a, b Query[T], eq func(a, b T) bool) bool {
	if lenDiff, ok := fastLenDiff(a, b); ok && lenDiff != 0 {
		return false
	}

	var end int
	for a, b := range zipSeq(a.Seq(), b.Seq(), &end) {
		if !eq(a, b) {
			return false
		}
	}
	return end == 0
}

// SequenceCmp compares elements pairwise from a and b in sequence order and
// returns when one of the following occurs:
//
//  1. if c = cmp(aElem, bElem) != 0, returns c
//  2. If either query is exhausted, returns 0 if both are exausted, < 0 if exhausted(a), > 0 if exhausted(b).
//
// This is known as lexicographical order and is equivalent to the > operator on
// strings.
func SequenceCmp[T any](a, b Query[T], cmp CmpFn[T]) int {
	var end int
	for a, b := range zipSeq(a.Seq(), b.Seq(), &end) {
		if c := cmp(a, b); c != 0 {
			return c
		}
	}
	return end
}

// SequenceGreater returns SequenceLess(b, a).
func SequenceGreater[T constraints.Ordered](a, b Query[T]) bool {
	return SequenceLess(b, a)
}

// SequenceGreaterCmp returns SequenceLessCmp(b, a, cmp).
func SequenceGreaterCmp[T any](a, b Query[T], cmp CmpFn[T]) bool {
	return SequenceLessCmp(b, a, cmp)
}

// SequenceLess compares elements pairwise from a and b in sequence order
// and returns true when one of the following occurs:
//
//  1. If two elements differ, returns aElem < bElem.
//  2. If either query is exhausted, returns !exhausted(b).
//
// This is known as lexicographical order and is analogous to the < operator on
// strings.
func SequenceLess[T constraints.Ordered](a, b Query[T]) bool {
	return SequenceLessCmp(a, b, Cmp[T])
}

// SequenceLessCmp returns SequenceCmp(a, b) < 0.
func SequenceLessCmp[T any](a, b Query[T], cmp CmpFn[T]) bool {
	return SequenceCmp(a, b, cmp) < 0
}
