package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marcelocantos/linq"
)

func TestRepeat(t *testing.T) {
	t.Parallel()

	assert.True(t, linq.Repeat(0, 0).Empty())
	assert.Equal(t, 10, linq.Repeat(0, 10).Count())
	assert.Equal(t, []int{1, 1, 1, 1, 1}, linq.Repeat(1, 5).ToSlice())
}
