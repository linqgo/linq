package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestWhere(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t,
		[]int{1, 3, 5},
		linq.Iota2(1, 6).Where(func(t int) bool { return t%2 == 1 }),
	)
}
