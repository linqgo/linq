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
	"cmp"
	"iter"
)

// Max returns the highest number in seq or ok=false if seq is empty.
func Max[R cmp.Ordered](seq iter.Seq[R]) (R, bool) {
	return Aggregate(seq, max[R])
}

// MaxBy returns the element(s) in q with the highest key.
func MaxBy[T any, R cmp.Ordered](q Query[T], key func(T) R) Query[T] {
	return bestBy(q.Seq(), key, greater[R])
}

// Min returns the lowest number in seq or ok=false if seq is empty.
func Min[R cmp.Ordered](seq iter.Seq[R]) (R, bool) {
	return Aggregate(seq, min[R])
}

// MinBy returns the element in q with the lowest key or ok = false if q is
// empty.
func MinBy[T any, K cmp.Ordered](q Query[T], key func(T) K) Query[T] {
	return bestBy(q.Seq(), key, less[K])
}

func bestBy[T any, O cmp.Ordered](seq iter.Seq[T], key func(T) O, better func(a, b O) bool) Query[T] {
	return FromSeq(func(yield func(T) bool) {
		var acc []T
		var best O
		i := 0
		for t := range seq {
			k := key(t)
			switch {
			case i == 0, better(k, best):
				best = k
				acc = acc[:0]
			case better(best, k):
				i++
				continue
			}
			acc = append(acc, t)
			i++
		}
		seqSlice(acc)(yield)
	})
}

func greater[O cmp.Ordered](a, b O) bool {
	return a > b
}

func less[O cmp.Ordered](a, b O) bool {
	return a < b
}

func max[O cmp.Ordered](a, b O) O {
	if a >= b {
		return a
	}
	return b
}

func min[O cmp.Ordered](a, b O) O {
	if a <= b {
		return a
	}
	return b
}
