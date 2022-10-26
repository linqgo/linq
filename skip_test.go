package linq_test

import (
	"testing"

	"github.com/linqgo/linq"
)

func TestSkip(t *testing.T) {
	t.Parallel()

	data := linq.From(1, 2, 3, 4, 5)

	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, data.Skip(0))
	assertQueryEqual(t, []int{3, 4, 5}, data.Skip(2))
	assertQueryEqual(t, []int{}, data.Skip(10))
	assertQueryEqual(t, []int{}, oneshot().Skip(2))

	assertOneShot(t, false, data.Skip(0))
	assertOneShot(t, true, oneshot().Skip(0))
	assertOneShot(t, true, oneshot().Skip(2))

	assertSome(t, 5, data.Skip(0).FastCount())
	assertSome(t, 3, data.Skip(2).FastCount())
	assertNo(t, oneshot().Skip(0).FastCount())
}

func TestSkipLast(t *testing.T) {
	t.Parallel()

	data := linq.From(1, 2, 3, 4, 5)

	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, data.SkipLast(0))
	assertQueryEqual(t, []int{1, 2, 3}, data.SkipLast(2))
	assertQueryEqual(t, []int{}, data.SkipLast(10))
	assertQueryEqual(t, []int{}, linq.From(1, 2, 3).Where(linq.False[int]).SkipLast(10))

	assertOneShot(t, false, data.SkipLast(0))
	assertOneShot(t, true, oneshot().SkipLast(0))

	assertSome(t, 5, data.SkipLast(0).FastCount())
	assertSome(t, 2, data.SkipLast(3).FastCount())
	assertNo(t, oneshot().SkipLast(0).FastCount())
}

func TestSkipWhile(t *testing.T) {
	t.Parallel()

	data := linq.From(1, 2, 3, 4, 5)

	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, data.SkipWhile(func(x int) bool { return x < 0 }))
	assertQueryEqual(t, []int{3, 4, 5}, data.SkipWhile(func(x int) bool { return x < 3 }))
	assertQueryEqual(t, []int{}, data.SkipWhile(func(x int) bool { return x < 10 }))

	assertOneShot(t, false, data.SkipWhile(linq.False[int]))
	assertOneShot(t, true, oneshot().SkipWhile(linq.False[int]))

	assertNo(t, data.SkipWhile(linq.False[int]).FastCount())
	assertNo(t, oneshot().SkipWhile(linq.False[int]).FastCount())
}

func TestSkipElementAt(t *testing.T) {
	t.Parallel()

	assertSome(t, 5, linq.Skip(linq.From(1, 2, 3, 4, 5), 3).FastElementAt(1))
	assertNo(t, linq.Skip(linq.From(1, 2, 3, 4, 5), 3).FastElementAt(3))
}
