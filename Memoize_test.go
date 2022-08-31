package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

var fiveInts = func() <-chan int {
	c := make(chan int, 5)
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
	return c
}

func TestMemoize(t *testing.T) {
	t.Parallel()

	q := linq.FromChannel(fiveInts())

	assertQueryEqual(t, []int{0, 1, 2, 3, 4}, q)
	assertQueryEqual(t, []int{}, q)

	m := linq.FromChannel(fiveInts()).Memoize()

	assertQueryEqual(t, []int{0, 1, 2, 3, 4}, m)
	assertQueryEqual(t, []int{0, 1, 2, 3, 4}, m)
}

func TestMemoizeParalle(t *testing.T) {
	t.Parallel()

	m := linq.FromChannel(fiveInts()).Memoize()
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				assertQueryEqual(t, []int{0, 1, 2, 3, 4}, m)
			}
		}()
	}
}
