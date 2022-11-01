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

func TestMaybeFlatMap(t *testing.T) {
	t.Parallel()

	type mofunc func(i int) linq.Maybe[int]

	f := func(i int) linq.Maybe[int] { return linq.NewMaybe(i+14, i >= 0) }
	g := func(i int) linq.Maybe[int] { return linq.Some(-i) }

	assertSome(t, 56, linq.Some(42).FlatMap(f))
	assertNo(t, linq.No[int]().FlatMap(f))
	assertNo(t, linq.Some(-42).FlatMap(f))

	// Left-identity: (unit(x) >>= f) = f(x)
	for _, x := range []int{5, -5} {
		for _, unit := range []mofunc{f, g} {
			assert.Equal(t, unit(x), linq.Some(x).FlatMap(unit), x)
		}
	}

	mm := []linq.Maybe[int]{linq.No[int](), linq.Some(5), linq.Some(-5)}

	// Right-identity: (ma >>= unit) = ma
	for _, m := range mm {
		assert.Equal(t, m, m.FlatMap(linq.Some[int]), m)
	}

	// Associativity: (ma >>= λx → (f(x) >>= g)) = ((ma >>= f) >>= g)
	for _, m := range mm {
		for i, f := range []mofunc{f, g, linq.Some[int]} {
			for j, g := range []mofunc{f, g, linq.Some[int]} {
				assert.Equal(t,
					m.FlatMap(f).FlatMap(g),
					m.FlatMap(func(x int) linq.Maybe[int] { return f(x).FlatMap(g) }),
					m, i, j,
				)
			}
		}
	}
}
