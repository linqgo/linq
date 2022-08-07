package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestAppend(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t,
		[]int{1, 2, 3, 4, 5, 6, 7},
		linq.From(1, 2, 3, 4, 5).Append(6).Append(7))
}
