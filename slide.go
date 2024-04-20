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
	"github.com/linqgo/linq/internal/num"
	"github.com/linqgo/linq/internal/ring"
)

// Slide implements a sliding window. All output elements represent windows over
// the input elements. Windows are represented as Deltas indicating which
// element is entering and which, if any, elements are exiting the window as it
// slides along.
func Slide[T any](q Query[T], expired func(older, current T) bool) Query[Delta[T]] {
	return FromSeq(func(yield func(Delta[T]) bool) {
		window := ring.New[T]()
		for t := range q.Range() {
			// Empty all expired elements
			out := make([]T, 0, 1)
			window.Push(t)
			for !window.Empty() && expired(window.Head(), t) {
				out = append(out, window.Pop())
			}
			if !yield(newDelta(From(out...), t)) {
				return
			}
		}
	})
}

// SlideAll implements an ever-expanding window of all elements seen to date.
func SlideAll[T any](q Query[T]) Query[Delta[T]] {
	return Slide(q, func(T, T) bool { return false })
}

// SlideFixed implements a sliding window of at most windowSize elements.
func SlideFixed[T any](q Query[T], windowSize int) Query[Delta[T]] {
	return windowValues(SlideTime(Index(q), windowSize))
}

// SlideTime implements a sliding window of elements that are younger than the
// specified expiryAge.
func SlideTime[Time num.RealNumber, T any](
	q Query[KV[Time, T]],
	expiryAge Time,
) Query[Delta[KV[Time, T]]] {
	return Slide(q, func(older, current KV[Time, T]) bool {
		return current.Key-older.Key >= expiryAge
	})
}

func windowValues[K, V any](q Query[Delta[KV[K, V]]]) Query[Delta[V]] {
	return Select(q,
		func(d Delta[KV[K, V]]) Delta[V] {
			return newDelta(
				Select(d.Outs, Value[KV[K, V]]),
				d.In.Value,
			)
		},
	)
}

type Delta[T any] struct {
	Outs Query[T]
	In   T
}

func newDelta[T any](outs Query[T], in T) Delta[T] {
	return Delta[T]{Outs: outs, In: in}
}
