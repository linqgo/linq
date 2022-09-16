package linq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertRingEmptyFull[T any](t *testing.T, r *ring[T], empty, full bool) bool {
	t.Helper()
	return assert.Equal(t, empty, r.Empty()) &&
		assert.Equal(t, full, r.Full())
}

func assertRingPush[T any](t *testing.T, r *ring[T], x T, full bool) bool {
	t.Helper()
	r.Push(x)
	return assertRingEmptyFull(t, r, false, full)
}

func assertRingPushPanics[T any](t *testing.T, r *ring[T]) bool {
	t.Helper()
	var x T
	return assert.Panics(t, func() { r.Push(x) })
}

func assertRingPop[T any](t *testing.T, r *ring[T], i int, empty bool) bool {
	t.Helper()
	return assert.Equal(t, i, r.Pop()) &&
		assertRingEmptyFull(t, r, empty, false)
}

func assertRingPopPanics[T any](t *testing.T, r *ring[T]) bool {
	t.Helper()
	return assert.Panics(t, func() { r.Pop() })
}

func TestRing(t *testing.T) {
	t.Parallel()

	r := newRing[int](3)
	assert.True(t, r.Empty())
	assert.False(t, r.Full())

	assertRingPush(t, &r, 1, false)
	assertRingPush(t, &r, 2, false)
	assertRingPush(t, &r, 3, true)

	assertRingPushPanics(t, &r)

	assertRingPop(t, &r, 1, false)
	assertRingPop(t, &r, 2, false)

	assertRingPush(t, &r, 4, false)
	assertRingPush(t, &r, 5, true)

	assertRingPushPanics(t, &r)

	assertRingPop(t, &r, 3, false)
	assertRingPop(t, &r, 4, false)
	assertRingPop(t, &r, 5, true)

	assertRingPopPanics(t, &r)
}

func TestRingEnumerator(t *testing.T) {
	t.Parallel()

	r := newRing[int](3)

	r.Push(1)
	r.Push(2)

	r.Pop()
	r.Pop()

	x, ok := r.Enumerator()()
	assert.False(t, ok, x)
}

func TestRingEnumeratorPartial(t *testing.T) {
	t.Parallel()

	r := newRing[int](3)

	r.Push(1)
	r.Push(2)

	r.Pop()

	x, ok := r.Enumerator()()
	assert.True(t, ok)
	assert.Equal(t, 2, x)
}
