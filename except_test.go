package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestExcept(t *testing.T) {
	t.Parallel()

	f := linq.From[int]
	assertQueryEqual(t, []int{1, 3, 5}, linq.Except(f(1, 2, 3, 4, 5), f(2, 4)))
	assertQueryEqual(t, []int{}, linq.Except(f(1, 2, 3), f(1, 2, 3)))
	assertQueryEqual(t, []int{}, linq.Except(f(1, 2, 3), f(1, 2, 3, 4, 5)))
	assertQueryEqual(t, []int{1, 2, 3}, linq.Except(f(1, 2, 3), f(4, 5)))
	assertQueryEqual(t, []int{1, 2}, linq.Except(f(1, 2, 3), f(3, 4, 5)))
}
