package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/linqgo/linq"
)

func TestContains(t *testing.T) {
	t.Parallel()

	assert.False(t, linq.Contains(linq.None[int](), 42))
	assert.False(t, linq.Contains(linq.From(1, 2, 3, 4, 5), 42))
	assert.True(t, linq.Contains(linq.From(1, 2, 3, 4, 5), 3))
}
