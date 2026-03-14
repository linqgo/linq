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

// Count returns the number of elements in q.
func (q Query[T]) Count() int {
	if c, ok := q.FastCount(); ok {
		return c
	}
	return Count(q.Seq())
}

// CountLimit returns a limited count, c, such that min(limit, Count(q)) <= c <=
// Count(q). This is useful for learning something about the size of the input
// without necessarily consuming it. One example is activating pagination
// controls for a result with at least 11 elements.
//
// If the query has a FastCount() the return value is the true count.
func (q Query[T]) CountLimit(limit int) int {
	if c, ok := q.FastCount(); ok {
		return min(limit, c)
	}
	return CountLimit(q.Seq(), limit)
}

// FastCount returns the number of elements in q if it can be computed in O(1)
// time, otherwise the second return value is false.
func (q Query[T]) FastCount() (int, bool) {
	count := q.fastCount()
	return count, count >= 0
}

// Count returns the number of elements in seq.
func Count[T any](seq iter.Seq[T]) int {
	n := 0
	for t := range seq {
		_ = t
		n++
	}
	return n
}

// CountLimit returns a limited count. This is useful for learning something
// about the size of the input without having to consume all of it. One example
// is determining whether pagination is required.
//
// To count up to N, but also learn if there are more than N elements:
//
//	n := CountLimit(N + 1)
//	n, more := min(n, N), n > N
func CountLimit[T any](seq iter.Seq[T], limit int) int {
	n := 0
	for t := range seq {
		if n == limit {
			return n
		}
		_ = t
		n++
	}
	return n
}
