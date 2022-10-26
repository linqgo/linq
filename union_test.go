package linq_test

import (
	"testing"

	"github.com/linqgo/linq"
)

func TestUnion(t *testing.T) {
	t.Parallel()

	f := linq.From[int]
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, linq.Union(f(1, 2, 3, 4, 5), f(2, 4)))
	assertQueryEqual(t, []int{1, 2, 3}, linq.Union(f(1, 2, 3), f(1, 2, 3)))
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, linq.Union(f(1, 2, 3), f(1, 2, 3, 4, 5)))
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, linq.Union(f(1, 2, 3), f(4, 5)))
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, linq.Union(f(1, 2, 3), f(3, 4, 5)))

	assertOneShot(t, false, linq.Union(f(1, 2, 3), f(3, 4, 5)))
	assertOneShot(t, true, linq.Union(oneshot(), f(3, 4, 5)))
	assertOneShot(t, true, linq.Union(f(1, 2, 3), oneshot()))
	assertOneShot(t, true, linq.Union(oneshot(), oneshot()))

	assertSome(t, 3, linq.Union(f(), f(3, 4, 5)).FastCount())
	assertSome(t, 3, linq.Union(f(1, 2, 3), f()).FastCount())
	assertNo(t, linq.Union(f(1, 2, 3), f(3, 4, 5)).FastCount())
	assertNo(t, linq.Union(oneshot(), f(3, 4, 5)).FastCount())
	assertNo(t, linq.Union(f(1, 2, 3), oneshot()).FastCount())
	assertNo(t, linq.Union(oneshot(), oneshot()).FastCount())
}
