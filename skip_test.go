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

	assertOneShot(t, false, data.Skip(0))
	assertOneShot(t, true, oneshot.Skip(0))
}

func TestSkipLast(t *testing.T) {
	t.Parallel()

	data := linq.From(1, 2, 3, 4, 5)

	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, data.SkipLast(0))
	assertQueryEqual(t, []int{1, 2, 3}, data.SkipLast(2))
	assertQueryEqual(t, []int{}, data.SkipLast(10))

	assertOneShot(t, false, data.SkipLast(0))
	assertOneShot(t, true, oneshot.SkipLast(0))
}

func TestSkipWhile(t *testing.T) {
	t.Parallel()

	data := linq.From(1, 2, 3, 4, 5)

	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, data.SkipWhile(func(x int) bool { return x < 0 }))
	assertQueryEqual(t, []int{3, 4, 5}, data.SkipWhile(func(x int) bool { return x < 3 }))
	assertQueryEqual(t, []int{}, data.SkipWhile(func(x int) bool { return x < 10 }))

	assertOneShot(t, false, data.SkipWhile(linq.False[int]))
	assertOneShot(t, true, oneshot.SkipWhile(linq.False[int]))
}
