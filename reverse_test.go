package linq_test

import (
	"testing"

	"github.com/linqgo/linq"
)

func TestReverse(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t, []int{5, 4, 3, 2, 1}, linq.Iota2(1, 6).Reverse())

	assertOneShot(t, false, linq.Iota2(1, 6).Reverse())
	assertOneShot(t, true, oneshot().Reverse())

	assertSome(t, 5, linq.Iota2(1, 6).Reverse().FastCount())
	assertNo(t, oneshot().Reverse().FastCount())
}
