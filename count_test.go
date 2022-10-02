package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/linqgo/linq"
)

func TestCount(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 0, linq.From[int]().Count())
	assert.Equal(t, 5, linq.From(1, 2, 3, 4, 5).Count())
	assert.Equal(t, 1, linq.From(42).Count())
	assert.Equal(t, 7,
		linq.Concat(
			linq.Concat(
				linq.From(1, 2, 3),
				linq.From(4),
			),
			linq.From(5, 6, 7),
		).Count(),
	)
}

func TestCountLimit(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 0, linq.From[int]().CountLimit(10))
	assert.Equal(t, 5, linq.From(1, 2, 3, 4, 5).CountLimit(3))
	assert.Equal(t, 5, linq.From(1, 2, 3, 4, 5).CountLimit(10))

	c := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		c <- i
	}
	assert.Equal(t, 3, linq.FromChannel(c).CountLimit(3))
}

func TestCountLimitTrue(t *testing.T) {
	t.Parallel()

	assert.Equal(t, maybe(0, true), maybe(linq.From[int]().CountLimitTrue(10)))
	assert.Equal(t, maybe(5, true), maybe(linq.From(1, 2, 3, 4, 5).CountLimitTrue(3)))
	assert.Equal(t, maybe(5, true), maybe(linq.From(1, 2, 3, 4, 5).CountLimitTrue(10)))

	assert.Equal(t, maybe(3, false), maybe(chanof(1, 2, 3, 4, 5).CountLimitTrue(3)))
	assert.Equal(t, maybe(3, true), maybe(chanof(1, 2, 3).CountLimitTrue(5)))
}

func TestFastCount(t *testing.T) {
	t.Parallel()

	assertFastCountEqual(t, 5, linq.Iota1(5))
	assert.Equal(t, 5, linq.Iota1(5).MustFastCount())
	assertNoFastCount(t, linq.Iota[int]())
	assert.Panics(t, func() { linq.Iota[int]().MustFastCount() })
}
