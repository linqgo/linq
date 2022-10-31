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
// elements are entering and exiting the window as it slides along.
//
// Maximality is an important concept in the discussion below. A maximal window
// is one for which expired(tail, head) is false, but which would become true if
// the window were extended at either end. Put another way, a window is maximal
// if it has no expired elements but becomes invalid if grown.
//
// Initial output is determined by the slideIn parameter. As the input is read,
// the sliding window grows by one element for each input element. Thereafter,
// every output window is maximal. If the slideIn parameter is true, the initial
// non-maximal windows will be output. If false, they will be skipped.
//
// The maximality criterion implies that all output windows bar the initial ones
// satisfy len(in) > 0 && len(out) > 0. This is because adding elements to a
// maximal window without removing any would necessarily produce an invalid
// window, whereas removing elements from a maximal window without adding any
// would necessarily produce a non-maximal window.
//
// For some windows, len(in) > 1 && len(out) > 1, which means that the window
// has slid forward more than one step. How this might happen can be seen when
// considering how a maximal window is advanced. The first step is to remove
// stale elements from the tail until it is possible to add new elements to the
// head. This could require the removal of more than element. Once enough
// elements are removed, the window must be grown until it is maximal again, and
// this could in turn require more than one addition. The example below
// illustrates one such scenario. The second delta represents the window sliding
// from {1, 2, 3} to {3, 5, 5}. This skips {2, 3, 5}, which contains expired
// element 2 and is thus invalid. In other scenarios, skipped windows will be
// valid but non-maximal.
//
// Example:
//
//	Slide(
//	    linq.From(1, 2, 3, 5, 5, 6, 7, 8),
//	    func(tail, head int) bool { return tail < head - 2 },
//	)
//	// Window:  {1,2,3}     {3,5,5}     {5,5,6,7}    {6,7,8}
//	// Delta:  ⁺{1,2,3}  ⁻{1,2}⁺{5,5}  ⁻{3}⁺{6,7}  ⁻{5,5}⁺{8}
func Slide[T any](
	q Query[T],
	slideIn bool,
	expired func(older, current T) bool,
) Query[Delta[T]] {
	return NewQuery(func() Enumerator[Delta[T]] {
		next := q.Enumerator()
		t, ok := next().Get()
		window := ring.New[T]()
		return func() Maybe[Delta[T]] {
			if !ok {
				return No[Delta[T]]()
			}
			// Empty all expired elements
			out := make([]T, 0, 1)
			window.Push(t)
			for !window.Empty() && expired(window.Head(), t) {
				out = append(out, window.Pop())
				slideIn = false
			}

			// Add elements till window is maximal unless slideIn
			in := []T{t}
			t, ok = next().Get()
			if !slideIn {
				head := window.Head()
				for ; ok; t, ok = next().Get() {
					if expired(head, t) {
						break
					}
					in = append(in, t)
					window.Push(t)
				}
			}
			return Some(newDelta(From(out...), From(in...)))
		}
	})
}

// SlideAll implements an ever-expanding window of all elements seen to date.
func SlideAll[T any](q Query[T]) Query[Delta[T]] {
	return NewQuery(func() Enumerator[Delta[T]] {
		next := q.Enumerator()
		return func() Maybe[Delta[T]] {
			if t, ok := next().Get(); ok {
				return Some(newDelta(None[T](), From(t)))
			}
			return No[Delta[T]]()
		}
	})
}

// SlideFixed implements a sliding window of at most windowSize elements.
func SlideFixed[T any](q Query[T], windowSize int, slideIn bool) Query[Delta[T]] {
	return windowValues(SlideTime(Index(q), windowSize, slideIn))
}

// SlideTime implements a sliding window of elements that are younger than the
// specified expiryAge.
func SlideTime[Time num.RealNumber, T any](
	q Query[KV[Time, T]],
	expiryAge Time,
	slideIn bool,
) Query[Delta[KV[Time, T]]] {
	return Slide(q, slideIn, func(older, current KV[Time, T]) bool {
		return older.Key <= current.Key-expiryAge
	})
}

func windowValues[K, V any](
	q Query[Delta[KV[K, V]]],
) Query[Delta[V]] {
	return Select(q,
		func(d Delta[KV[K, V]]) Delta[V] {
			return newDelta(
				Select(d.Outs, Value[KV[K, V]]),
				Select(d.Ins, Value[KV[K, V]]),
			)
		},
	)
}

type Delta[T any] struct {
	Outs, Ins Query[T]
}

func newDelta[T any](out, in Query[T]) Delta[T] {
	return Delta[T]{Outs: out, Ins: in}
}
