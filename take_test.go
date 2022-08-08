package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestTake(t *testing.T) {
	t.Parallel()

	data := linq.From(1, 2, 3, 4, 5)

	assertQueryEqual(t, []int{}, data.Take(0))
	assertQueryEqual(t, []int{1, 2}, data.Take(2))
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, data.Take(10))
}

func TestTakeLast(t *testing.T) {
	t.Parallel()

	data := linq.From(1, 2, 3, 4, 5)

	assertQueryEqual(t, []int{}, data.TakeLast(0))
	assertQueryEqual(t, []int{4, 5}, data.TakeLast(2))
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, data.TakeLast(10))
}

func TestTakeWhile(t *testing.T) {
	t.Parallel()

	data := linq.From(1, 2, 3, 4, 5)

	assertQueryEqual(t, []int{}, data.TakeWhile(func(i int) bool { return i < 0 }))
	assertQueryEqual(t, []int{1, 2}, data.TakeWhile(func(i int) bool { return i < 3 }))
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, data.TakeWhile(func(i int) bool { return i < 10 }))
}

func TestTakeWhileI(t *testing.T) {
	t.Parallel()

	data := linq.Iota3(5, 0, -1)

	assertQueryEqual(t, []int{5, 4, 3, 2, 1}, data.TakeWhileI(func(i, x int) bool { return i*x < 10 }))
	assertQueryEqual(t, []int{5, 4, 3}, data.TakeWhileI(func(i, x int) bool { return i < x }))
	assertQueryEqual(t, []int{}, data.TakeWhileI(func(i, x int) bool { return i*x > 0 }))
}
