// Copyright 2022-2024 Marcelo Cantos
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

package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/linqgo/linq/v2"
)

func maybe[T any](t T, ok bool) func() (T, bool) {
	return func() (T, bool) { return t, ok }
}

func must[T any](t T, ok bool) T {
	if !ok {
		panic("no value")
	}
	return t
}

func assertSome[T any](t *testing.T, expected T, next func() (T, bool)) bool {
	t.Helper()
	v, ok := next()
	return assert.True(t, ok) && assert.Equal(t, expected, v)
}

func assertNo[T any](t *testing.T, next func() (T, bool)) bool {
	t.Helper()
	v, ok := next()
	return assert.False(t, ok, v)
}

func assertHaveQuery[T any](t *testing.T, expected []T, m func() (linq.Query[T], bool)) bool {
	t.Helper()

	v, valid := m()
	return assert.True(t, valid) && assertQueryEqual(t, expected, v)
}

func assertQueryElementsMatch[T any](t *testing.T, expected []T, q linq.Query[T]) bool {
	t.Helper()

	s := q.ToSlice()
	if len(s) == 0 && len(expected) == 0 {
		return true
	}
	return assert.Equal(t, q.Empty(), q.OneShot()) &&
		assert.ElementsMatch(t, expected, s)
}

func assertQueryEqual[T any](t *testing.T, expected []T, q linq.Query[T]) bool {
	t.Helper()

	s := q.ToSlice()
	if len(s) == 0 && len(expected) == 0 {
		return true
	}
	return assert.Equal(t, expected, s)
}

func assertAll[R any](
	t *testing.T,
	test func(r R) bool,
	r ...R,
) bool {
	t.Helper()
	pass := true
	for i, r := range r {
		pass = assert.True(t, test(r), i) && pass
	}
	return pass
}

func oneshot() linq.Query[int] {
	return chanof(42)
}

var slowcount = oneshot()

func chanof[T any](t ...T) linq.Query[T] {
	c := make(chan T, len(t))
	for _, t := range t {
		c <- t
	}
	close(c)
	return linq.FromChannel(c)
}

func assertOneShot[T any](t *testing.T, oneshot bool, q linq.Query[T]) bool {
	t.Helper()
	return assert.Equal(t, oneshot, q.OneShot())
}
