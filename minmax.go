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

import "golang.org/x/exp/constraints"

// Max returns the highest number in q or ok=false if q is empty.
func Max[R constraints.Ordered](q Query[R]) Maybe[R] {
	return Aggregate(q, max[R])
}

// MaxBy returns the element(s) in q with the highest key.
func MaxBy[T any, R constraints.Ordered](q Query[T], key func(T) R) Query[T] {
	return bestBy(q, key, greater[R])
}

// Min returns the highest number in q or ok=false if q is empty.
func Min[R constraints.Ordered](q Query[R]) Maybe[R] {
	return Aggregate(q, min[R])
}

// MinBy returns the element in q with the highest key or ok = false if q is
// empty.
func MinBy[T any, K constraints.Ordered](q Query[T], key func(T) K) Query[T] {
	return bestBy(q, key, less[K])
}

func bestBy[T any, O constraints.Ordered](q Query[T], key func(T) O, better func(a, b O) bool) Query[T] {
	return FromSeq(func(yield func(T) bool) {
		var acc []T
		var best O
		for i, t := range q.IRange() {
			k := key(t)
			switch {
			case i == 0, better(k, best):
				best = k
				acc = acc[:0]
			case better(best, k):
				continue
			}
			acc = append(acc, t)
		}
		for _, t := range acc {
			if !yield(t) {
				return
			}
		}
	})
}

func greater[O constraints.Ordered](a, b O) bool {
	return a > b
}

func less[O constraints.Ordered](a, b O) bool {
	return a < b
}

func max[O constraints.Ordered](a, b O) O {
	if a >= b {
		return a
	}
	return b
}

func min[O constraints.Ordered](a, b O) O {
	if a <= b {
		return a
	}
	return b
}
