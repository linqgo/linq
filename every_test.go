package linq_test

import (
	"testing"

	"github.com/linqgo/linq"
)

func TestEvery(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t, []int{0, 2, 4, 6}, linq.Iota1(8).Every(2))
	assertQueryEqual(t, []int{10, 13, 16, 19}, linq.Iota1(20).EveryFrom(10, 3))

	assertOneShot(t, false, linq.Iota2(1, 6).Every(2))
	assertOneShot(t, true, oneshot().Every(2))
	assertOneShot(t, false, linq.Iota2(1, 6).EveryFrom(10, 3))
	assertOneShot(t, true, oneshot().EveryFrom(10, 3))

	assertSome(t, 3, linq.Iota2(1, 6).Every(2).FastCount())
	assertSome(t, 3, linq.Iota2(1, 7).Every(2).FastCount())
	assertSome(t, 4, linq.Iota2(1, 8).Every(2).FastCount())

	assertSome(t, 4, linq.Iota1(20).EveryFrom(10, 3).FastCount())
	assertSome(t, 4, linq.Iota1(21).EveryFrom(10, 3).FastCount())
	assertSome(t, 4, linq.Iota1(22).EveryFrom(10, 3).FastCount())
	assertSome(t, 5, linq.Iota1(23).EveryFrom(10, 3).FastCount())

	assertNo(t, oneshot().Every(2).FastCount())
	assertNo(t, oneshot().EveryFrom(10, 3).FastCount())
}
