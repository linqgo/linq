package linq_test

import (
	"testing"

	"github.com/linqgo/linq"
)

func TestTake(t *testing.T) {
	t.Parallel()

	data := linq.From(1, 2, 3, 4, 5)

	assertQueryEqual(t, []int{}, data.Take(0))
	assertQueryEqual(t, []int{1, 2}, data.Take(2))
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, data.Take(10))

	assertOneShot(t, false, data.Take(999))
	assertOneShot(t, true, oneshot().Take(999))

	assertSome(t, 5, data.Take(999).FastCount())
	assertNo(t, oneshot().Take(999).FastCount())
}

func TestTakeLast(t *testing.T) {
	t.Parallel()

	data := linq.From(1, 2, 3, 4, 5)

	assertQueryEqual(t, []int{}, data.TakeLast(0))
	assertQueryEqual(t, []int{4, 5}, data.TakeLast(2))
	assertQueryEqual(t, []int{3, 4, 5}, data.TakeLast(3))
	assertQueryEqual(t, []int{2, 3, 4, 5}, data.TakeLast(4))
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, data.TakeLast(10))

	assertOneShot(t, false, data.TakeLast(999))
	assertOneShot(t, true, oneshot().TakeLast(999))

	assertSome(t, 5, data.TakeLast(999).FastCount())
	assertNo(t, oneshot().TakeLast(999).FastCount())
}

func TestTakeLastOneShot(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t, []int{4, 5}, oneshotN(1, 2, 3, 4, 5).TakeLast(2))
}

func TestTakeWhile(t *testing.T) {
	t.Parallel()

	data := linq.From(1, 2, 3, 4, 5)

	assertQueryEqual(t, []int{}, data.TakeWhile(func(i int) bool { return i < 0 }))
	assertQueryEqual(t, []int{1, 2}, data.TakeWhile(func(i int) bool { return i < 3 }))
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, data.TakeWhile(func(i int) bool { return i < 10 }))

	assertOneShot(t, false, data.TakeWhile(linq.True[int]))
	assertOneShot(t, true, oneshot().TakeWhile(linq.True[int]))

	assertSome(t, 0, linq.None[int]().TakeWhile(linq.True[int]).FastCount())
	assertNo(t, data.TakeWhile(linq.True[int]).FastCount())
	assertNo(t, oneshot().TakeLast(999).FastCount())
}

func TestTakeElementAt(t *testing.T) {
	t.Parallel()

	assertSome(t, 2, linq.Take(linq.From(1, 2, 3, 4, 5), 3).FastElementAt(1))
	assertNo(t, linq.Take(linq.From(1, 2, 3, 4, 5), 3).FastElementAt(3))
}
