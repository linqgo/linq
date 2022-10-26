package linq_test

import (
	"testing"

	"github.com/linqgo/linq"
)

func TestFromGetter(t *testing.T) {
	t.Parallel()

	g := linq.FromGetter(func(i int) linq.Maybe[int] {
		if 0 <= i && i < 5 {
			return linq.Some(i * i)
		}
		return linq.No[int]()
	})

	assertQueryEqual(t, []int{0, 1, 4, 9, 16}, g)
	assertSome(t, 0, g.FastElementAt(0))
	assertSome(t, 16, g.FastElementAt(4))
	assertNo(t, g.FastElementAt(5))
	assertNo(t, g.FastElementAt(-1))
}

func TestToGetter(t *testing.T) {
	t.Parallel()

	g := linq.Iota1(10).Select(func(i int) int { return i * i }).ToGetter()

	assertSome(t, 0, g(0))
	assertSome(t, 36, g(6))
	assertNo(t, g(10))
	assertNo(t, g(-1))
}
