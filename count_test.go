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

package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/linqgo/linq"
)

func TestCount(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 0, linq.From[int]().Count())
	assert.Equal(t, 5, linq.From(1, 2, 3, 4, 5).Count())
	assert.Equal(t, 1, linq.From(42).Count())
	assert.Equal(t, 7,
		linq.Concat(
			linq.Concat(
				linq.From(1, 2, 3),
				linq.From(4),
			),
			linq.From(5, 6, 7),
		).Count(),
	)
}

func TestCountLimit(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 0, linq.From[int]().CountLimit(10))
	assert.Equal(t, 5, linq.From(1, 2, 3, 4, 5).CountLimit(3))
	assert.Equal(t, 5, linq.From(1, 2, 3, 4, 5).CountLimit(10))

	c := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		c <- i
	}
	assert.Equal(t, 3, linq.FromChannel(c).CountLimit(3))
}

type countTrue[T any] struct {
	t      T
	isTrue bool
}

func ct[T any](t T, isTrue bool) countTrue[T] {
	return countTrue[T]{t: t, isTrue: isTrue}
}

func TestCountLimitTrue(t *testing.T) {
	t.Parallel()

	assert.Equal(t, ct(0, true), ct(linq.From[int]().CountLimitTrue(10)))
	assert.Equal(t, ct(5, true), ct(linq.From(1, 2, 3, 4, 5).CountLimitTrue(3)))
	assert.Equal(t, ct(5, true), ct(linq.From(1, 2, 3, 4, 5).CountLimitTrue(10)))

	assert.Equal(t, ct(3, false), ct(chanof(1, 2, 3, 4, 5).CountLimitTrue(3)))
	assert.Equal(t, ct(3, true), ct(chanof(1, 2, 3).CountLimitTrue(5)))
}

func TestFastCount(t *testing.T) {
	t.Parallel()

	assertSome(t, 5, linq.Iota1(5).FastCount())
	assertNo(t, linq.Iota[int]().FastCount())
}
