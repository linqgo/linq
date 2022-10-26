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

// TODO: Figure out Go-specific embodiments of Cast.

// // Cast returns a Query that contains all the elements of q that have type U.
// func Cast[U, T realNumber](q Query[T]) Query[U] {
// 	return NewQuery(func() Enumerator[U] {
// 		next := q.Enumerator()
// 		return func() (U, bool) {
// 			if t, ok := next(); ok {
// 				return U(t), true
// 			}
// 			var u U
// 			return u, false
// 		}
// 	})
// }
