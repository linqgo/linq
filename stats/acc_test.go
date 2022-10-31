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

package stats_test

import (
	"testing"

	"github.com/linqgo/linq"
	"github.com/linqgo/linq/stats"
)

func TestAccMean(t *testing.T) {
	t.Parallel()

	odds := linq.Iota3(1, 11, 2)

	assertQueryInEpsilon(t, []float64{2.0, 2.0, 2.0, 2.0, 2.0},
		stats.AccMean(linq.SlideAll(linq.Repeat(2.0, 5))), 1.001)
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, stats.AccMean(linq.SlideAll(odds)))
	assertQueryEqual(t, []int{2, 4, 6, 8}, stats.AccMean(linq.SlideFixed(odds, 2, false)))
	assertQueryEqual(t, []int{1, 2, 4, 6, 8}, stats.AccMean(linq.SlideFixed(odds, 2, true)))
}

// func TestAccMeanWt(t *testing.T) {
// 	t.Parallel()
// }

func TestAccGeometricMean(t *testing.T) {
	t.Parallel()

	assertQueryInEpsilon(t, []float64{2.0, 2.0, 2.0, 2.0, 2.0},
		stats.AccGeometricMean(linq.SlideAll(linq.Repeat(2.0, 5))), 1.001)
	assertQueryInEpsilon(t, []float64{1.0, 1.7321, 2.4662, 3.2011, 3.9363},
		stats.AccGeometricMean(linq.SlideAll(linq.Iota3(1.0, 11.0, 2.0))), 1.001)
}

func TestAccHarmonicMean(t *testing.T) {
	t.Parallel()

	assertQueryInEpsilon(t, []float64{2.0, 2.0, 2.0, 2.0, 2.0},
		stats.AccHarmonicMean(linq.SlideAll(linq.Repeat(2.0, 5))), 1.001)
	assertQueryInEpsilon(t, []float64{1.0, 1.5, 1.9565, 2.3864, 2.7975},
		stats.AccHarmonicMean(linq.SlideAll(linq.Iota3(1.0, 11.0, 2.0))), 1.001)
}

func TestAccProduct(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t, []int{2, 4, 8, 16, 32}, stats.AccProduct(linq.SlideAll(linq.Repeat(2, 5))))
	assertQueryEqual(t, []int{1, 2, 6, 24, 120, 720}, stats.AccProduct(linq.SlideAll(linq.Iota2(1, 7))))
}

func TestAccSum(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t, []int{1, 4, 9, 16, 25}, stats.AccSum(linq.SlideAll(linq.Iota3(1, 11, 2))))
}
