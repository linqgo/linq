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

	// assertOneShot(t, false, data.Take(999))
	// assertOneShot(t, true, oneshot().Take(999))

	// assertFastCountEqual(t, 5, data.Take(999))
	// assertNoFastCount(t, oneshot().Take(999))
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

	assertFastCountEqual(t, 5, data.TakeLast(999))
	assertNoFastCount(t, oneshot().TakeLast(999))
}

func TestTakeWhile(t *testing.T) {
	t.Parallel()

	data := linq.From(1, 2, 3, 4, 5)

	assertQueryEqual(t, []int{}, data.TakeWhile(func(i int) bool { return i < 0 }))
	assertQueryEqual(t, []int{1, 2}, data.TakeWhile(func(i int) bool { return i < 3 }))
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, data.TakeWhile(func(i int) bool { return i < 10 }))

	assertOneShot(t, false, data.TakeWhile(linq.True[int]))
	assertOneShot(t, true, oneshot().TakeWhile(linq.True[int]))

	assertFastCountEqual(t, 0, linq.None[int]().TakeWhile(linq.True[int]))
	assertNoFastCount(t, data.TakeWhile(linq.True[int]))
	assertNoFastCount(t, oneshot().TakeLast(999))
}
