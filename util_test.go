package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marcelocantos/linq"
)

func TestDeref(t *testing.T) {
	t.Parallel()

	assert.Equal(t, struct{ X int }{42}, linq.Deref(&struct{ X int }{42}))
}

func TestFalse(t *testing.T) {
	t.Parallel()

	assert.False(t, linq.False(56))
}

func TestLess(t *testing.T) {
	t.Parallel()

	assert.False(t, linq.Less(10, 5))
	assert.False(t, linq.Less(5, 5))
	assert.True(t, linq.Less(5, 10))
}

func TestGreater(t *testing.T) {
	t.Parallel()

	assert.False(t, linq.Greater("10", "5"))
	assert.False(t, linq.Greater("5", "5"))
	assert.True(t, linq.Greater("5", "10"))
}

func TestLongerMap(t *testing.T) {
	t.Parallel()

	assert.False(t, linq.LongerMap(map[int]int{9: 1}, map[int]int{9: 3, 8: 4}))
	assert.False(t, linq.LongerMap(map[int]int{9: 1, 8: 2}, map[int]int{9: 3, 8: 4}))
	assert.True(t, linq.LongerMap(map[int]int{9: 1, 8: 2, 7: 3}, map[int]int{9: 3, 8: 4}))
}

func TestLongerSlice(t *testing.T) {
	t.Parallel()

	assert.False(t, linq.LongerSlice([]int{1}, []int{3, 4}))
	assert.False(t, linq.LongerSlice([]int{1, 2}, []int{3, 4}))
	assert.True(t, linq.LongerSlice([]int{1, 2, 3}, []int{3, 4}))
}

func TestPointer(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 56, *linq.Pointer(56))
}

func TestShorterMap(t *testing.T) {
	t.Parallel()

	assert.True(t, linq.ShorterMap(map[int]int{9: 1}, map[int]int{9: 3, 8: 4}))
	assert.False(t, linq.ShorterMap(map[int]int{9: 1, 8: 2}, map[int]int{9: 3, 8: 4}))
	assert.False(t, linq.ShorterMap(map[int]int{9: 1, 8: 2, 7: 3}, map[int]int{9: 3, 8: 4}))
}

func TestShorterSlice(t *testing.T) {
	t.Parallel()

	assert.True(t, linq.ShorterSlice([]int{1}, []int{3, 4}))
	assert.False(t, linq.ShorterSlice([]int{1, 2}, []int{3, 4}))
	assert.False(t, linq.ShorterSlice([]int{1, 2, 3}, []int{3, 4}))
}

func TestTrue(t *testing.T) {
	t.Parallel()

	assert.True(t, linq.True(42))
}

func TestZero(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "", linq.Zero[int, string](42))
}
