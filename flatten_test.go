package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestFlatten(t *testing.T) {
	t.Parallel()

	q := linq.Flatten(linq.From(linq.From(1, 2), linq.From(3, 4)))
	assertQueryEqual(t, []int{1, 2, 3, 4}, q)

	assertOneShot(t, false, q)
	assertOneShot(t, true, linq.Flatten(linq.FromChannel(make(chan linq.Query[int]))))
}

func TestFlattenSlices(t *testing.T) {
	t.Parallel()

	q := linq.FlattenSlices(linq.From([]int{1, 2}, []int{3, 4}))
	assertQueryEqual(t, []int{1, 2, 3, 4}, q)

	assertOneShot(t, false, q)
	assertOneShot(t, true, linq.FlattenSlices(linq.FromChannel(make(chan []int))))
}
