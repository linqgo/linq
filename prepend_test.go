package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestPrepend(t *testing.T) {
	t.Parallel()

	q := linq.Prepend(6, 7, 8)(linq.From(1, 2, 3, 4, 5))
	assertQueryEqual(t, []int{6, 7, 8, 1, 2, 3, 4, 5}, q)

	assertOneShot(t, false, q)
	assertOneShot(t, true, oneshot.Reverse())
}
