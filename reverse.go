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

import "slices"

// Reverse returns a query with the elements of q in reverse.
func (q Query[T]) Reverse() Query[T] {
	return Reverse(q)
}

// Reverse returns a query with the elements of q in reverse.
func Reverse[T any](q Query[T]) Query[T] {
	var get Getter[T]
	if q.count >= 0 {
		if qget := q.getter(); qget != nil {
			last := q.count - 1
			return FromArray(ArrayFromLenGet(q.count, func(i int) T {
				if t, ok := qget(last - i); ok {
					return t
				}
				panic("element not found")
			}))
		}
	}
	return FromSeq(
		func(yield func(t T) bool) {
			s := q.ToSlice()
			slices.Reverse(s)
			for _, t := range s {
				if !yield(t) {
					return
				}
			}
		},
		OneShotOption[T](q.OneShot()),
		FastCountOption[T](q.fastCount()),
		FastGetOption(get),
	)
}
