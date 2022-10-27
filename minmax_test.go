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

	"github.com/linqgo/linq"
)

var (
	testNums  = linq.Select(linq.Iota2(1, 11), func(i int) float64 { return float64(i) })
	emptyNums = linq.None[float64]()
)

func TestMax(t *testing.T) {
	t.Parallel()

	for _, data := range []linq.Query[float64]{testNums, linq.Reverse(testNums)} {
		assertSome(t, 10.0, linq.Max(data))
		assertNo(t, linq.Max(emptyNums))
	}
}

func TestMaxBy(t *testing.T) {
	t.Parallel()

	type Person = linq.KV[string, int]

	peeps := linq.FromMap(map[string]int{"John": 42, "Sanjiv": 22, "Andrea": 35})
	noone := linq.FromMap(map[string]int{})

	name := func(kv Person) string { return kv.Key }
	age := func(kv Person) int { return kv.Value }

	assertSome(t, linq.NewKV("Sanjiv", 22), linq.MaxBy(peeps, name))
	assertNo(t, linq.MaxBy(noone, name))
	assertSome(t, linq.NewKV("John", 42), linq.MaxBy(peeps, age))
	assertNo(t, linq.MaxBy(noone, age))
}

func TestMin(t *testing.T) {
	t.Parallel()

	for _, data := range []linq.Query[float64]{testNums, linq.Reverse(testNums)} {
		assertSome(t, 1.0, linq.Min(data))
		assertNo(t, linq.Min(emptyNums))
	}
}

func TestMinBy(t *testing.T) {
	t.Parallel()

	type Person = linq.KV[string, int]

	peeps := linq.FromMap(map[string]int{"John": 42, "Andrea": 35, "Sanjiv": 22})
	noone := linq.FromMap(map[string]int{})

	name := linq.Key[Person]
	age := linq.Value[Person]

	assertSome(t, linq.NewKV("Andrea", 35), linq.MinBy(peeps, name))
	assertNo(t, linq.MinBy(noone, name))
	assertSome(t, linq.NewKV("Sanjiv", 22), linq.MinBy(peeps, age))
	assertNo(t, linq.MinBy(noone, age))
}
