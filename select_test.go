package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestSelect(t *testing.T) {
	t.Parallel()

	square := func(x int) int { return x * x }
	q := linq.Select(linq.Iota1(5), square)
	assertQueryEqual(t, []int{0, 1, 4, 9, 16}, q)

	assertOneShot(t, false, q)
	assertOneShot(t, true, linq.Select(oneshot(), square))

	assertFastCountEqual(t, 5, q)
	assertNoFastCount(t, linq.Select(oneshot(), square))
}

func primeFactors(n int) linq.Query[int] {
	return linq.NewQuery(func() linq.Enumerator[int] {
		i, s := 2, 1
		return func() (int, bool) {
			for ; i <= n; i, s = i+s, 2 {
				if n%i == 0 {
					n /= i
					return i, true
				}
			}
			return 0, false
		}
	})
}

func TestSelectMany(t *testing.T) {
	t.Parallel()

	q := linq.SelectMany(linq.From(42, 56), primeFactors)
	assertQueryEqual(t, []int{2, 3, 7, 2, 2, 2, 7}, q)

	assertOneShot(t, false, q)
	assertOneShot(t, true, linq.SelectMany(oneshot(), primeFactors))

	assertNoFastCount(t, q)
	assertNoFastCount(t, linq.SelectMany(oneshot(), primeFactors))
}
