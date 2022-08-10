package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestFlatten(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t,
		[]int{1, 2, 3, 4},
		linq.Flatten(linq.From(linq.From(1, 2), linq.From(3, 4))),
	)
}

func TestFlattenSlices(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t,
		[]int{1, 2, 3, 4},
		linq.FlattenSlices(linq.From([]int{1, 2}, []int{3, 4})),
	)
}
