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

// All returns true if pred returns true for all elements in q, including if q
// is empty.
func (q Query[T]) All(pred func(t T) bool) bool {
	return All(q.Seq(), pred)
}

// Any returns true if pred returns true for at least one element in q.
func (q Query[T]) Any(pred func(t T) bool) bool {
	return Any(q.Seq(), pred)
}

// Empty returns true if q has no elements.
func (q Query[T]) Empty() bool {
	return Empty(q.Seq())
}

// All returns true if pred returns true for all elements in seq, including if seq
// is empty.
func All[T any](seq iter.Seq[T], pred func(t T) bool) bool {
	return !Any(seq, Not(pred))
}

// Any returns true if pred returns true for at least one element in seq.
func Any[T any](seq iter.Seq[T], pred func(t T) bool) bool {
	for t := range seq {
		if pred(t) {
			return true
		}
	}
	return false
}

// Empty returns true if seq has no elements.
func Empty[T any](seq iter.Seq[T]) bool {
	for t := range seq {
		_ = t
		return false
	}
	return true
}
