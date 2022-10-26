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
				return qget(last - i).Must()
			}))
		}
	}
	return NewQuery(
		func() Enumerator[T] {
			data := q.ToSlice()
			return func() Maybe[T] {
				var e T
				last := len(data) - 1
				if last >= 0 {
					e, data = data[last], data[:last]
				}
				return NewMaybe(e, last >= 0)
			}
		},
		OneShotOption[T](q.OneShot()),
		FastCountOption[T](q.fastCount()),
		FastGetOption(get),
	)
}
