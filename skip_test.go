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

	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, data.SkipWhile(func(i int) bool { return i < 0 }))
	assertQueryEqual(t, []int{3, 4, 5}, data.SkipWhile(func(i int) bool { return i < 3 }))
	assertQueryEqual(t, []int{}, data.SkipWhile(func(i int) bool { return i < 10 }))
}
