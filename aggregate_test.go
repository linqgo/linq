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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/linqgo/linq/v2"
)

func add(a, b int) int { return a + b }

func TestAggregate(t *testing.T) {
	t.Parallel()

	assertNo(t, maybe(linq.Iota1(0).Aggregate(add)))
	assertSome(t, 15, maybe(linq.Iota2(1, 6).Aggregate(add)))
}

func TestAggregateSeed(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 42+15,
		linq.From(1, 2, 3, 4, 5).
			AggregateSeed(42, func(a, b int) int { return a + b }),
	)

	assert.Equal(t, 42,
		linq.From[int]().
			AggregateSeed(42, func(a, b int) int { return a + b }),
	)

	assert.Equal(t, ".1.2.3.4.5",
		linq.AggregateSeed(linq.From(1, 2, 3, 4, 5), "",
			func(a string, b int) string { return fmt.Sprintf("%s.%d", a, b) },
		),
	)
}
