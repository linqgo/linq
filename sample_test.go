// Copyright 2022 Marcelo Cantos
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0//
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

func TestRandomSample(t *testing.T) {
	t.Parallel()

	assert.InEpsilon(t, 1000, linq.Iota1(10000).Sample(0.1).Count(), 1.1)

	assertQueryEqual(t, []int{1, 3, 4, 5, 6, 9}, linq.Iota1(10).SampleSeed(0.6, 0))
	assertQueryEqual(t, []int{0, 8, 9, 12}, linq.Iota1(20).SampleSeed(0.3, 1234))

	assertOneShot(t, false, linq.Iota1(10000).Sample(0.1))
	assertOneShot(t, true, oneshot().Sample(0.1))
	assertOneShot(t, false, linq.Iota1(10).SampleSeed(0.6, 0))
	assertOneShot(t, true, oneshot().SampleSeed(0.6, 0))
}
