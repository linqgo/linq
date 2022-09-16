package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestChannel(t *testing.T) {
	t.Parallel()

	c := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		c <- i
	}
	close(c)
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, linq.FromChannel(c))

	assertOneShot(t, true, linq.FromChannel(c))
	assertNoFastCount(t, linq.FromChannel(c))
}
