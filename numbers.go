package linq

import (
	"math"

	"golang.org/x/exp/constraints"
)

type number interface {
	realNumber | constraints.Complex
}

type realNumber interface {
	constraints.Integer | constraints.Float
}

var nan = math.NaN()
