package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestIntersect(t *testing.T) {
	t.Parallel()

	f := linq.From[int]
	assertQueryEqual(t, []int{2, 4}, linq.Intersect(f(1, 2, 3, 4, 5), f(2, 4)))
	assertQueryEqual(t, []int{1, 2, 3}, linq.Intersect(f(1, 2, 3), f(1, 2, 3)))
	assertQueryEqual(t, []int{1, 2, 3}, linq.Intersect(f(1, 2, 3), f(1, 2, 3, 4, 5)))
	assertQueryEqual(t, []int{}, linq.Intersect(f(1, 2, 3), f(4, 5)))
	assertQueryEqual(t, []int{3}, linq.Intersect(f(1, 2, 3), f(3, 4, 5)))

	assertOneShot(t, false, linq.Intersect(f(1, 2, 3), f(3, 4, 5)))
	assertOneShot(t, false, linq.Intersect(oneshot(), f()))
	assertOneShot(t, true, linq.Intersect(oneshot(), f(3, 4, 5)))
	assertOneShot(t, false, linq.Intersect(f(), oneshot()))
	assertOneShot(t, true, linq.Intersect(f(1, 2, 3), oneshot()))
	assertOneShot(t, true, linq.Intersect(oneshot(), oneshot()))

	assertFastCountEqual(t, 0, linq.Intersect(f(), f()))
	assertFastCountEqual(t, 0, linq.Intersect(f(), slowcount))
	assertFastCountEqual(t, 0, linq.Intersect(slowcount, f()))
	assertNoFastCount(t, linq.Intersect(f(1, 2, 3), f(3, 4, 5)))
	assertNoFastCount(t, linq.Intersect(oneshot(), f(3, 4, 5)))
	assertNoFastCount(t, linq.Intersect(f(1, 2, 3), oneshot()))
	assertNoFastCount(t, linq.Intersect(oneshot(), oneshot()))
}
