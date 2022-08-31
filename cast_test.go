package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

// OfType returns a Query that contains all the elements of q that have type U.
func TestOfType(t *testing.T) {
	t.Parallel()

	data := linq.From[any](1, "hello", 2, 3, "goodbye")

	assertQueryEqual(t, []int{1, 2, 3}, linq.OfType[int](data))
	assertQueryEqual(t, []string{"hello", "goodbye"}, linq.OfType[string](data))
}
