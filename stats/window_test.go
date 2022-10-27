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

	"github.com/linqgo/linq/stats"
)

func TestWindowAll(t *testing.T) {
	t.Parallel()

	s := stats.WindowAll(chanof(1, 2, 3)).ToSlice()
	for i, kv := range s {
		assert.Equal(t, i+1, kv.Key)
		assertQueryEqual(t, []int{}, kv.Value)
	}
}

func TestWindowFixed(t *testing.T) {
	t.Parallel()

	s := stats.WindowFixed(chanof(1, 2, 3, 4, 5), 2).ToSlice()
	for i, kv := range s {
		t.Log(i, kv.Key, kv.Value.ToSlice())
	}
	for i, kv := range s[:2] {
		assert.Equal(t, i+1, kv.Key)
		assertQueryEqual(t, []int{}, kv.Value)
	}
	for i, kv := range s[2:] {
		assert.Equal(t, i+3, kv.Key)
		assertQueryEqual(t, []int{i + 1}, kv.Value)
	}
}
