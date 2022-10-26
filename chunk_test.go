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

	assertSome(t, 3, linq.ChunkSlices(data, 2).FastCount())
	assertSome(t, 2, linq.ChunkSlices(data.Skip(1), 2).FastCount())
	assertNo(t, linq.ChunkSlices(slowcount, 2).FastCount())
}

func TestChunkElementAt(t *testing.T) {
	t.Parallel()

	data := linq.Chunk(linq.Iota1(100), 3)
	assertSomeQuery(t, []int{0, 1, 2}, data.FastElementAt(0))
	assertSomeQuery(t, []int{30, 31, 32}, data.FastElementAt(10))
	assertSomeQuery(t, []int{99}, data.FastElementAt(33))
	assertNo(t, data.FastElementAt(34))
	assertNo(t, data.FastElementAt(-1))
}
