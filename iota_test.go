package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestIota(t *testing.T) {
	t.Parallel()

	iota := linq.Iota[int]().Enumerator()
	for i := 0; i < 10; i++ {
		assertResultEqual(t, i, maybe(iota()))
	}
}

func TestIota123(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t, []int{}, linq.Iota1(0))
	assertQueryEqual(t, []int{0, 1, 2}, linq.Iota1(3))
	assertQueryEqual(t, []int{4, 5, 6}, linq.Iota2(4, 7))
}
