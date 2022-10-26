package linq_test

import (
	"testing"

	"github.com/linqgo/linq"
)

func TestSelect(t *testing.T) {
	t.Parallel()

	square := func(x int) int { return x * x }
	q := linq.Iota1(5).Select(square)
	assertQueryEqual(t, []int{0, 1, 4, 9, 16}, q)

	assertOneShot(t, false, q)
	assertOneShot(t, true, oneshot().Select(square))

	assertSome(t, 5, q.FastCount())
	assertNo(t, oneshot().Select(square).FastCount())
}

func primeFactors(n int) linq.Query[int] {
	return linq.NewQuery(func() linq.Enumerator[int] {
		i, s := 2, 1
		return func() linq.Maybe[int] {
			for ; i <= n; i, s = i+s, 2 {
				if n%i == 0 {
					n /= i
					return linq.Some(i)
				}
			}
			return linq.No[int]()
		}
	})
}

func TestSelectMany(t *testing.T) {
	t.Parallel()

	q := linq.SelectMany(linq.From(42, 56), primeFactors)
	assertQueryEqual(t, []int{2, 3, 7, 2, 2, 2, 7}, q)

	assertOneShot(t, false, q)
	assertOneShot(t, true, linq.SelectMany(oneshot(), primeFactors))

	assertNo(t, q.FastCount())
	assertNo(t, linq.SelectMany(oneshot(), primeFactors).FastCount())
}
