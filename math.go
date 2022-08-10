package linq

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Average returns the arithmetic mean of the numbers in q or ok=false if q is
// empty.
func Average[R realNumber](q Query[R]) (mean R, ok bool) {
	return aggregateThen(q.Enumerator(), 0, add[R], func(sum R, n int) R {
		return sum / R(n)
	})
}

// AverageOrNan returns the arithmetic mean of the numbers in q or NaN if q is
// empty.
func AverageOrNaN[R realNumber](q Query[R]) R {
	return valueOrNaN(Average(q))
}

// MustAverage returns the arithmetic mean of the numbers in q or panics if q is
// empty.
func MustAverage[R realNumber](q Query[R]) R {
	return valueOrPanicEmpty(Average(q))
}

// GeometricMean returns the geometric mean of the numbers in q or ok=false if q
// is empty.
func GeometricMean[R realNumber](q Query[R]) (mean R, ok bool) {
	return aggregateThen(q.Enumerator(), 0, mul[R], func(product R, n int) R {
		return R(math.Pow(float64(product), float64(n)))
	})
}

// GeometricMeanOrNan returns the geometric mean of the numbers in q or NaN if q
// is empty.
func GeometricMeanOrNaN[R realNumber](q Query[R]) R {
	return valueOrNaN(GeometricMean(q))
}

// MustGeometricMean returns the geometric mean of the numbers in q or panics if
// q is empty.
func MustGeometricMean[R realNumber](q Query[R]) R {
	return valueOrPanicEmpty(GeometricMean(q))
}

// HarmonicMean returns the harmonic mean of the numbers in q or ok=false if q
// is empty.
func HarmonicMean[F constraints.Float](q Query[F]) (mean F, ok bool) {
	return aggregateThen(q.Enumerator(), 0, recipAdd[F], func(recipSum F, n int) F {
		return F(n) / F(recipSum)
	})
}

// HarmonicMeanOrNan returns the harmonic mean of the numbers in q or NaN if q
// is empty.
func HarmonicMeanOrNaN[F constraints.Float](q Query[F]) F {
	return valueOrNaN(HarmonicMean(q))
}

// MustHarmonicMean returns the harmonic mean of the numbers in q or panics if q
// is empty.
func MustHarmonicMean[F constraints.Float](q Query[F]) F {
	return valueOrPanicEmpty(HarmonicMean(q))
}

// Max returns the highest number in q or ok=false if q is empty.
func Max[R realNumber](q Query[R]) (_ R, ok bool) {
	return Aggregate(q, max[R])
}

// MaxOrNaN returns the highest number in q or NaN if .
func MaxOrNaN[R realNumber](q Query[R]) R {
	return valueOrNaN(Max(q))
}

// MustMax returns the highest number in q or panics if q is empty.
func MustMax[R realNumber](q Query[R]) R {
	return valueOrPanicEmpty(Max(q))
}

// Min returns the highest number in q or ok=false if q is empty.
func Min[R realNumber](q Query[R]) (_ R, ok bool) {
	return Aggregate(q, min[R])
}

// MinOrNaN returns the highest number in q or NaN if q is empty.
func MinOrNaN[R realNumber](q Query[R]) R {
	return valueOrNaN(Min(q))
}

// MustMin returns the lowest number in q or panics of q is empty.
func MustMin[R realNumber](q Query[R]) R {
	return valueOrPanicEmpty(Min(q))
}

// Product returns the product of the numbers in q or 1 if q is empty.
func Product[R number](q Query[R]) R {
	return aggregate(q.Enumerator(), 1, mul[R])
}

// Sum returns the sum of the numbers in q or 0 if q is empty.
func Sum[R number](q Query[R]) R {
	return aggregate(q.Enumerator(), 0, add[R])
}

type number interface {
	realNumber | constraints.Complex
}

type realNumber interface {
	constraints.Integer | constraints.Float
}

var nan = math.NaN()

func max[R realNumber](a, b R) R {
	if a >= b {
		return a
	}
	return b
}

func min[R realNumber](a, b R) R {
	if a <= b {
		return a
	}
	return b
}

func add[N number](a, b N) N {
	return a + b
}

func mul[N number](a, b N) N {
	return a * b
}

func recipAdd[R constraints.Float](a, b R) R {
	return a + 1/b
}

func valueOrNaN[R realNumber](r R, ok bool) R {
	return valueElse(r, ok, R(nan))
}

func valueOrPanicEmpty[R realNumber](r R, ok bool) R {
	return valueOrPanic(r, ok, emptySourceError)
}
