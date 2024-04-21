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

	"github.com/linqgo/linq/v2"
	"github.com/linqgo/linq/v2/internal/num"
)

func maybe[T any](t T, ok bool) func() (T, bool) {
	return func() (T, bool) { return t, ok }
}

func assertNo[T any](t *testing.T, m func() (T, bool)) bool {
	t.Helper()

	v, valid := m()
	return assert.False(t, valid, v)
}

func assertSome[T any](t *testing.T, expected T, m func() (T, bool)) bool {
	t.Helper()

	v, valid := m()
	return assert.True(t, valid) && assert.Equal(t, expected, v)
}

func assertSomeInEpsilon[T any](t *testing.T, expected T, m func() (T, bool), ε float64) bool {
	t.Helper()

	v, valid := m()
	return assert.True(t, valid) && assert.InEpsilon(t, expected, v, ε)
}

func assertQueryEqual[T any](t *testing.T, expected []T, q linq.Query[T]) bool {
	t.Helper()

	s := q.ToSlice()
	if len(s) == 0 && len(expected) == 0 {
		return true
	}
	return assert.Equal(t, expected, s)
}

func assertQueryInEpsilon[R num.RealNumber](t *testing.T, expected []R, q linq.Query[R], ε R) bool {
	t.Helper()

	s := q.ToSlice()
	if len(s) != len(expected) {
		return assert.Failf(t, "query has unexpected length", "%v", s)
	}
	if len(s) == 0 && len(expected) == 0 {
		return true
	}
	return assert.InEpsilonSlice(t, expected, s, float64(ε))
}
