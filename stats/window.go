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

package stats

import (
	"github.com/linqgo/linq"
	"github.com/linqgo/linq/internal/num"
	"github.com/linqgo/linq/internal/ring"
)

// Window implements a sliding window. Each output element adds an input element
// to the sliding window and expels any element already in the window that
// satisfy the expired predicate. The return KV holds the new element and a
// query containing any expired elements.
func Window[T any](
	q linq.Query[T],
	expired func(older, current T) bool,
) linq.Query[linq.KV[T, linq.Query[T]]] {
	return linq.NewQuery(func() linq.Enumerator[linq.KV[T, linq.Query[T]]] {
		next := q.Enumerator()
		buf := ring.New[T](0)
		return func() linq.Maybe[linq.KV[T, linq.Query[T]]] {
			if t, ok := next().Get(); ok {
				buf.Push(t)
				var morgue []T
				for !buf.Empty() {
					h := buf.Head()
					if !expired(h, t) {
						break
					}
					buf.Pop()
					morgue = append(morgue, h)
				}
				return linq.Some(linq.NewKV(t, linq.From(morgue...)))
			}
			return linq.No[linq.KV[T, linq.Query[T]]]()
		}
	})
}

// WindowTime implements an ever-expanding window of all elements seen to date.
func WindowAll[T any](q linq.Query[T]) linq.Query[linq.KV[T, linq.Query[T]]] {
	return Window(q, func(_, _ T) bool { return false })
}

// WindowTime implements a sliding window of at most windowSize elements.
func WindowFixed[T any](q linq.Query[T], windowSize int) linq.Query[linq.KV[T, linq.Query[T]]] {
	return windowValues(WindowTime(linq.Index(q), windowSize))
}

// WindowTime implements a sliding window of elements that are younger than the
// specified expiryAge.
func WindowTime[Time num.RealNumber, T any](
	q linq.Query[linq.KV[Time, T]],
	expiryAge Time,
) linq.Query[linq.KV[linq.KV[Time, T], linq.Query[linq.KV[Time, T]]]] {
	return Window(q, func(older, current linq.KV[Time, T]) bool {
		return older.Key <= current.Key-expiryAge
	})
}

func windowValues[K, V any](
	q linq.Query[linq.KV[linq.KV[K, V], linq.Query[linq.KV[K, V]]]],
) linq.Query[linq.KV[V, linq.Query[V]]] {
	return linq.Select(q,
		func(kv linq.KV[linq.KV[K, V], linq.Query[linq.KV[K, V]]]) linq.KV[V, linq.Query[V]] {
			return linq.NewKV(kv.Key.Value, linq.SelectValues(kv.Value))
		},
	)
}
