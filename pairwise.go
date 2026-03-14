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

import "iter"

// Pairwise returns a seq of consecutive pairs from seq.
func Pairwise[T any](seq iter.Seq[T]) iter.Seq[KV[T, T]] {
	return func(yield func(KV[T, T]) bool) {
		var a T
		i := 0
		seq(func(b T) bool {
			cont := i == 0 || yield(NewKV(a, b))
			a = b
			i++
			return cont
		})
	}
}

// PairwiseQuery returns a query of consecutive pairs from q.
func PairwiseQuery[T any](q Query[T]) Query[KV[T, T]] {
	return Pipe(q, Pairwise(q.Seq()))
}
