// Copyright 2022 Marcelo Cantos
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0//
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package linq

import "golang.org/x/exp/constraints"

// Iota returns a query with all integers from 0 up.
func Iota[I constraints.Integer]() Query[I] {
	return NewQuery(
		func() Enumerator[I] {
			var i I
			i--
			return func() Maybe[I] {
				i++
				return Some(i)
			}
		},
		FastGetOption(func(i int) Maybe[I] { return Some(I(i)) }),
	)
}

// Iota1 returns a query with all integers in the range [0, stop).
func Iota1[I constraints.Integer](stop I) Query[I] {
	return Iota3(0, stop, 1)
}

// Iota2 returns a query with all integers in the range [start, stop).
func Iota2[I constraints.Integer](start, stop I) Query[I] {
	return Iota3(start, stop, 1)
}

// Iota3 returns a query with every step-th integer in the range [start, stop).
func Iota3[I constraints.Integer](start, stop, step I) Query[I] {
	switch {
	case step > 0:
		n := int((stop-start-1)/step + 1)
		return NewQuery(
			func() Enumerator[I] {
				i := start - step
				return func() Maybe[I] {
					i += step
					return NewMaybe(i, i < stop)
				}
			},
			FastCountOption[I](n),
			FastGetOption(LenGetGetter(n, func(i int) I {
				return start + step*I(i)
			})),
		)
	case step < 0:
		n := int((start-stop-1)/-step + 1)
		return NewQuery(
			func() Enumerator[I] {
				i := start - step
				return func() Maybe[I] {
					i += step
					return NewMaybe(i, i > stop)
				}
			},
			FastCountOption[I](n),
			FastGetOption(LenGetGetter(n, func(i int) I {
				return start + step*I(i)
			})),
		)
	default:
		if start == stop {
			return None[I]()
		}
		panic(ZeroIotaStepError)
	}
}
