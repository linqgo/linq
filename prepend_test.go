package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestPrepend(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t,
		[]int{1, 2, 3, 4, 5, 6, 7},
		linq.Prepend(6)(linq.From(1, 2, 3, 4, 5)).Prepend(7))
}
