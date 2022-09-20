package linq_test

import (
	"testing"

	"github.com/linqgo/linq"
)

func TestFlatten(t *testing.T) {
	t.Parallel()

	q := linq.Flatten(linq.From(linq.From(1, 2), linq.From(3, 4)))
	assertQueryEqual(t, []int{1, 2, 3, 4}, q)

	assertOneShot(t, false, q)
	assertOneShot(t, true, linq.Flatten(linq.FromChannel(make(chan linq.Query[int]))))

	assertFastCountEqual(t, 0, linq.Flatten(linq.None[linq.Query[int]]()))
	assertNoFastCount(t, q)
	assertNoFastCount(t, linq.Flatten(linq.FromChannel(make(chan linq.Query[int]))))
}

func TestFlattenSlices(t *testing.T) {
	t.Parallel()

	q := linq.FlattenSlices(linq.From([]int{1, 2}, []int{3, 4}))
	assertQueryEqual(t, []int{1, 2, 3, 4}, q)

	assertOneShot(t, false, q)
	assertOneShot(t, true, linq.FlattenSlices(linq.FromChannel(make(chan []int))))

	assertFastCountEqual(t, 0, linq.FlattenSlices(linq.None[[]int]()))
	assertNoFastCount(t, q)
	assertNoFastCount(t, linq.FlattenSlices(linq.FromChannel(make(chan []int))))
}
