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

	assertSome(t, 0, linq.Flatten(linq.None[linq.Query[int]]()).FastCount())
	assertNo(t, q.FastCount())
	assertNo(t, linq.Flatten(linq.FromChannel(make(chan linq.Query[int]))).FastCount())
}

func TestFlattenSlices(t *testing.T) {
	t.Parallel()

	q := linq.FlattenSlices(linq.From([]int{1, 2}, []int{3, 4}))
	assertQueryEqual(t, []int{1, 2, 3, 4}, q)

	assertOneShot(t, false, q)
	assertOneShot(t, true, linq.FlattenSlices(linq.FromChannel(make(chan []int))))

	assertSome(t, 0, linq.FlattenSlices(linq.None[[]int]()).FastCount())
	assertNo(t, q.FastCount())
	assertNo(t, linq.FlattenSlices(linq.FromChannel(make(chan []int))).FastCount())
}
