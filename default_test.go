package linq_test

import (
	"testing"

	"github.com/linqgo/linq"
)

func TestDefaultIfEmpty(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t, []int{42}, linq.From[int]().DefaultIfEmpty(42))
	assertQueryEqual(t, []int{1, 2, 3}, linq.From(1, 2, 3).DefaultIfEmpty(42))
	assertQueryEqual(t, []int{42}, oneshot().DefaultIfEmpty(56))
	assertQueryEqual(t, []int{56}, oneshot().Skip(2).DefaultIfEmpty(56))

	assertOneShot(t, false, linq.From[int]().DefaultIfEmpty(42))
	assertOneShot(t, false, linq.From(1, 2, 3).DefaultIfEmpty(42))
	assertOneShot(t, true, oneshot().DefaultIfEmpty(42))

	assertFastCountEqual(t, 1, linq.From[int]().DefaultIfEmpty(42))
	assertFastCountEqual(t, 3, linq.From(1, 2, 3).DefaultIfEmpty(42))
	assertNoFastCount(t, oneshot().DefaultIfEmpty(42))
}
