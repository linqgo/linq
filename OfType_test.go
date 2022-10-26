package linq_test

import (
	"testing"

	"github.com/linqgo/linq"
)

// OfType returns a Query that contains all the elements of q that have type U.
func TestOfType(t *testing.T) {
	t.Parallel()

	data := linq.From[any](1, "hello", 2, 3, "goodbye")

	assertQueryEqual(t, []int{1, 2, 3}, linq.OfType[int](data))
	assertQueryEqual(t, []string{"hello", "goodbye"}, linq.OfType[string](data))

	assertOneShot(t, false, linq.OfType[int](data))
	assertOneShot(t, true, linq.OfType[int](oneshot()))

	assertSome(t, 0, linq.OfType[int](linq.None[any]()).FastCount())
	assertNo(t, linq.OfType[int](data).FastCount())
	assertNo(t, linq.OfType[int](oneshot()).FastCount())
}
