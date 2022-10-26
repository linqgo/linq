package linq_test

import (
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
		assertSome(t, 5.5, linq.Average(data))
		assertNo(t, linq.Average(emptyNums))
	}
}

func TestGeometricMean(t *testing.T) {
	t.Parallel()

	for _, data := range []linq.Query[float64]{testNums, linq.Reverse(testNums)} {
		assertSomeInEpsilon(t, 4.529, linq.GeometricMean(data), 1.001)
		assertNo(t, linq.GeometricMean(emptyNums))
	}
}

func TestHarmonicMean(t *testing.T) {
	t.Parallel()

	for _, data := range []linq.Query[float64]{testNums, linq.Reverse(testNums)} {
		assertSomeInEpsilon(t, 3.414, linq.HarmonicMean(data), 1.001)
		assertNo(t, linq.HarmonicMean(emptyNums))
	}
}

func TestMax(t *testing.T) {
	t.Parallel()

	for _, data := range []linq.Query[float64]{testNums, linq.Reverse(testNums)} {
		assertSome(t, 10.0, linq.Max(data))
		assertNo(t, linq.Max(emptyNums))
	}
}

func TestMaxBy(t *testing.T) {
	t.Parallel()

	type Person = linq.KV[string, int]

	peeps := linq.FromMap(map[string]int{"John": 42, "Sanjiv": 22, "Andrea": 35})
	noone := linq.FromMap(map[string]int{})

	name := func(kv Person) string { return kv.Key }
	age := func(kv Person) int { return kv.Value }

	assertSome(t, linq.NewKV("Sanjiv", 22), linq.MaxBy(peeps, name))
	assertNo(t, linq.MaxBy(noone, name))
	assertSome(t, linq.NewKV("John", 42), linq.MaxBy(peeps, age))
	assertNo(t, linq.MaxBy(noone, age))
}

func TestMin(t *testing.T) {
	t.Parallel()

	for _, data := range []linq.Query[float64]{testNums, linq.Reverse(testNums)} {
		assertSome(t, 1.0, linq.Min(data))
		assertNo(t, linq.Min(emptyNums))
	}
}

func TestMinBy(t *testing.T) {
	t.Parallel()

	type Person = linq.KV[string, int]

	peeps := linq.FromMap(map[string]int{"John": 42, "Andrea": 35, "Sanjiv": 22})
	noone := linq.FromMap(map[string]int{})

	name := func(kv Person) string { return kv.Key }
	age := func(kv Person) int { return kv.Value }

	assertSome(t, linq.NewKV("Andrea", 35), linq.MinBy(peeps, name))
	assertNo(t, linq.MinBy(noone, name))
	assertSome(t, linq.NewKV("Sanjiv", 22), linq.MinBy(peeps, age))
	assertNo(t, linq.MinBy(noone, age))
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
