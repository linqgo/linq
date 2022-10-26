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
	"math"

	"golang.org/x/exp/constraints"
)

// Average returns the arithmetic mean of the numbers in q or ok = false if q is
// empty.
func Average[R realNumber](q Query[R]) Maybe[R] {
	if sum, n := aggregateN(q, 0, add[R]); n > 0 {
		return Some(sum / R(n))
	}
	return No[R]()
}

// GeometricMean returns the geometric mean of the numbers in q or ok=false if q
// is empty.
func GeometricMean[R realNumber](q Query[R]) Maybe[R] {
	if product, n := aggregateN(q, 0, mul[R]); n > 0 {
		return Some(R(math.Pow(float64(product), float64(n))))
	}
	return No[R]()
}

// HarmonicMean returns the harmonic mean of the numbers in q or ok = false if q
// is empty.
func HarmonicMean[F constraints.Float](q Query[F]) Maybe[F] {
	if recipSum, n := aggregateN(q, 0, recipAdd[F]); n > 0 {
		return Some(F(n) / F(recipSum))
	}
	return No[F]()
}

// Max returns the highest number in q or ok=false if q is empty.
func Max[R realNumber](q Query[R]) Maybe[R] {
	return Aggregate(q, max[R])
}

// MaxBy returns the element in q with the highest key or ok = false if q is
// empty.
func MaxBy[T any, R constraints.Ordered](q Query[T], key func(T) R) Maybe[T] {
	return bestBy(q, key, greater[R])
}

// Min returns the highest number in q or ok=false if q is empty.
func Min[R realNumber](q Query[R]) Maybe[R] {
	return Aggregate(q, min[R])
}

// MinBy returns the element in q with the highest key or ok = false if q is
// empty.
func MinBy[T any, K constraints.Ordered](q Query[T], key func(T) K) Maybe[T] {
	return bestBy(q, key, less[K])
}

// Product returns the product of the numbers in q or 1 if q is empty.
func Product[R number](q Query[R]) R {
	return aggregate(q, 1, mul[R])
}

// Sum returns the sum of the numbers in q or 0 if q is empty.
func Sum[R number](q Query[R]) R {
	return aggregate(q, 0, add[R])
}

func add[N number](a, b N) N {
	return a + b
}

func bestBy[T any, O constraints.Ordered](q Query[T], key func(T) O, better func(a, b O) bool) Maybe[T] {
	next := q.Enumerator()
	bestValue, ok := next().Get()
	if !ok {
		return No[T]()
	}
	bestKey := key(bestValue)
	for u, ok := next().Get(); ok; u, ok = next().Get() {
		k := key(u)
		if better(k, bestKey) {
			bestValue, bestKey = u, k
		}
	}
	return Some(bestValue)
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

func mul[N number](a, b N) N {
	return a * b
}

func recipAdd[R constraints.Float](a, b R) R {
	return a + 1/b
}
