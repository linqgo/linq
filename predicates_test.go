package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marcelocantos/linq"
)

func TestAll(t *testing.T) {
	t.Parallel()

	assert.True(t, linq.Iota1(5).All(func(t int) bool { return t < 10 }))
	assert.False(t, linq.Iota1(5).All(func(t int) bool { return t < 3 }))
	assert.True(t, linq.Iota1(0).All(func(t int) bool { return false }))
}

func TestAny(t *testing.T) {
	t.Parallel()

	assert.True(t, linq.Iota1(5).Any(func(t int) bool { return t < 3 }))
	assert.False(t, linq.Iota1(5).Any(func(t int) bool { return t > 10 }))
	assert.False(t, linq.Iota1(0).Any(func(t int) bool { return true }))
}

func TestEmpty(t *testing.T) {
	t.Parallel()

	assert.True(t, linq.None[int]().Empty())
	assert.False(t, linq.From(1, 2, 3).Empty())
}
