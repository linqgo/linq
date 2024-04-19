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

// FirstComp returns the element in q that precedes every other element or ok =
// false if q is empty.
func (q Query[T]) FirstComp(precedes func(a, b T) bool) Maybe[T] {
	return FirstComp(q, precedes)
}

// LastComp returns the element in q that precedes every other element or ok =
// false if q is empty.
func (q Query[T]) LastComp(precedes func(a, b T) bool) Maybe[T] {
	return LastComp(q, precedes)
}

// FirstComp returns the element in q that precedes every other element or ok =
// false if q is empty.
func FirstComp[T any](q Query[T], precedes func(a, b T) bool) Maybe[T] {
	return firstBy(q, Identity[T], precedes)
}

// LastComp returns the element in q that precedes every other element or ok =
// false if q is empty.
func LastComp[T any](q Query[T], precedes func(a, b T) bool) Maybe[T] {
	return lastBy(q, Identity[T], precedes)
}

func firstBy[T, K any](q Query[T], key func(T) K, precedes func(a, b K) bool) Maybe[T] {
	var firstValue T
	var firstKey K
	ok := false
	for i, t := range q.IRange() {
		k := key(t)
		if i == 0 || precedes(k, firstKey) {
			firstValue, firstKey = t, k
		}
		ok = true
	}
	return NewMaybe(firstValue, ok)
}

func lastBy[T, K any](q Query[T], key func(T) K, precedes func(a, b K) bool) Maybe[T] {
	return firstBy(q, key, SwapArgs(precedes))
}
