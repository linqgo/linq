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

func TestAll(t *testing.T) {
	t.Parallel()

	assert.True(t, linq.Iota1(5).All(func(t int) bool { return t < 10 }))
	assert.False(t, linq.Iota1(5).All(func(t int) bool { return t < 3 }))
	assert.True(t, linq.Iota1(0).All(linq.False[int]))
}

func TestAny(t *testing.T) {
	t.Parallel()

	assert.True(t, linq.Iota1(5).Any(func(t int) bool { return t < 3 }))
	assert.False(t, linq.Iota1(5).Any(func(t int) bool { return t > 10 }))
	assert.False(t, linq.Iota1(0).Any(linq.True[int]))
}

func TestEmpty(t *testing.T) {
	t.Parallel()

	assert.True(t, linq.None[int]().Empty())
	assert.False(t, linq.From(1, 2, 3).Empty())
}
