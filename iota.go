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

import "github.com/linqgo/linq/v2/internal/num"

// Iota returns a query with all integers from 0 up.
func Iota[R num.RealNumber]() Query[R] {
	return FromSeq(
		func(yield func(R) bool) {
			for i := R(0); ; i++ {
				if !yield(i) {
					return
				}
			}
		},
		FastGetOption(func(i int) (R, bool) { return R(i), true }),
	)
}

// Iota1 returns a query with all integers in the range [0, stop).
func Iota1[R num.RealNumber](stop R) Query[R] {
	return Iota3(0, stop, 1)
}

// Iota2 returns a query with all integers in the range [start, stop).
func Iota2[R num.RealNumber](start, stop R) Query[R] {
	return Iota3(start, stop, 1)
}

// Iota3 returns a query with every step-th integer in the range [start, stop).
func Iota3[R num.RealNumber](start, stop, step R) Query[R] {
	switch {
	case step > 0:
		n := (stop-start-1)/step + 1
		return FromSeq(
			func(yield func(R) bool) {
				seqN(n)(func(i R) bool {
					return yield(start + step*i)
				})
			},
			FastCountOption[R](int(n)),
			FastGetOption(LenGetGetter(int(n), func(i int) R {
				return start + step*R(i)
			})),
		)
	case step < 0:
		n := int((start-stop-1)/-step + 1)
		return FromSeq(
			func(yield func(R) bool) {
				for i := start; i > stop; i += step {
					if !yield(i) {
						return
					}
				}
			},
			FastCountOption[R](n),
			FastGetOption(LenGetGetter(n, func(i int) R {
				return start + step*R(i)
			})),
		)
	default:
		if start == stop {
			return None[R]()
		}
		panic(ZeroIotaStepError)
	}
}
