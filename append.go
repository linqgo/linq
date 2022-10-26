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

// Append returns a query with the elements of q followed by the elements of t.
func (q Query[T]) Append(t ...T) Query[T] {
	return Append(q, t...)
}

// Append returns a query with the elements of q followed by the elements of t.
func Append[T any](q Query[T], t ...T) Query[T] {
	return q.Concat(From(t...))
}
