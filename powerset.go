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
	"math/bits"
	"unsafe"
)

// PowerSet returns the power set of seq as an iter.Seq of iter.Seq.
func PowerSet[T any](seq iter.Seq[T]) iter.Seq[iter.Seq[T]] {
	return func(yield func(iter.Seq[T]) bool) {
		var cache []T
		if !yield(func(func(T) bool) {}) {
			return
		}
		i := 0
		for t := range seq {
			cache = append(cache, t)
			hi := 1 << i
			stop := false
			seqN(hi)(func(mask int) bool {
				if !yield(powerSubSeq(cache, uint64(hi+mask))) {
					stop = true
					return false
				}
				return true
			})
			if stop {
				return
			}
			i++
		}
	}
}

// PowerSetQuery returns the power set of q as a Query of Queries.
func PowerSetQuery[T any](q Query[T]) Query[Query[T]] {
	// Calculate the number of bits available in a positive int (minus one
	// because the high bit is reserved for negative ints).
	const positiveIntBits = int(8*unsafe.Sizeof(int(0))) - 1

	// Slight problem: If FastCount(q) >= 64, then the actual count can't be
	// represented with an int.
	count := -1
	if c, ok := q.FastCount(); ok && c < positiveIntBits {
		count = 1 << c
	}

	return Pipe(q,
		Select(PowerSet(q.Seq()), func(s iter.Seq[T]) Query[T] {
			// Collect the seq into a slice for Query wrapping
			var data []T
			for t := range s {
				data = append(data, t)
			}
			return From(data...)
		}),
		FastCountOption[Query[T]](count),
	)
}

func powerSubSeq[T any](cache []T, mask uint64) iter.Seq[T] {
	return func(yield func(T) bool) {
		for m := mask; m != 0; m &= m - 1 {
			if !yield(cache[bits.TrailingZeros64(m)]) {
				return
			}
		}
	}
}
