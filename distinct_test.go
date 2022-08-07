package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestDistinct(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t, []int{}, linq.Distinct(linq.None[int]()))

	assertQueryEqual(t,
		[]int{1, 2, 3, 4, 5, 6, 7},
		linq.Distinct(linq.Concat(
			linq.From(1, 2, 3, 4, 5),
			linq.From(3, 4, 5, 6, 7),
		)),
	)
}
