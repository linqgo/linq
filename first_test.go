package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/linqgo/linq"
)

func TestFirstComp(t *testing.T) {
	t.Parallel()

	q := linq.From(2, 8, 5, 1)
	assertResultEqual(t, 8, maybe(q.FirstComp(linq.Greater[int])))
	assertNoResult(t, maybe(linq.None[int]().FirstComp(linq.Greater[int])))
}

func TestFirstCompElse(t *testing.T) {
	t.Parallel()

	q := linq.From(2, 8, 5, 1)
	assert.Equal(t, 8, q.FirstCompElse(linq.Greater[int], 0))
	assert.Equal(t, 0, linq.None[int]().FirstCompElse(linq.Greater[int], 0))
}

func TestLastComp(t *testing.T) {
	t.Parallel()

	q := linq.From(2, 8, 5, 1)
	assertResultEqual(t, 8, maybe(q.LastComp(linq.Less[int])))
	assertNoResult(t, maybe(linq.None[int]().LastComp(linq.Less[int])))
}

func TestLastCompElse(t *testing.T) {
	t.Parallel()

	q := linq.From(2, 8, 5, 1)
	assert.Equal(t, 8, q.LastCompElse(linq.Less[int], 0))
	assert.Equal(t, 0, linq.None[int]().LastCompElse(linq.Less[int], 0))
}

func TestMustFirstComp(t *testing.T) {
	t.Parallel()

	q := linq.From(2, 8, 5, 1)
	assert.Equal(t, 8, q.MustFirstComp(linq.Greater[int]))
	assert.Panics(t, func() { linq.None[int]().MustFirstComp(linq.Greater[int]) })
}

func TestMustLastComp(t *testing.T) {
	t.Parallel()

	q := linq.From(2, 8, 5, 1)
	assert.Equal(t, 8, q.MustLastComp(linq.Less[int]))
	assert.Panics(t, func() { linq.None[int]().MustLastComp(linq.Less[int]) })
}
