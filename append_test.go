package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestAppend(t *testing.T) {
	t.Parallel()

	q := linq.From(1, 2, 3, 4, 5).Append(6).Append(7)
	assertQueryEqual(t, []int{1, 2, 3, 4, 5, 6, 7}, q)

	assertOneShot(t, false, q)
	assertOneShot(t, true, oneshot.Append(6).Append(7))
}
