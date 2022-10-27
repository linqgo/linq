// Copyright 2022 Marcelo Cantos
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ring_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/linqgo/linq/internal/ring"
)

func assertRingEmpty[T any](t *testing.T, r *ring.Ring[T], empty bool) bool {
	t.Helper()
	return assert.Equal(t, empty, r.Empty())
}

func assertRingPush[T any](t *testing.T, r *ring.Ring[T], x T) bool {
	t.Helper()
	r.Push(x)
	return assertRingEmpty(t, r, false)
}

func assertRingPop[T any](t *testing.T, r *ring.Ring[T], i int, empty bool) bool {
	t.Helper()
	return assert.Equal(t, i, r.Pop()) &&
		assertRingEmpty(t, r, empty)
}

func assertRingPopPanics[T any](t *testing.T, r *ring.Ring[T]) bool {
	t.Helper()
	return assert.Panics(t, func() { r.Pop() })
}

func TestRing(t *testing.T) {
	t.Parallel()

	r := ring.New[int](3)
	assert.True(t, r.Empty())

	assertRingPush(t, &r, 1)
	assertRingPush(t, &r, 2)
	assertRingPush(t, &r, 3)
	assertRingPush(t, &r, 4)

	assertRingPop(t, &r, 1, false)
	assertRingPop(t, &r, 2, false)

	assertRingPush(t, &r, 5)
	assertRingPush(t, &r, 6)
	assertRingPush(t, &r, 7)
	assertRingPush(t, &r, 8)

	assertRingPop(t, &r, 3, false)
	assertRingPop(t, &r, 4, false)
	assertRingPop(t, &r, 5, false)
	assertRingPop(t, &r, 6, false)
	assertRingPop(t, &r, 7, false)
	assertRingPop(t, &r, 8, true)

	assertRingPopPanics(t, &r)
}

func TestRingEnumerator(t *testing.T) {
	t.Parallel()

	r := ring.New[int](3)

	r.Push(1)
	r.Push(2)

	r.Pop()
	r.Pop()

	x, ok := r.Enumerator()()
	assert.False(t, ok, x)
}

func TestRingEnumeratorPartial(t *testing.T) {
	t.Parallel()

	r := ring.New[int](3)

	r.Push(1)
	r.Push(2)

	r.Pop()

	next := r.Enumerator()
	x, ok := next()
	assert.True(t, ok)
	assert.Equal(t, 2, x)
	_, ok = next()
	assert.False(t, ok)

	r.Push(3)
	r.Pop()
	r.Push(4)

	next = r.Enumerator()
	x, ok = next()
	assert.True(t, ok)
	assert.Equal(t, 3, x)
	x, ok = next()
	assert.True(t, ok)
	assert.Equal(t, 4, x)
	_, ok = next()
	assert.False(t, ok)
}
