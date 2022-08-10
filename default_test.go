package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
	"github.com/stretchr/testify/assert"
)

func TestDefaultIfEmpty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, []int{42}, linq.From[int]().DefaultIfEmpty(42))
}
