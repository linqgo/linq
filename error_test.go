package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/linqgo/linq"
)

func TestErrorError(t *testing.T) {
	t.Parallel()

	assert.EqualError(t, linq.EmptySourceError, "empty source")
}
