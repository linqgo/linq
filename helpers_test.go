package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/linqgo/linq"
)

func assertNo[T any](t *testing.T, m linq.Maybe[T]) bool {
	t.Helper()

	v, valid := m.Get()
	return assert.False(t, valid, v)
}

func assertSome[T any](t *testing.T, expected T, m linq.Maybe[T]) bool {
	t.Helper()

	v, valid := m.Get()
	return assert.True(t, valid) && assert.Equal(t, expected, v)
}

func assertSomeInEpsilon[T any](t *testing.T, expected T, m linq.Maybe[T], ε float64) bool {
	t.Helper()

	v, valid := m.Get()
	return assert.True(t, valid) && assert.InEpsilon(t, expected, v, ε)
}

func assertSomeQuery[T any](t *testing.T, expected []T, m linq.Maybe[linq.Query[T]]) bool {
	t.Helper()

	v, valid := m.Get()
	return assert.True(t, valid) && assertQueryEqual(t, expected, v)
}

func assertQueryElementsMatch[T any](t *testing.T, expected []T, q linq.Query[T]) bool {
	t.Helper()

	s := q.ToSlice()
	if len(s) == 0 && len(expected) == 0 {
		return true
	}
	return assert.Equal(t, q.Empty(), q.OneShot()) &&
		assert.ElementsMatch(t, expected, s) &&
		assertExhaustedEnumeratorBehavesWell(t, q)
}

func assertQueryEqual[T any](t *testing.T, expected []T, q linq.Query[T]) bool {
	t.Helper()

	s := q.ToSlice()
	if len(s) == 0 && len(expected) == 0 {
		return true
	}
	return assert.Equal(t, expected, s) &&
		assertExhaustedEnumeratorBehavesWell(t, q)
}

func assertExhaustedEnumeratorBehavesWell[T any](t *testing.T, q linq.Query[T]) bool {
	t.Helper()

	next := q.Enumerator()
	linq.Drain(next)
	var m linq.Maybe[T]
	return assert.NotPanics(t, func() { m = next() }) && assertNo(t, m)
}

func oneshot() linq.Query[int] {
	return oneshotN(42)
}

func oneshotN[T any](data ...T) linq.Query[T] {
	c := make(chan T, len(data))
	for _, e := range data {
		c <- e
	}
	close(c)
	return linq.FromChannel(c)
}

var slowcount = oneshot()

func assertOneShot[T any](t *testing.T, oneshot bool, q linq.Query[T]) bool {
	t.Helper()
	return assert.Equal(t, oneshot, q.OneShot())
}
