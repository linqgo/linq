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

// AccMean accumulates the arithmetic mean of the input values within a sliding
// window. Use the Slide... functions to construct a suitable input window.
func AccMean[R num.RealNumber](q linq.Query[linq.Delta[R]]) linq.Query[R] {
	return linq.PipeOneToOne(q, func() func(r linq.Delta[R]) R {
		sum := R(0)
		n := 0
		return func(r linq.Delta[R]) R {
			outs, nOuts := aggregateN(r.Outs, 0, add[R])
			ins, nIns := aggregateN(r.Ins, 0, add[R])
			sum += ins - outs
			n += nIns - nOuts
			return sum / R(n)
		}
	})
}

// // AccMean accumulates the arithmetic mean of the input values within
// // a sliding window. Use the Slide... functions to construct a suitable input
// // window.
// func AccMeanWt[W, R num.RealNumber](
// 	q linq.Query[linq.KV[linq.KV[W, R], linq.Query[linq.KV[W, R]]]],
// ) linq.Query[R] {
// 	return linq.PipeOneToOne(q,
// 		func() func(r linq.KV[linq.KV[W, R], linq.Query[linq.KV[W, R]]]) R {
// 			var totalWt W
// 			var sum R
// 			return func(r linq.KV[linq.KV[W, R], linq.Query[linq.KV[W, R]]]) R {
// 				birthWt, birth := r.Key.KV()
// 				agg, _ := aggregateN(r.Value, linq.NewKV[W, R](0, 0), addWt[W, R])
// 				deathsWt, deaths := agg.KV()
// 				totalWt += birthWt - deathsWt
// 				sum += birth - deaths
// 				return sum / R(totalWt)
// 			}
// 		},
// 	)
// }

// AccGeometricMean accumulates the geometric mean of the input values within a
// sliding window. Use the Slide... functions to construct a suitable input
// window.
func AccGeometricMean[R num.RealNumber](q linq.Query[linq.Delta[R]]) linq.Query[R] {
	return linq.PipeOneToOne(q, func() func(r linq.Delta[R]) R {
		product := R(1)
		n := 0
		return func(r linq.Delta[R]) R {
			outs, nOuts := aggregateN(r.Outs, 1, mul[R])
			ins, nIns := aggregateN(r.Ins, 1, mul[R])
			product *= ins / outs
			n += nIns - nOuts
			return R(math.Pow(float64(product), 1/float64(n)))
		}
	})
}

// AccHarmonicMean accumulates the harmonic mean of the input values within a
// sliding window. Use the Slide... functions to construct a suitable input
// window.
func AccHarmonicMean[F constraints.Float](q linq.Query[linq.Delta[F]]) linq.Query[F] {
	return linq.PipeOneToOne(q, func() func(r linq.Delta[F]) F {
		recipSum := F(0)
		n := 0
		return func(r linq.Delta[F]) F {
			outs, nOuts := aggregateN(r.Outs, 0, recipAdd[F])
			ins, nIns := aggregateN(r.Ins, 0, recipAdd[F])
			recipSum += ins - outs
			n += nIns - nOuts
			return F(n) / F(recipSum)
		}
	})
}

// AccProduct accumulates the product of the input values within a sliding
// window. Use the Slide... functions to construct a suitable input window.
func AccProduct[R num.RealNumber](q linq.Query[linq.Delta[R]]) linq.Query[R] {
	return linq.PipeOneToOne(q, func() func(r linq.Delta[R]) R {
		product := R(1)
		return func(r linq.Delta[R]) R {
			product *= Product(r.Ins) / Product(r.Outs)
			return product
		}
	})
}

// AccSum accumulates the sum of the input values within a sliding window. Use
// the Slide... functions to construct a suitable input window.
func AccSum[R num.RealNumber](q linq.Query[linq.Delta[R]]) linq.Query[R] {
	return linq.PipeOneToOne(q, func() func(r linq.Delta[R]) R {
		sum := R(0)
		return func(r linq.Delta[R]) R {
			sum += Sum(r.Ins) - Sum(r.Outs)
			return sum
		}
	})
}

// func addWt[W, R num.RealNumber](a, b linq.KV[W, R]) linq.KV[W, R] {
// 	wa, ta := a.KV()
// 	wb, tb := b.KV()
// 	return linq.NewKV(wa+wb, R(wa)*ta+R(wb)*tb)
// }
