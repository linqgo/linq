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

func TestCount(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 0, linq.From[int]().Count())
	assert.Equal(t, 5, linq.From(1, 2, 3, 4, 5).Count())
	assert.Equal(t, 1, linq.From(42).Count())
	assert.Equal(t, 7,
		linq.From(1, 2, 3).Concat(linq.From(4)).Concat(linq.From(5, 6, 7)).Count(),
	)
}

func TestCountLimit(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 0, linq.From[int]().CountLimit(10))
	assert.Equal(t, 3, linq.From(1, 2, 3, 4, 5).CountLimit(3))
	assert.Equal(t, 5, linq.From(1, 2, 3, 4, 5).CountLimit(10))

	assert.Equal(t, 0, chanof(1, 2, 3, 4, 5).CountLimit(0))
	assert.Equal(t, 4, chanof(1, 2, 3, 4, 5).CountLimit(4))
	assert.Equal(t, 5, chanof(1, 2, 3, 4, 5).CountLimit(5))
	assert.Equal(t, 5, chanof(1, 2, 3, 4, 5).CountLimit(6))
	assert.Equal(t, 5, chanof(1, 2, 3, 4, 5).CountLimit(10))
}

type countTrue[T any] struct {
	t      T
	isTrue bool
}

func ct[T any](t T, isTrue bool) countTrue[T] {
	return countTrue[T]{t: t, isTrue: isTrue}
}

func TestFastCount(t *testing.T) {
	t.Parallel()

	assertSome(t, 5, linq.Iota1(5).FastCount)
	assertNo(t, linq.Iota[int]().FastCount)
}
