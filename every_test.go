package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestEvery(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t, []int{0, 2, 4, 6}, linq.Iota1(8).Every(2))
	assertQueryEqual(t, []int{10, 13, 16, 19}, linq.Iota1(20).EveryFrom(10, 3))
}
