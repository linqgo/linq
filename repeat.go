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

package linq

import "golang.org/x/exp/constraints"

// Repeat returns a query with value repeated count times.
func Repeat[T any, I constraints.Integer](value T, count I) Query[T] {
	if count == 0 {
		return None[T]()
	}
	n := int(count)
	return NewQuery(
		func() Enumerator[T] {
			var i I = 0
			return func() Maybe[T] {
				if i < count {
					i++
					return Some(value)
				}
				return No[T]()
			}
		},
		FastCountOption[T](int(count)),
		FastGetOption(func(i int) Maybe[T] {
			return NewMaybe(value, 0 <= i && i < n)
		}),
	)
}

// RepeatForever returns a query with value repeated forever.
func RepeatForever[T any](value T) Query[T] {
	return NewQuery(
		func() Enumerator[T] {
			return func() Maybe[T] {
				return Some(value)
			}
		},
		FastGetOption(func(i int) Maybe[T] {
			return NewMaybe(value, 0 <= i)
		}),
	)
}
