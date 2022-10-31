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
	"math"

	"golang.org/x/exp/constraints"

	"github.com/linqgo/linq"
	"github.com/linqgo/linq/internal/num"
)

// Mean returns the arithmetic mean of the numbers in q or ok = false if q is
// empty.
//
// This function is equivalent to "github.com/linqgo/linq".Average, which is
// retained for parity with .Net's Enumerable class.
func Mean[R num.RealNumber](q linq.Query[R]) linq.Maybe[R] {
	return linq.Average(q)
}

// GeometricMean returns the geometric mean of the numbers in q or ok=false if q
// is empty.
func GeometricMean[R num.RealNumber](q linq.Query[R]) linq.Maybe[R] {
	if product, n := aggregateN(q, 0, mul[R]); n > 0 {
		return linq.Some(R(math.Pow(float64(product), float64(n))))
	}
	return linq.No[R]()
}

// HarmonicMean returns the harmonic mean of the numbers in q or ok = false if q
// is empty.
func HarmonicMean[F constraints.Float](q linq.Query[F]) linq.Maybe[F] {
	if recipSum, n := aggregateN(q, 0, recipAdd[F]); n > 0 {
		return linq.Some(F(n) / F(recipSum))
	}
	return linq.No[F]()
}

// Product returns the product of the numbers in q or 1 if q is empty.
func Product[R num.Number](q linq.Query[R]) R {
	return aggregate(q, 1, mul[R])
}

// Sum returns the sum of the num.Numbers in q or 0 if q is empty.
//
// This function is equivalent to "github.com/linqgo/linq".Sum, which is
// retained for parity with .Net's Enumerable class.
func Sum[R num.Number](q linq.Query[R]) R {
	return linq.Sum(q)
}

func add[N num.Number](a, b N) N {
	return a + b
}

func mul[N num.Number](a, b N) N {
	return a * b
}

func recipAdd[R constraints.Float](a, b R) R {
	return a + 1/b
}
