package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"

	"github.com/stretchr/testify/assert"
)

func TestMath(t *testing.T) {
	t.Parallel()

	data := linq.Select(linq.Iota2(1, 11), func(i int) float64 { return float64(i) })
	assert.EqualValues(t, 5.5, linq.Average(data))
	assert.EqualValues(t, 10, linq.Max(data))
	assert.EqualValues(t, 1, linq.Min(data))
	assert.EqualValues(t, 55, linq.Sum(data))

	data = linq.Reverse(data)
	assert.EqualValues(t, 5.5, linq.Average(data))
	assert.EqualValues(t, 10, linq.Max(data))
	assert.EqualValues(t, 1, linq.Min(data))
	assert.EqualValues(t, 55, linq.Sum(data))
}
