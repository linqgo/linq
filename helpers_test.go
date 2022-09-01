package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marcelocantos/linq"
)

type result[T any] struct {
	t  T
	ok bool
}

func maybe[T any](t T, ok bool) result[T] {
	return result[T]{t: t, ok: ok}
}

func assertNoResult[T any](t *testing.T, r result[T]) bool {
	t.Helper()

	return assert.False(t, r.ok, r.t)
}

func assertResultEqual[T any](t *testing.T, expected T, r result[T]) bool {
	t.Helper()

	return assert.True(t, r.ok) && assert.Equal(t, expected, r.t)
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
	return assert.Equal(t, q.Empty(), q.OneShot()) &&
		assert.Equal(t, expected, s) &&
		assertExhaustedEnumeratorBehavesWell(t, q)
}

func assertExhaustedEnumeratorBehavesWell[T any](t *testing.T, q linq.Query[T]) bool {
	t.Helper()

	next := q.Enumerator()
	for _, ok := next(); ok; {
		_, ok = next()
	}
	var ok bool
	return assert.NotPanics(t, func() { _, ok = next() }) &&
		assert.False(t, ok)
}

var oneshot = func() linq.Query[int] {
	c := make(chan int, 1)
	c <- 42
	close(c)
	return linq.FromChannel(c)
}()

func assertOneShot[T any](t *testing.T, oneshot bool, q linq.Query[T]) bool {
	return assert.Equal(t, oneshot, q.OneShot())
}
