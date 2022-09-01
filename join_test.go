package linq_test

import (
	"fmt"
	"testing"

	"github.com/marcelocantos/linq"
)

func TestJoin(t *testing.T) {
	t.Parallel()

	type XY struct{ x, y int }
	type YZ struct{ y, z int }
	type XYZ struct{ x, y, z int }

	a := []XY{{1, 20}, {1, 30}, {2, 30}, {2, 40}}
	b := []YZ{{30, 200}, {30, 250}, {20, 100}}

	assertQueryElementsMatch(t,
		[]XYZ{{1, 20, 100}, {1, 30, 200}, {1, 30, 250}, {2, 30, 200}, {2, 30, 250}},
		linq.Join(
			linq.From(a...),
			linq.From(b...),
			func(e XY) int { return e.y },
			func(e YZ) int { return e.y },
			func(a XY, b YZ) XYZ { return XYZ{a.x, a.y, b.z} },
		),
	)

	assertQueryElementsMatch(t,
		[]XYZ{{1, 20, 100}, {1, 30, 200}, {1, 30, 250}, {2, 30, 200}, {2, 30, 250}},
		linq.Join(
			linq.From(b...),
			linq.From(a...),
			func(e YZ) int { return e.y },
			func(e XY) int { return e.y },
			func(b YZ, a XY) XYZ { return XYZ{a.x, a.y, b.z} },
		),
	)

	f := func(a, b linq.Query[int]) linq.Query[string] {
		return linq.Join(
			a,
			b,
			func(i int) int { return i % 2 },
			func(i int) int { return i % 3 },
			func(a, b int) string { return fmt.Sprintf("%d,%d", a, b) },
		)
	}
	assertQueryElementsMatch(t,
		[]string{"0,6", "1,7", "0,9", "2,6", "2,9", "3,7", "4,6", "4,9"},
		f(linq.Iota2(0, 5), linq.Iota2(5, 10)))

	assertOneShot(t, false, f(linq.Iota2(0, 5), linq.Iota2(5, 10)))
	assertOneShot(t, true, f(oneshot, linq.Iota2(5, 10)))
	assertOneShot(t, true, f(linq.Iota2(0, 5), oneshot))
	assertOneShot(t, true, f(oneshot, oneshot))
}
