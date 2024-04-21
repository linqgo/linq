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

// DefaultIfEmpty returns q if not empty, otherwise it returns a query
// containing alt.
func (q Query[T]) DefaultIfEmpty(alt T) Query[T] {
	return DefaultIfEmpty(q, alt)
}

// DefaultIfEmpty returns q if not empty, otherwise it returns a query
// containing alt.
func DefaultIfEmpty[T any](q Query[T], alt T) Query[T] {
	count := q.fastCount()
	switch count {
	case -1:
		return Pipe(q, func(yield func(T) bool) {
			delivered := false
			for t := range q.Seq() {
				if !yield(t) {
					return
				}
				delivered = true
			}
			if !delivered {
				yield(alt)
			}
		})
	case 0:
		return From(alt)
	default:
		return q
	}
}
