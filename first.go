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

// FirstCmp returns the element in q that cmp every other element or ok =
// false if q is empty.
func (q Query[T]) FirstCmp(cmp func(a, b T) int) (T, bool) {
	return FirstCmp(q, cmp)
}

// LastCmp returns the element in q that cmp every other element or ok =
// false if q is empty.
func (q Query[T]) LastCmp(cmp func(a, b T) int) (T, bool) {
	return LastCmp(q, cmp)
}

// FirstCmp returns the element in q that cmp every other element or ok =
// false if q is empty.
func FirstCmp[T any](q Query[T], cmp func(a, b T) int) (T, bool) {
	return firstBy(q, Identity[T], cmp)
}

// LastCmp returns the element in q that cmp every other element or ok =
// false if q is empty.
func LastCmp[T any](q Query[T], cmp func(a, b T) int) (T, bool) {
	return lastBy(q, Identity[T], cmp)
}

func firstBy[T, K any](q Query[T], key func(T) K, cmp func(a, b K) int) (T, bool) {
	var firstValue T
	var firstKey K
	ok := false
	for i, t := range q.ISeq() {
		k := key(t)
		if i == 0 || cmp(k, firstKey) < 0 {
			firstValue, firstKey = t, k
		}
		ok = true
	}
	return firstValue, ok
}

func lastBy[T, K any](q Query[T], key func(T) K, cmp func(a, b K) int) (T, bool) {
	return firstBy(q, key, SwapArgs(cmp))
}
