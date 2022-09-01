package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestWhere(t *testing.T) {
	t.Parallel()

	q := linq.Iota2(1, 6).Where(func(t int) bool { return t%2 == 1 })
	assertQueryEqual(t, []int{1, 3, 5}, q)

	assertOneShot(t, false, q.Where(linq.False[int]))
	assertOneShot(t, true, oneshot.Where(linq.False[int]))
}
