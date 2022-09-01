package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestConcat(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t,
		[]int{1, 2, 3, 4, 5},
		linq.Concat(linq.From(1, 2), linq.From(3), linq.From(4, 5)),
	)
	assertQueryEqual(t,
		[]int{1, 2, 4, 5},
		linq.Concat(linq.From(1, 2), linq.None[int](), linq.From(4, 5)),
	)

	assertOneShot(t, false, linq.Concat(linq.From(1, 2), linq.From(3)))
	assertOneShot(t, true, linq.Concat(oneshot, linq.From(1, 2)))
	assertOneShot(t, true, linq.Concat(linq.From(1, 2), oneshot))
}
