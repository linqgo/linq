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

	"github.com/stretchr/testify/assert"
)

func TestAverage(t *testing.T) {
	t.Parallel()

	for _, data := range []linq.Query[float64]{testNums, linq.Reverse(testNums)} {
		assertSome(t, 5.5, maybe(linq.Average(data)))
		assertNo(t, maybe(linq.Average(emptyNums)))
	}
}

func TestSum(t *testing.T) {
	t.Parallel()

	for _, data := range []linq.Query[float64]{testNums, linq.Reverse(testNums)} {
		assert.EqualValues(t, 55, linq.Sum(data))
	}
}
