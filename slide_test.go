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

func TestSlide(t *testing.T) {
	data := [][][]int{
		{nil, {1}},
		{nil, {2}},
		{nil, {3}},
		{{1, 2}, {5, 5}},
		{{3}, {6, 7}},
		{{5, 5}, {8}},
	}
	slide := func(slideIn bool) linq.Query[[][]int] {
		return linq.Select(
			linq.Slide(
				linq.From(1, 2, 3, 5, 5, 6, 7, 8),
				slideIn,
				func(tail, head int) bool { return tail < head-2 },
			),
			func(d linq.Delta[int]) [][]int {
				return [][]int{
					d.Outs.ToSlice(),
					d.Ins.ToSlice(),
				}
			},
		)
	}

	assertQueryEqual(t, data, slide(true))
	assertQueryEqual(t,
		append([][][]int{{nil, {1, 2, 3}}}, data[3:]...),
		slide(false))
}

func TestSlideAll(t *testing.T) {
	t.Parallel()

	s := linq.SlideAll(chanof(1, 2, 3)).ToSlice()
	for i, kv := range s {
		assertQueryEqual(t, []int{i + 1}, kv.Ins)
		assertQueryEqual(t, []int{}, kv.Outs)
	}
}

func TestSlideFixed(t *testing.T) {
	t.Parallel()

	s := linq.SlideFixed(chanof(1, 2, 3, 4, 5), 2, true).ToSlice()
	for i, kv := range s {
		t.Log(i, kv.Ins.ToSlice(), kv.Outs.ToSlice())
	}
	for i, kv := range s[:2] {
		assertQueryEqual(t, []int{i + 1}, kv.Ins)
		assertQueryEqual(t, []int{}, kv.Outs)
	}
	for i, kv := range s[2:] {
		assertQueryEqual(t, []int{i + 3}, kv.Ins)
		assertQueryEqual(t, []int{i + 1}, kv.Outs)
	}
}
