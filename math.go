package linq

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Average returns the mean of the numbers in q or panic if q is empty.
func Average[N number](q Query[N]) N {
	sum, n := aggregate(q.Enumerator(), 0, add[N])
	return sum / N(n)
}

// Max returns the highest number in q.
func Max[N number](q Query[N]) N {
	return AggregateElse(q, max[N], N(math.NaN()))
}

// Min returns the lowest number in q.
func Min[N number](q Query[N]) N {
	return AggregateElse(q, min[N], N(math.NaN()))
}

// Sum returns the sum of the numbers in q.
func Sum[N number](q Query[N]) N {
	sum, _ := aggregate(q.Enumerator(), 0, add[N])
	return sum
}

type number interface {
	constraints.Integer | constraints.Float
}

func max[N number](a, b N) N {
	if a >= b {
		return a
	}
	return b
}

func min[N number](a, b N) N {
	if a <= b {
		return a
	}
	return b
}

func add[N number](a, b N) N {
	return a + b
}
