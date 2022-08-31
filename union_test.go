package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestUnion(t *testing.T) {
	t.Parallel()

	f := linq.From[int]
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, linq.Union(f(1, 2, 3, 4, 5), f(2, 4)))
	assertQueryEqual(t, []int{1, 2, 3}, linq.Union(f(1, 2, 3), f(1, 2, 3)))
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, linq.Union(f(1, 2, 3), f(1, 2, 3, 4, 5)))
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, linq.Union(f(1, 2, 3), f(4, 5)))
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, linq.Union(f(1, 2, 3), f(3, 4, 5)))
}
