package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestEvery(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t, []int{0, 2, 4, 6}, linq.Iota1(8).Every(2))
	assertQueryEqual(t, []int{10, 13, 16, 19}, linq.Iota1(20).EveryFrom(10, 3))

	assertOneShot(t, false, linq.Iota2(1, 6).Every(2))
	assertOneShot(t, true, oneshot.Every(2))
	assertOneShot(t, false, linq.Iota2(1, 6).EveryFrom(10, 3))
	assertOneShot(t, true, oneshot.EveryFrom(10, 3))
}
