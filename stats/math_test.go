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

	"github.com/stretchr/testify/assert"
)

var (
	testNums  = linq.Select(linq.Iota2(1, 11), func(i int) float64 { return float64(i) })
	emptyNums = linq.None[float64]()
)

func TestAverage(t *testing.T) {
	t.Parallel()

	for _, data := range []linq.Query[float64]{testNums, linq.Reverse(testNums)} {
		assertSome(t, 5.5, stats.Mean(data))
		assertNo(t, stats.Mean(emptyNums))
	}
}

func TestGeometricMean(t *testing.T) {
	t.Parallel()

	for _, data := range []linq.Query[float64]{testNums, linq.Reverse(testNums)} {
		assertSomeInEpsilon(t, 4.529, stats.GeometricMean(data), 1.001)
		assertNo(t, stats.GeometricMean(emptyNums))
	}
}

func TestHarmonicMean(t *testing.T) {
	t.Parallel()

	for _, data := range []linq.Query[float64]{testNums, linq.Reverse(testNums)} {
		assertSomeInEpsilon(t, 3.414, stats.HarmonicMean(data), 1.001)
		assertNo(t, stats.HarmonicMean(emptyNums))
	}
}

func TestProduct(t *testing.T) {
	t.Parallel()

	for _, data := range []linq.Query[float64]{testNums, linq.Reverse(testNums)} {
		assert.EqualValues(t, 3628800, stats.Product(data))
	}
}

func TestSum(t *testing.T) {
	t.Parallel()

	for _, data := range []linq.Query[float64]{testNums, linq.Reverse(testNums)} {
		assert.EqualValues(t, 55, stats.Sum(data))
	}
}
