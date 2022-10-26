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

// Prepend returns a query with the elements of t followed by the elements of q.
//
// Be careful! Prepend puts the tail arguments, t, in front of q. To avoid
// confusion and bugs, consider using the corresponding global Prepend function.
func (q Query[T]) Prepend(t ...T) Query[T] {
	return Prepend(t...)(q)
}

// Prepend returns a query with the elements of t followed by the elements of q.
//
// So that t... args appear before q, Prepend takes just t and returns a func
// that takes q.
func Prepend[T any](t ...T) func(q Query[T]) Query[T] {
	return func(q Query[T]) Query[T] {
		return From(t...).Concat(q)
	}
}
