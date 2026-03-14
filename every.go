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

// Every returns a query that contains every nth element from q.
func (q Query[T]) Every(n int) Query[T] {
	return q.EveryFrom(0, n)
}

// EveryFrom returns a query that contains every nth element from q, starting at the
// start-th element.
func (q Query[T]) EveryFrom(start, n int) Query[T] {
	return Pipe(q, EveryFrom(q.Seq(), start, n),
		ComputedFastCountOption[T](q.fastCount(), func(count int) int {
			return (count-start-1)/n + 1
		}))
}

// Every returns a seq that contains every nth element from seq.
func Every[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	return EveryFrom(seq, 0, n)
}

// EveryFrom returns a seq that contains every nth element from seq, starting at the
// start-th element.
func EveryFrom[T any](seq iter.Seq[T], start, n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		i := 0
		seq(func(t T) bool {
			cont := !(start <= i && (i-start)%n == 0) || yield(t)
			i++
			return cont
		})
	}
}
