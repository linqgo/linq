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

package linq

import (
	"iter"

	"github.com/linqgo/linq/v2/internal/num"
)

// Average returns the arithmetic mean of the numbers in seq or ok = false if seq is
// empty.
//
// This function is equivalent to "./stats".Mean and is retained here for parity
// with .Net's Enumerable class.
func Average[R num.RealNumber](seq iter.Seq[R]) (R, bool) {
	if sum, n := aggregateNEnum(seq, 0, add[R]); n > 0 {
		return sum / R(n), true
	}
	return no[R]()
}

// Sum returns the sum of the num.Numbers in seq or 0 if seq is empty.
//
// This function is equivalent to "./stats".Sum and is retained here for parity
// with .Net's Enumerable class.
func Sum[R num.Number](seq iter.Seq[R]) R {
	return AggregateSeed(seq, 0, add[R])
}

func add[N num.Number](a, b N) N {
	return a + b
}
