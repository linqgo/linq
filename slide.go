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
	"iter"

	"github.com/linqgo/linq/v2/internal/num"
	"github.com/linqgo/linq/v2/internal/ring"
)

type Delta[T any] struct {
	Outs Query[T]
	In   T
}

func newDelta[T any](outs Query[T], in T) Delta[T] {
	return Delta[T]{Outs: outs, In: in}
}

// Slide implements a sliding window over a seq. Windows are represented as
// Deltas indicating which element is entering and which, if any, elements are
// exiting the window.
func Slide[T any](seq iter.Seq[T], expired func(older, current T) bool) iter.Seq[Delta[T]] {
	return func(yield func(Delta[T]) bool) {
		window := ring.New[T]()
		for t := range seq {
			out := make([]T, 0, 1)
			window.Push(t)
			for !window.Empty() && expired(window.Head(), t) {
				out = append(out, window.Pop())
			}
			if !yield(newDelta(From(out...), t)) {
				return
			}
		}
	}
}

// SlideQuery implements a sliding window over a Query.
func SlideQuery[T any](q Query[T], expired func(older, current T) bool) Query[Delta[T]] {
	return Pipe(q, Slide(q.Seq(), expired))
}

// SlideAll implements an ever-expanding window of all elements seen to date.
func SlideAll[T any](seq iter.Seq[T]) iter.Seq[Delta[T]] {
	return Slide(seq, func(T, T) bool { return false })
}

// SlideAllQuery implements an ever-expanding window over a Query.
func SlideAllQuery[T any](q Query[T]) Query[Delta[T]] {
	return SlideQuery(q, func(T, T) bool { return false })
}

// SlideFixed implements a sliding window of at most windowSize elements.
func SlideFixed[T any](seq iter.Seq[T], windowSize int) iter.Seq[Delta[T]] {
	return windowValues(SlideTime(indexSeq(seq), windowSize))
}

// SlideFixedQuery implements a sliding window of at most windowSize elements
// over a Query.
func SlideFixedQuery[T any](q Query[T], windowSize int) Query[Delta[T]] {
	return Pipe(q, SlideFixed(q.Seq(), windowSize))
}

// SlideTime implements a sliding window of elements younger than expiryAge.
func SlideTime[Time num.RealNumber, T any](
	seq iter.Seq[KV[Time, T]],
	expiryAge Time,
) iter.Seq[Delta[KV[Time, T]]] {
	return Slide(seq, func(older, current KV[Time, T]) bool {
		return current.Key-older.Key >= expiryAge
	})
}

// SlideTimeQuery implements a sliding window of elements younger than expiryAge
// over a Query.
func SlideTimeQuery[Time num.RealNumber, T any](
	q Query[KV[Time, T]],
	expiryAge Time,
) Query[Delta[KV[Time, T]]] {
	return SlideQuery(q, func(older, current KV[Time, T]) bool {
		return current.Key-older.Key >= expiryAge
	})
}

func windowValues[K, V any](seq iter.Seq[Delta[KV[K, V]]]) iter.Seq[Delta[V]] {
	return Select(seq,
		func(d Delta[KV[K, V]]) Delta[V] {
			return newDelta(
				FromSeq(Select(d.Outs.Seq(), Value[KV[K, V]])),
				d.In.Value,
			)
		},
	)
}

// indexSeq adds integer indices to elements of a seq.
func indexSeq[T any](seq iter.Seq[T]) iter.Seq[KV[int, T]] {
	return func(yield func(KV[int, T]) bool) {
		i := 0
		for t := range seq {
			if !yield(NewKV(i, t)) {
				return
			}
			i++
		}
	}
}
