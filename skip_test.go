package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestSkip(t *testing.T) {
	t.Parallel()

	data := linq.From(1, 2, 3, 4, 5)

	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, data.Skip(0))
	assertQueryEqual(t, []int{3, 4, 5}, data.Skip(2))
	assertQueryEqual(t, []int{}, data.Skip(10))
}

func TestSkipLast(t *testing.T) {
	t.Parallel()

	data := linq.From(1, 2, 3, 4, 5)

	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, data.SkipLast(0))
	assertQueryEqual(t, []int{1, 2, 3}, data.SkipLast(2))
	assertQueryEqual(t, []int{}, data.SkipLast(10))
}

func TestSkipWhile(t *testing.T) {
	t.Parallel()

	data := linq.From(1, 2, 3, 4, 5)

	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, data.SkipWhile(func(x int) bool { return x < 0 }))
	assertQueryEqual(t, []int{3, 4, 5}, data.SkipWhile(func(x int) bool { return x < 3 }))
	assertQueryEqual(t, []int{}, data.SkipWhile(func(x int) bool { return x < 10 }))
}

func TestSkipWhileI(t *testing.T) {
	t.Parallel()

	data := linq.Iota3(5, 0, -1)

	assertQueryEqual(t, []int{}, data.SkipWhileI(func(i, x int) bool { return i*x < 10 }))
	assertQueryEqual(t, []int{2, 1}, data.SkipWhileI(func(i, x int) bool { return i < x }))
	assertQueryEqual(t, []int{5, 4, 3, 2, 1}, data.SkipWhileI(func(i, x int) bool { return i*x > 0 }))
}
