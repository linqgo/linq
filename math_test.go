package linq_test

import (
	"math"
	"testing"

	"github.com/marcelocantos/linq"

	"github.com/stretchr/testify/assert"
)

func TestMath(t *testing.T) {
	t.Parallel()

	data := linq.Select(linq.Iota2(1, 11), func(i int) float64 { return float64(i) })
	empty := linq.None[float64]()

	for _, data := range []linq.Query[float64]{data, linq.Reverse(data)} {
		assert.EqualValues(t, 5.5, linq.MustAverage(data))
		assert.PanicsWithError(t, "empty source", func() { linq.MustAverage(empty) })
		assert.EqualValues(t, 5.5, linq.AverageOrNaN(data))
		assert.True(t, math.IsNaN(linq.AverageOrNaN(empty)))

		assert.InEpsilon(t, 4.529, linq.MustGeometricMean(data), 1.001)
		assert.PanicsWithError(t, "empty source", func() { linq.MustGeometricMean(empty) })
		assert.InEpsilon(t, 4.529, linq.GeometricMeanOrNaN(data), 1.001)
		assert.True(t, math.IsNaN(linq.GeometricMeanOrNaN(empty)))

		assert.InEpsilon(t, 3.414, linq.MustHarmonicMean(data), 1.001)
		assert.PanicsWithError(t, "empty source", func() { linq.MustHarmonicMean(empty) })
		assert.InEpsilon(t, 3.414, linq.HarmonicMeanOrNaN(data), 1.001)
		assert.True(t, math.IsNaN(linq.HarmonicMeanOrNaN(empty)))

		assert.EqualValues(t, 10, linq.MustMax(data))
		assert.Panics(t, func() { linq.MustMax(empty) })
		assert.EqualValues(t, 10, linq.MaxOrNaN(data))
		assert.True(t, math.IsNaN(linq.MaxOrNaN(empty)))

		assert.EqualValues(t, 1, linq.MustMin(data))
		assert.Panics(t, func() { linq.MustMin(empty) })
		assert.EqualValues(t, 1, linq.MinOrNaN(data))
		assert.True(t, math.IsNaN(linq.MinOrNaN(empty)))

		assert.EqualValues(t, 3628800, linq.Product(data))

		assert.EqualValues(t, 55, linq.Sum(data))
	}
}
