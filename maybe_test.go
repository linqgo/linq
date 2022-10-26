package linq_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/linqgo/linq"
)

func TestMaybeMustPanics(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() { linq.No[int]().Must() })
}

func TestMaybeElse(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 42, linq.Some(42).Else(56))
	assert.Equal(t, 56, linq.No[int]().Else(56))
}

func TestMaybeElseNaN(t *testing.T) {
	t.Parallel()

	assert.True(t, math.IsNaN(linq.ElseNaN(linq.No[float64]())))
}
