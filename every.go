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

// Every returns a query that contains every nth element from q.
func (q Query[T]) Every(n int) Query[T] {
	return Every(q, n)
}

// Every returns a query that contains every nth element from q, starting at the
// start-th element.
func (q Query[T]) EveryFrom(start, n int) Query[T] {
	return EveryFrom(q, start, n)
}

// Every returns a query that contains every nth element from q.
func Every[T any](q Query[T], n int) Query[T] {
	return EveryFrom(q, 0, n)
}

// Every returns a query that contains every nth element from q, starting at the
// start-th element.
func EveryFrom[T any](q Query[T], start, n int) Query[T] {
	return Pipe(q, func(next Enumerator[T]) Enumerator[T] {
		skip := start
		return func() Maybe[T] {
			for t := next(); t.Valid(); t = next() {
				if skip == 0 {
					skip = n - 1
					return t
				}
				skip--
			}
			return No[T]()
		}
	}, ComputedFastCountOption[T](q.fastCount(), func(count int) int {
		return (count-start-1)/n + 1
	}))
}
