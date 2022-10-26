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

func TestFromMap(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t, []linq.KV[int, int]{}, linq.FromMap(map[int]int{}))
	assertQueryElementsMatch(t,
		[]linq.KV[int, int]{{1, 1}, {2, 3}},
		linq.FromMap(map[int]int{1: 1, 2: 3}),
	)

	assertOneShot(t, false, linq.FromMap(map[int]int{1: 2, 2: 3}))
}

func TestMustToMapKV(t *testing.T) {
	t.Parallel()

	assert.Equal(t,
		map[int]int{1: 10, 2: 20},
		linq.MustToMapKV(linq.From(linq.NewKV(1, 10), linq.NewKV(2, 20))),
	)

	assert.Panics(t, func() {
		linq.MustToMapKV(linq.From(linq.NewKV(1, 10), linq.NewKV(1, 20)))
	})
}

func TestMustToMap(t *testing.T) {
	t.Parallel()

	assert.Equal(t,
		map[int]int{1: 1, 2: 4, 3: 9, 4: 16, 5: 25},
		linq.MustToMap(
			linq.From(1, 2, 3, 4, 5),
			func(i int) linq.KV[int, int] { return linq.NewKV(i, i*i) },
		),
	)

	assert.Panics(t, func() {
		linq.MustToMap(
			linq.From(1, 2, 3, 4, 5),
			func(i int) linq.KV[int, int] { return linq.NewKV(i%2, i) },
		)
	})
}

func TestSelectKeys(t *testing.T) {
	t.Parallel()

	peeps := linq.FromMap(map[string]int{"John": 42, "Sanjiv": 22, "Andrea": 35})

	assertQueryElementsMatch(t,
		[]string{"John", "Sanjiv", "Andrea"},
		linq.SelectKeys(peeps),
	)
}

func TestSelectValues(t *testing.T) {
	t.Parallel()

	peeps := linq.FromMap(map[string]int{"John": 42, "Sanjiv": 22, "Andrea": 35})

	assertQueryElementsMatch(t, []int{42, 22, 35}, linq.SelectValues(peeps))
}
