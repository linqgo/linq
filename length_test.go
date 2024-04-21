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

	"github.com/stretchr/testify/assert"

	"github.com/linqgo/linq/v2"
)

func TestShorter(t *testing.T) {
	t.Parallel()

	data := []string{"", "hello", "hello!", "z"}

	for _, a := range data {
		qa := linq.FromString(a)
		for _, b := range data {
			qb := linq.FromString(b)

			sr := len(a) < len(b)
			assertAll(t, func(qr bool) bool {
				return assert.Equal(t, sr, qr,
					"len(%q) < len(%q) expected %v, got %v", a, b, sr, qr)
			},
				qa.Shorter(qb), linq.Shorter(qa, qb),
				qa.Shorter(chanof([]rune(b)...)), linq.Shorter(qa, chanof([]rune(b)...)),
				chanof([]rune(a)...).Shorter(qb), linq.Shorter(chanof([]rune(a)...), qb),
				must(qa.FastShorter(qb)), must(linq.FastShorter(qa, qb)),
			)

			assertAll(t, func(qr func() (bool, bool)) bool {
				v, valid := qr()
				return assert.True(t, valid) && assert.Equal(t, sr, v,
					"len(%q) < len(%q) expected %v, got %v", a, b, sr, v)
			},
				maybe(qa.FastShorter(qb)), maybe(linq.FastShorter(qa, qb)),
			)

			assertAll(t, func(qr func() (bool, bool)) bool { return assertNo(t, qr) },
				maybe(qa.FastShorter(chanof([]rune(b)...))),
				maybe(linq.FastShorter(qa, chanof([]rune(b)...))),
				maybe(chanof([]rune(a)...).FastShorter(qb)),
				maybe(linq.FastShorter(chanof([]rune(a)...), qb)),
			)
		}
	}
}

func TestLonger(t *testing.T) {
	t.Parallel()

	data := []string{"", "hello", "hello!", "z"}

	for _, a := range data {
		qa := linq.FromString(a)
		for _, b := range data {
			qb := linq.FromString(b)

			sr := len(a) > len(b)
			assertAll(t, func(qr bool) bool {
				return assert.Equal(t, sr, qr,
					"len(%q) < len(%q) expected %v, got %v", a, b, sr, qr)
			},
				qa.Longer(qb), linq.Longer(qa, qb),
				qa.Longer(chanof([]rune(b)...)), linq.Longer(qa, chanof([]rune(b)...)),
				chanof([]rune(a)...).Longer(qb), linq.Longer(chanof([]rune(a)...), qb),
				must(qa.FastLonger(qb)), must(linq.FastLonger(qa, qb)),
			)

			assertAll(t, func(qr func() (bool, bool)) bool {
				v, valid := qr()
				return assert.True(t, valid) && assert.Equal(t, sr, v,
					"len(%q) < len(%q) expected %v, got %v", a, b, sr, v)
			},
				maybe(qa.FastLonger(qb)), maybe(linq.FastLonger(qa, qb)),
			)

			assertAll(t, func(qr func() (bool, bool)) bool { return assertNo(t, qr) },
				maybe(qa.FastLonger(chanof([]rune(b)...))),
				maybe(linq.FastLonger(qa, chanof([]rune(b)...))),
				maybe(chanof([]rune(a)...).FastLonger(qb)),
				maybe(linq.FastLonger(chanof([]rune(a)...), qb)),
			)
		}
	}
}
