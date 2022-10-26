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

// Count returns the number of elements in q.
func (q Query[T]) Count() int {
	return Count(q)
}

// CountLimit returns a limited count, c, such that min(limit, Count(q)) <= c <=
// Count(q). This is useful for learning something about the size of the input
// without necessarily consuming it. One example is activating pagination
// controls for a result with at least 11 elements.
//
// If the query has a FastCount(), the return value is the true count.
func (q Query[T]) CountLimit(limit int) int {
	return CountLimit(q, limit)
}

// CountLimitTrue returns a limited count, c, such that min(limit, Count(q)) <=
// c <= Count(q). This is useful for learning something about the size of the
// input without necessarily consuming it. One example is activating pagination
// controls for a result with at least 11 elements.
//
// If the query has a FastCount(), the return value is the true count.
//
// The second return value is true if the returned count is the true count.
func (q Query[T]) CountLimitTrue(limit int) (int, bool) {
	return CountLimitTrue(q, limit)
}

// FastCount returns the number of elements in q if it can be computed in O(1)
// time, otherwise the second return value is false.
func (q Query[T]) FastCount() Maybe[int] {
	return FastCount(q)
}

// Count returns the number of elements in q.
func Count[T any](q Query[T]) int {
	if c, ok := q.FastCount().Get(); ok {
		return c
	}
	return Drain(q.Enumerator())
}

// CountLimit returns a limited count, c, such that min(limit, Count(q)) <= c <=
// Count(q). This is useful for learning something about the size of the input
// without necessarily consuming it. One example is activating pagination
// controls for a result with at least 11 elements.
//
// If the query has a FastCount(), the return value is the true count.
func CountLimit[T any](q Query[T], limit int) int {
	c, _ := CountLimitTrue(q, limit)
	return c
}

// CountLimitTrue returns a limited count, c, such that min(limit, Count(q)) <=
// c <= Count(q). This is useful for learning something about the size of the
// input without necessarily consuming it. One example is activating pagination
// controls for a result with at least 11 elements.
//
// If the query has a FastCount(), the return value is the true count.
//
// The second return value is true if the returned count is the true count.
func CountLimitTrue[T any](q Query[T], limit int) (int, bool) {
	if c, ok := FastCount(q).Get(); ok {
		return c, true
	}

	if n := Drain(Take(q, limit+1).Enumerator()); n <= limit {
		return n, true
	}
	return limit, false
}

// FastCount returns the number of elements in q if it can be computed in O(1)
// time, otherwise the second return value is false.
func FastCount[T any](q Query[T]) Maybe[int] {
	count := q.fastCount()
	return NewMaybe(count, count >= 0)
}
