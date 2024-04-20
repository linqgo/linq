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

import "iter"

type Getter[T any] func(i int) (T, bool)

func (g Getter[T]) Seq() iter.Seq[T] {
	return func(yield func(t T) bool) {
		seqIota(func(i int) bool {
			t, ok := g(i)
			return ok && yield(t)
		})
	}
}

// ArrayGetter returns a Getter for an Array.
func ArrayGetter[T any](a Array[T]) Getter[T] {
	return LenGetGetter(a.Len(), a.Get)
}

// LenGetGetter returns a Getter for a len/get pair.
func LenGetGetter[T any](n int, get func(i int) T) Getter[T] {
	return func(i int) (T, bool) {
		if 0 <= i && i < n {
			return get(i), true
		}
		return no[T]()
	}
}

func FromGetter[T any](get Getter[T]) Query[T] {
	return FromSeq(
		get.Seq(),
		FastGetOption(get),
	)
}

// ToGetter returns a Getter providing access to the elements of q.
func (q Query[T]) ToGetter() Getter[T] {
	return ToGetter(q)
}

// ToGetter returns a Getter providing access to the elements of q.
func ToGetter[T any](q Query[T]) Getter[T] {
	return q.ElementAt
}
