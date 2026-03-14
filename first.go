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

import "iter"

// FirstCmp returns the element in q that cmp every other element or ok =
// false if q is empty.
func (q Query[T]) FirstCmp(cmp func(a, b T) int) (T, bool) {
	return FirstCmp(q.Seq(), cmp)
}

// LastCmp returns the element in q that cmp every other element or ok =
// false if q is empty.
func (q Query[T]) LastCmp(cmp func(a, b T) int) (T, bool) {
	return LastCmp(q.Seq(), cmp)
}

// FirstCmp returns the element in seq that cmp every other element or ok =
// false if seq is empty.
func FirstCmp[T any](seq iter.Seq[T], cmp func(a, b T) int) (T, bool) {
	return firstBy(seq, Identity[T], cmp)
}

// LastCmp returns the element in seq that cmp every other element or ok =
// false if seq is empty.
func LastCmp[T any](seq iter.Seq[T], cmp func(a, b T) int) (T, bool) {
	return lastBy(seq, Identity[T], cmp)
}

func firstBy[T, K any](seq iter.Seq[T], key func(T) K, cmp func(a, b K) int) (T, bool) {
	var firstValue T
	var firstKey K
	ok := false
	i := 0
	for t := range seq {
		k := key(t)
		if i == 0 || cmp(k, firstKey) < 0 {
			firstValue, firstKey = t, k
		}
		ok = true
		i++
	}
	return firstValue, ok
}

func lastBy[T, K any](seq iter.Seq[T], key func(T) K, cmp func(a, b K) int) (T, bool) {
	return firstBy(seq, key, SwapArgs(cmp))
}
