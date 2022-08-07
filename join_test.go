package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestJoin(t *testing.T) {
	t.Parallel()

	type XY struct{ x, y int }
	type YZ struct{ y, z int }
	type XYZ struct{ x, y, z int }

	A := []XY{{1, 20}, {1, 30}, {2, 30}, {2, 40}}
	B := []YZ{{30, 200}, {30, 250}, {20, 100}}

	assertQueryEqual(t,
		[]XYZ{{1, 20, 100}, {1, 30, 200}, {1, 30, 250}, {2, 30, 200}, {2, 30, 250}},
		linq.Join(
			linq.From(A...),
			linq.From(B...),
			func(e XY) int { return e.y },
			func(e YZ) int { return e.y },
			func(a XY, b YZ) XYZ { return XYZ{a.x, a.y, b.z} },
		),
	)
}
