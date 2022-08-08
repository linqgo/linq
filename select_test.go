package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestSelect(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t,
		[]int{0, 1, 4, 9, 16},
		linq.Select(linq.Iota1(5), func(x int) int { return x * x }),
	)
}

func TestSelectI(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t,
		[]int{0, 4, 6, 6, 4},
		linq.SelectI(linq.Iota3(5, 0, -1), func(i, x int) int { return i * x }),
	)
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

	assertQueryEqual(t,
		[]int{2, 3, 7, 2, 2, 2, 7},
		linq.SelectMany(linq.From(42, 56), func(e int) linq.Query[int] {
			return primeFactors(e)
		}),
	)
}

func TestSelectManyI(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t,
		[]int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1},
		linq.SelectManyI(linq.Iota3(5, 0, -1), func(i, x int) linq.Query[int] {
			return linq.Repeat(x, i)
		}),
	)
}
