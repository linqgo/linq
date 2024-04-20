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

type Array[T any] interface {
	Len() int
	Get(i int) T
}

func ArrayFromLenGet[T any](n int, get func(i int) T) Array[T] {
	return lenGetArray[T]{n: n, get: get}
}

func FromArray[T any](a Array[T]) Query[T] {
	return FromSeq(
		func(yield func(T) bool) {
			for i, n := 0, a.Len(); i < n; i++ {
				if !yield(a.Get(i)) {
					return
				}
			}
		},
		FastCountOption[T](a.Len()),
		FastGetOption(ArrayGetter(a)),
	)
}

// ToArray returns an Array interface containing the elements of q.
func (q Query[T]) ToArray() Array[T] {
	return ToArray(q)
}

// ToSlice returns a slice containing the elements of q.
func ToArray[T any](q Query[T]) Array[T] {
	return queryArray[T]{q: q}
}

type lenGetArray[T any] struct {
	n   int
	get func(i int) T
}

func (a lenGetArray[T]) Len() int {
	return a.n
}

func (a lenGetArray[T]) Get(i int) T {
	return a.get(i)
}

type queryArray[T any] struct {
	q Query[T]
}

func (a queryArray[T]) Len() int {
	return a.q.Count()
}

func (a queryArray[T]) Get(i int) T {
	if t, ok := a.q.ElementAt(i); ok {
		return t
	}
	panic("element not found")
}
