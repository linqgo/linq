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

	"github.com/linqgo/linq/v2"
)

var (
	testNums  = linq.Select(linq.Iota2(1, 11).Seq(), func(i int) float64 { return float64(i) })
	emptyNums = linq.None[float64]()
)

func TestMax(t *testing.T) {
	t.Parallel()

	assertSome(t, 10.0, maybe(linq.Max(testNums)))
	assertNo(t, maybe(linq.Max(emptyNums.Seq())))

	// Also test with reversed data via Query
	rev := linq.FromSeq(testNums).Reverse()
	assertSome(t, 10.0, maybe(linq.Max(rev.Seq())))
}

func TestMaxBy(t *testing.T) {
	t.Parallel()

	type Person = linq.KV[string, int]

	peeps := linq.FromMap(map[string]int{"John": 42, "Sanjiv": 22, "Andrea": 35})
	noone := linq.FromMap(map[string]int{})

	name := func(kv Person) string { return kv.Key }
	age := func(kv Person) int { return kv.Value }

	assertSome(t, linq.NewKV("Sanjiv", 22), maybe(linq.MaxBy(peeps, name).Single()))
	assertNo(t, maybe(linq.MaxBy(noone, name).Single()))
	assertSome(t, linq.NewKV("John", 42), maybe(linq.MaxBy(peeps, age).Single()))
	assertNo(t, maybe(linq.MaxBy(noone, age).Single()))

	repeats := linq.From(1, 2, 4, 2, 3, 1)
	mod3 := func(i int) int { return i % 3 }
	assertQueryEqual(t, []int{2, 2}, (linq.MaxBy(repeats, mod3)))
	assertQueryEqual(t, []int{2}, linq.MaxBy(repeats, mod3).Take(1))
}

func TestMin(t *testing.T) {
	t.Parallel()

	assertSome(t, 1.0, maybe(linq.Min(testNums)))
	assertNo(t, maybe(linq.Min(emptyNums.Seq())))

	rev := linq.FromSeq(testNums).Reverse()
	assertSome(t, 1.0, maybe(linq.Min(rev.Seq())))
}

func TestMinBy(t *testing.T) {
	t.Parallel()

	type Person = linq.KV[string, int]

	peeps := linq.FromMap(map[string]int{"John": 42, "Andrea": 35, "Sanjiv": 22})
	noone := linq.FromMap(map[string]int{})

	name := linq.Key[Person]
	age := linq.Value[Person]

	assertSome(t, linq.NewKV("Andrea", 35), maybe(linq.MinBy(peeps, name).Single()))
	assertNo(t, maybe(linq.MinBy(noone, name).Single()))
	assertSome(t, linq.NewKV("Sanjiv", 22), maybe(linq.MinBy(peeps, age).Single()))
	assertNo(t, maybe(linq.MinBy(noone, age).Single()))
}
