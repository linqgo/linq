package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestExcept(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t,
		[]int{1, 3, 5},
		linq.Except(linq.From(1, 2, 3, 4, 5), linq.From(2, 4)),
	)
	assertQueryEqual(t,
		[]int{},
		linq.Except(linq.From(1, 2, 3), linq.From(1, 2, 3)),
	)
	assertQueryEqual(t,
		[]int{},
		linq.Except(linq.From(1, 2, 3), linq.From(1, 2, 3, 4, 5)),
	)
	assertQueryEqual(t,
		[]int{1, 2, 3},
		linq.Except(linq.From(1, 2, 3), linq.From(4, 5)),
	)
	assertQueryEqual(t,
		[]int{1, 2},
		linq.Except(linq.From(1, 2, 3), linq.From(3, 4, 5)),
	)
}
