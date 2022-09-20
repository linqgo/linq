package linq_test

import (
	"testing"

	"github.com/linqgo/linq"
)

func TestDistinct(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t, []int{}, linq.Distinct(linq.None[int]()))

	assertQueryEqual(t,
		[]int{1, 2, 3, 4, 5, 6, 7},
		linq.Distinct(linq.Concat(
			linq.From(1, 2, 3, 4, 5),
			linq.From(3, 4, 5, 6, 7),
		)),
	)

	assertOneShot(t, false, linq.Distinct(linq.From(1, 2, 3, 2, 3, 4, 3, 4, 5)))
	assertOneShot(t, true, linq.Distinct(oneshot()))

	assertFastCountEqual(t, 0, linq.Distinct(linq.None[int]()))
	assertFastCountEqual(t, 1, linq.Distinct(linq.From(1)))
	assertFastCountEqual(t, 1, linq.Distinct(linq.Concat(linq.None[int](), linq.From(1), linq.None[int]())))
	assertNoFastCount(t, linq.Distinct(linq.From(1, 2, 3, 2, 3, 4, 3, 4, 5)))
	assertNoFastCount(t, linq.Distinct(oneshot()))
}
