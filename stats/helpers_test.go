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

package stats_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/linqgo/linq"
	"github.com/linqgo/linq/internal/num"
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

func assertQueryEqual[T any](t *testing.T, expected []T, q linq.Query[T]) bool {
	t.Helper()

	s := q.ToSlice()
	if len(s) == 0 && len(expected) == 0 {
		return true
	}
	return assert.Equal(t, expected, s) &&
		assertExhaustedEnumeratorBehavesWell(t, q)
}

func assertQueryInEpsilon[R num.RealNumber](t *testing.T, expected []R, q linq.Query[R], ε R) bool {
	t.Helper()

	s := q.ToSlice()
	if len(s) == 0 && len(expected) == 0 {
		return true
	}
	return assert.InEpsilonSlice(t, expected, s, float64(ε)) &&
		assertExhaustedEnumeratorBehavesWell(t, q)
}

func assertExhaustedEnumeratorBehavesWell[T any](t *testing.T, q linq.Query[T]) bool {
	t.Helper()

	next := q.Enumerator()
	linq.Drain(next)
	var m linq.Maybe[T]
	return assert.NotPanics(t, func() { m = next() }) && assertNo(t, m)
}

func chanof[T any](t ...T) linq.Query[T] {
	c := make(chan T, len(t))
	for _, t := range t {
		c <- t
	}
	close(c)
	return linq.FromChannel(c)
}
