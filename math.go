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

import (
	"github.com/linqgo/linq/internal/num"
)

// Average returns the arithmetic mean of the numbers in q or ok = false if q is
// empty.
//
// This function is equivalent to "github.com/linqgo/linq/stats".Mean.
// It is retained here for parity with .Net's Enumerable class.
func Average[R num.RealNumber](q Query[R]) Maybe[R] {
	if sum, n := aggregateN(q, 0, add[R]); n > 0 {
		return Some(sum / R(n))
	}
	return No[R]()
}

// Sum returns the sum of the num.Numbers in q or 0 if q is empty.
//
// This function is equivalent to "github.com/linqgo/linq/stats".Sum. It is
// retained here for parity with .Net's Enumerable class.
func Sum[R num.Number](q Query[R]) R {
	return aggregate(q, 0, add[R])
}

func add[N num.Number](a, b N) N {
	return a + b
}
