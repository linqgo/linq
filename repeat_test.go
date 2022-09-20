package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/linqgo/linq"
)

func TestRepeat(t *testing.T) {
	t.Parallel()

	assert.True(t, linq.Repeat(0, 0).Empty())
	assert.Equal(t, 10, linq.Repeat(0, 10).Count())
	assertQueryEqual(t, []int{1, 1, 1, 1, 1}, linq.Repeat(1, 5))

	assertOneShot(t, false, linq.Repeat(1, 5))

	assertFastCountEqual(t, 5, linq.Repeat(1, 5))
}

func TestRepeatForever(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 1000, linq.RepeatForever(42).Take(1000).Count())

	assertOneShot(t, false, linq.RepeatForever(42))

	assertNoFastCount(t, linq.RepeatForever(42))
}
