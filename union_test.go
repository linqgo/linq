package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestUnion(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t,
		[]int{1, 2, 3, 4, 5},
		linq.Union(linq.From(1, 2, 3, 4, 5), linq.From(2, 4)),
	)
	assertQueryEqual(t,
		[]int{1, 2, 3},
		linq.Union(linq.From(1, 2, 3), linq.From(1, 2, 3)),
	)
	assertQueryEqual(t,
		[]int{1, 2, 3, 4, 5},
		linq.Union(linq.From(1, 2, 3), linq.From(1, 2, 3, 4, 5)),
	)
	assertQueryEqual(t,
		[]int{1, 2, 3, 4, 5},
		linq.Union(linq.From(1, 2, 3), linq.From(4, 5)),
	)
	assertQueryEqual(t,
		[]int{1, 2, 3, 4, 5},
		linq.Union(linq.From(1, 2, 3), linq.From(3, 4, 5)),
	)
}
