package linq_test

import (
	"testing"

	"github.com/linqgo/linq"
)

func TestExcept(t *testing.T) {
	t.Parallel()

	f := linq.From[int]
	assertQueryEqual(t, []int{1, 3, 5}, linq.Except(f(1, 2, 3, 4, 5), f(2, 4)))
	assertQueryEqual(t, []int{}, linq.Except(f(1, 2, 3), f(1, 2, 3)))
	assertQueryEqual(t, []int{}, linq.Except(f(1, 2, 3), f(1, 2, 3, 4, 5)))
	assertQueryEqual(t, []int{1, 2, 3}, linq.Except(f(1, 2, 3), f(4, 5)))
	assertQueryEqual(t, []int{1, 2}, linq.Except(f(1, 2, 3), f(3, 4, 5)))

	assertOneShot(t, false, linq.Except(f(1, 2, 3), f(3, 4, 5)))
	assertOneShot(t, false, linq.Except(f(), oneshot()))
	assertOneShot(t, true, linq.Except(oneshot(), f(3, 4, 5)))
	assertOneShot(t, true, linq.Except(f(1, 2, 3), oneshot()))
	assertOneShot(t, true, linq.Except(oneshot(), oneshot()))

	assertNoFastCount(t, linq.Except(f(1, 2, 3), f(3, 4, 5)))
	assertFastCountEqual(t, 3, linq.Except(f(1, 2, 3), f()))
	assertFastCountEqual(t, 0, linq.Except(f(), oneshot()))
	assertNoFastCount(t, linq.Except(oneshot(), f(3, 4, 5)))
	assertNoFastCount(t, linq.Except(f(1, 2, 3), oneshot()))
	assertNoFastCount(t, linq.Except(oneshot(), oneshot()))
}
