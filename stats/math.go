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

package stats

import (
	"math"

	"github.com/linqgo/linq/v2"
	"github.com/linqgo/linq/v2/internal/num"
)

// Mean returns the arithmetic mean of the numbers in q or ok = false if q is
// empty.
//
// This function is equivalent to "..".Average, which is retained for parity
// with .Net's Enumerable class.
func Mean[R num.RealNumber](q linq.Query[R]) (R, bool) {
	return linq.Average(q.Seq())
}

// GeometricMean returns the geometric mean of the numbers in q or ok=false if q
// is empty.
func GeometricMean[R num.RealNumber](q linq.Query[R]) (R, bool) {
	if product, n := aggregateN(q, 0, mul[R]); n > 0 {
		return R(math.Pow(float64(product), float64(n))), true
	}
	var zero R
	return zero, false
}

// HarmonicMean returns the harmonic mean of the numbers in q or ok = false if q
// is empty.
func HarmonicMean[F num.Float](q linq.Query[F]) (F, bool) {
	if recipSum, n := aggregateN(q, 0, recipAdd[F]); n > 0 {
		return F(n) / F(recipSum), true
	}
	var zero F
	return zero, false
}

// Product returns the product of the numbers in q or 1 if q is empty.
func Product[R num.Number](q linq.Query[R]) R { return aggregate(q, 1, mul[R]) }

// Sum returns the sum of the num.Numbers in q or 0 if q is empty.
//
// This function is equivalent to "..".Sum, which is retained for parity with
// .Net's Enumerable class.
func Sum[R num.Number](q linq.Query[R]) R { return linq.Sum(q.Seq()) }

func add[N num.Number](a, b N) N             { return a + b }
func sub[N num.Number](a, b N) N             { return a - b }
func mul[N num.Number](a, b N) N             { return a * b }
func div[N num.Number](a, b N) N             { return a / b }
func recipAdd[R num.Float](a, b R) R { return a + 1/b }
func recipSub[R num.Float](a, b R) R { return a - 1/b }
