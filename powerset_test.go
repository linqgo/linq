package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestPowerSet(t *testing.T) {
	t.Parallel()

	powerset := func(q linq.Query[int]) linq.Query[[]int] {
		return linq.Select(linq.PowerSet(q), linq.ToSlice[int])
	}

	assertQueryElementsMatch(t, [][]int{nil}, powerset(linq.None[int]()))
	assertQueryElementsMatch(t, [][]int{nil, {1}}, powerset(linq.From(1)))
	assertQueryElementsMatch(t,
		[][]int{nil, {1}, {4}, {1, 4}},
		powerset(linq.From(1, 4)),
	)
	q := powerset(linq.From(1, 4, 9))
	assertQueryElementsMatch(t,
		[][]int{nil, {1}, {4}, {1, 4}, {9}, {1, 9}, {4, 9}, {1, 4, 9}},
		q)

	assertOneShot(t, false, q)
	assertOneShot(t, true, powerset(oneshot))
}
