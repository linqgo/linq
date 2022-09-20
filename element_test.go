package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/linqgo/linq"
)

func TestElementAt(t *testing.T) {
	t.Parallel()

	assertNoResult(t, maybe(linq.From[int]().ElementAt(0)))
	assertNoResult(t, maybe(linq.Iota2(1, 6).ElementAt(42)))
	assertResultEqual(t, 3, maybe(linq.Iota2(1, 6).ElementAt(2)))
	assert.Equal(t, 2, linq.Iota2(1, 6).ElementAtElse(1, 56))
	assert.Equal(t, 56, linq.Iota2(1, 6).ElementAtElse(42, 56))
	assert.Equal(t, 3, linq.Iota2(1, 6).MustElementAt(2))
	assert.Panics(t, func() { linq.Iota2(1, 6).MustElementAt(5) })
}

func TestFirst(t *testing.T) {
	t.Parallel()

	assertNoResult(t, maybe(linq.Iota1(0).First()))
	assert.Equal(t, 0, linq.Iota1(10).FirstElse(56))
	assert.Equal(t, 56, linq.Iota1(0).FirstElse(56))
	assertResultEqual(t, 1, maybe(linq.Iota2(1, 6).First()))
	assert.Equal(t, 1, linq.Iota2(1, 6).MustFirst())
}

func TestLast(t *testing.T) {
	t.Parallel()

	assertNoResult(t, maybe(linq.Iota1(0).Last()))
	assert.Equal(t, 9, linq.Iota1(10).LastElse(56))
	assert.Equal(t, 56, linq.Iota1(0).LastElse(56))
	assertResultEqual(t, 5, maybe(linq.Iota2(1, 6).Last()))
	assert.Equal(t, 5, linq.Iota2(1, 6).MustLast())
}
