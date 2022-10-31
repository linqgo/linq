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
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/linqgo/linq"
)

func TestMaybeMustPanics(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() { linq.No[int]().Must() })
}

func TestMaybeElse(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 42, linq.Some(42).Else(56))
	assert.Equal(t, 56, linq.No[int]().Else(56))
}

func TestMaybeElseNaN(t *testing.T) {
	t.Parallel()

	assert.True(t, math.IsNaN(linq.ElseNaN(linq.No[float64]())))
}

func TestMaybeMap(t *testing.T) {
	t.Parallel()

	f := func(i int) linq.Maybe[int] { return linq.NewMaybe(i+14, i >= 0) }
	assertSome(t, 56, linq.MaybeFlatMap(linq.Some(42), f))
	assertNo(t, linq.MaybeFlatMap(linq.No[int](), f))
	assertNo(t, linq.MaybeFlatMap(linq.Some(-42), f))
}
