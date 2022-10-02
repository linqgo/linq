package linq_test

import (
	"math"
	"testing"

	"github.com/linqgo/linq"

	"github.com/stretchr/testify/assert"
)

var (
	testNums  = linq.Select(linq.Iota2(1, 11), func(i int) float64 { return float64(i) })
	emptyNums = linq.None[float64]()
)

func TestAverage(t *testing.T) {
	t.Parallel()

	for _, data := range []linq.Query[float64]{testNums, linq.Reverse(testNums)} {
		assert.EqualValues(t, 5.5, linq.MustAverage(data))
		assert.PanicsWithError(t, "empty source", func() { linq.MustAverage(emptyNums) })
		assert.EqualValues(t, 5.5, linq.AverageOrNaN(data))
		assert.True(t, math.IsNaN(linq.AverageOrNaN(emptyNums)))
		assert.EqualValues(t, 5.5, linq.AverageElse(data, 42.0))
		assert.EqualValues(t, 42.0, linq.AverageElse(emptyNums, 42.0))
	}
}

func TestGeometricMean(t *testing.T) {
	t.Parallel()

	for _, data := range []linq.Query[float64]{testNums, linq.Reverse(testNums)} {
		assert.InEpsilon(t, 4.529, linq.MustGeometricMean(data), 1.001)
		assert.PanicsWithError(t, "empty source", func() { linq.MustGeometricMean(emptyNums) })
		assert.InEpsilon(t, 4.529, linq.GeometricMeanOrNaN(data), 1.001)
		assert.True(t, math.IsNaN(linq.GeometricMeanOrNaN(emptyNums)))
		assert.InEpsilon(t, 4.529, linq.GeometricMeanElse(data, 42.0), 1.001)
		assert.EqualValues(t, 42.0, linq.GeometricMeanElse(emptyNums, 42.0))
	}
}

func TestHarmonicMean(t *testing.T) {
	t.Parallel()

	for _, data := range []linq.Query[float64]{testNums, linq.Reverse(testNums)} {
		assert.InEpsilon(t, 3.414, linq.MustHarmonicMean(data), 1.001)
		assert.PanicsWithError(t, "empty source", func() { linq.MustHarmonicMean(emptyNums) })
		assert.InEpsilon(t, 3.414, linq.HarmonicMeanOrNaN(data), 1.001)
		assert.True(t, math.IsNaN(linq.HarmonicMeanOrNaN(emptyNums)))
		assert.InEpsilon(t, 3.414, linq.HarmonicMeanElse(data, 42.0), 1.001)
		assert.EqualValues(t, 42.0, linq.HarmonicMeanElse(emptyNums, 42.0))
	}
}

func TestMax(t *testing.T) {
	t.Parallel()

	for _, data := range []linq.Query[float64]{testNums, linq.Reverse(testNums)} {
		assert.EqualValues(t, 10, linq.MustMax(data))
		assert.Panics(t, func() { linq.MustMax(emptyNums) })
		assert.EqualValues(t, 10, linq.MaxOrNaN(data))
		assert.True(t, math.IsNaN(linq.MaxOrNaN(emptyNums)))
		assert.EqualValues(t, 10, linq.MaxElse(data, 42.0))
		assert.EqualValues(t, 42.0, linq.MaxElse(emptyNums, 42.0))
	}
}

func TestMaxBy(t *testing.T) {
	t.Parallel()

	type Person = linq.KV[string, int]

	peeps := linq.FromMap(map[string]int{"John": 42, "Sanjiv": 22, "Andrea": 35})
	noone := linq.FromMap(map[string]int{})

	name := func(kv Person) string { return kv.Key }
	age := func(kv Person) int { return kv.Value }

	assert.EqualValues(t, linq.NewKV("Sanjiv", 22), linq.MustMaxBy(peeps, name))
	assert.Panics(t, func() { linq.MustMaxBy(noone, name) })
	assert.EqualValues(t, linq.NewKV("John", 42), linq.MustMaxBy(peeps, age))
	assert.Panics(t, func() { linq.MustMaxBy(noone, age) })
}

func TestMin(t *testing.T) {
	t.Parallel()

	for _, data := range []linq.Query[float64]{testNums, linq.Reverse(testNums)} {
		assert.EqualValues(t, 1, linq.MustMin(data))
		assert.Panics(t, func() { linq.MustMin(emptyNums) })
		assert.EqualValues(t, 1, linq.MinOrNaN(data))
		assert.True(t, math.IsNaN(linq.MinOrNaN(emptyNums)))
		assert.EqualValues(t, 1, linq.MinElse(data, 42.0))
		assert.EqualValues(t, 42.0, linq.MinElse(emptyNums, 42.0))
	}
}

func TestMinBy(t *testing.T) {
	t.Parallel()

	type Person = linq.KV[string, int]

	peeps := linq.FromMap(map[string]int{"John": 42, "Andrea": 35, "Sanjiv": 22})
	noone := linq.FromMap(map[string]int{})

	name := linq.Key[Person]
	age := linq.Value[Person]

	assert.EqualValues(t, linq.NewKV("Andrea", 35), linq.MustMinBy(peeps, name))
	assert.Panics(t, func() { linq.MustMinBy(noone, name) })
	assert.EqualValues(t, linq.NewKV("Sanjiv", 22), linq.MustMinBy(peeps, age))
	assert.Panics(t, func() { linq.MustMinBy(noone, age) })
}

func TestProduct(t *testing.T) {
	t.Parallel()

	for _, data := range []linq.Query[float64]{testNums, linq.Reverse(testNums)} {
		assert.EqualValues(t, 3628800, linq.Product(data))
	}
}

func TestSum(t *testing.T) {
	t.Parallel()

	for _, data := range []linq.Query[float64]{testNums, linq.Reverse(testNums)} {
		assert.EqualValues(t, 55, linq.Sum(data))
	}
}
