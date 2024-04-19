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

// Aggregate applies an aggregator function to the elements of q and returns the
// aggregated result or !ok if q is empty.
func (q Query[T]) Aggregate(agg func(a, b T) T) Maybe[T] {
	return Aggregate(q, agg)
}

// AggregateSeed applies an aggregator function to the elements of q, using
// seed as the initial value, and returns the aggregated result.
//
// Use the global AggregateSeed function if the seed and result are not of
// type T (e.g., concatenate a Query[int] into a string).
func (q Query[T]) AggregateSeed(seed T, agg func(a, b T) T) T {
	return AggregateSeed(q, seed, agg)
}

// Aggregate applies an aggregator function to the elements of q and returns the
// aggregated result or !ok if q is empty.
func Aggregate[T any](q Query[T], agg func(a, b T) T) Maybe[T] {
	next, stop := iter.Pull(q.Range())
	defer stop()
	if seed, ok := next(); ok {
		agg, _ := aggregateNEnum(func(yield func(T) bool) {
			for t := range nextToSeq(next) {
				if !yield(t) {
					return
				}
			}
		}, seed, agg)
		return Some(agg)
	}
	return No[T]()
}

// AggregateSeed applies an aggregator function to the elements of q, using seed
// as the initial value, and returns the aggregated result.
func AggregateSeed[T, A any](q Query[T], seed A, agg func(a A, t T) A) A {
	return aggregate(q, seed, agg)
}

func aggregate[T, A any](q Query[T], acc A, agg func(a A, t T) A) A {
	t, _ := aggregateN(q, acc, agg)
	return t
}

func aggregateN[T, A any](q Query[T], acc A, agg func(a A, t T) A) (A, int) {
	return aggregateNEnum(q.Range(), acc, agg)
}

func aggregateNEnum[T, A any](seq iter.Seq[T], acc A, agg func(a A, t T) A) (A, int) {
	n := 0
	for e := range seq {
		acc = agg(acc, e)
		n++
	}
	return acc, n
}
