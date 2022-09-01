package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marcelocantos/linq"
)

func TestTrue(t *testing.T) {
	t.Parallel()

	assert.True(t, linq.True(42))
}

func TestFalse(t *testing.T) {
	t.Parallel()

	assert.False(t, linq.False(56))
}
