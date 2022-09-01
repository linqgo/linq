package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestDefaultIfEmpty(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t, []int{42}, linq.From[int]().DefaultIfEmpty(42))
	assertQueryEqual(t, []int{1, 2, 3}, linq.From(1, 2, 3).DefaultIfEmpty(42))

	assertOneShot(t, false, linq.From[int]().DefaultIfEmpty(42))
	assertOneShot(t, false, linq.From(1, 2, 3).DefaultIfEmpty(42))
	assertOneShot(t, true, oneshot.DefaultIfEmpty(42))
}
