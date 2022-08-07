package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestChunk(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t,
		[][]int{{1, 2}, {3, 4}, {5}},
		linq.ChunkSlices(linq.From(1, 2, 3, 4, 5), 2),
	)
}
