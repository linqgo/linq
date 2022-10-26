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
	"strconv"
	"testing"

	"github.com/linqgo/linq"
)

var fiveInts = func() <-chan int {
	c := make(chan int, 5)
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
	return c
}

func TestMemoize(t *testing.T) {
	t.Parallel()

	q := linq.FromChannel(fiveInts())
	assertQueryEqual(t, []int{0, 1, 2, 3, 4}, q)
	assertQueryEqual(t, []int{}, q)

	m := linq.FromChannel(fiveInts()).Memoize()
	assertQueryEqual(t, []int{0, 1, 2, 3, 4}, m)
	assertQueryEqual(t, []int{0, 1, 2, 3, 4}, m)

	assertOneShot(t, true, linq.OfType[int](q))
	assertOneShot(t, false, linq.OfType[int](m))
}

func TestMemoizeParallel(t *testing.T) {
	t.Parallel()

	m := linq.FromChannel(fiveInts()).Memoize()
	for i := 0; i < 10; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			for j := 0; j < 1000; j++ {
				if !assertQueryEqual(t, []int{0, 1, 2, 3, 4}, m) {
					break
				}
			}
		})
	}
}
