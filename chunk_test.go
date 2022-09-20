package linq_test

import (
	"testing"

	"github.com/linqgo/linq"
)

func TestChunk(t *testing.T) {
	t.Parallel()

	data := linq.From(1, 2, 3, 4, 5)

	assertQueryEqual(t,
		[][]int{{1, 2}, {3, 4}, {5}},
		linq.ChunkSlices(data, 2),
	)

	assertOneShot(t, false, linq.ChunkSlices(data, 2))
	assertOneShot(t, true, linq.ChunkSlices(oneshot(), 2))

	assertFastCountEqual(t, 3, linq.ChunkSlices(data, 2))
	assertFastCountEqual(t, 2, linq.ChunkSlices(data.Skip(1), 2))
	assertNoFastCount(t, linq.ChunkSlices(slowcount, 2))
}
