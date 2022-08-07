package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marcelocantos/linq"
)

func TestReverse(t *testing.T) {
	t.Parallel()

	assert.Equal(t,
		[]int{5, 4, 3, 2, 1},
		linq.Iota2(1, 6).Reverse().ToSlice(),
	)
}
