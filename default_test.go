package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marcelocantos/linq"
)

func TestDefaultIfEmpty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, []int{42}, linq.From[int]().DefaultIfEmpty(42).ToSlice())
}
