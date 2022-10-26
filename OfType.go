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

// OfType returns a Query that contains all the elements of q that have type U.
func OfType[U, T any](q Query[T]) Query[U] {
	return Pipe(q, func(next Enumerator[T]) Enumerator[U] {
		return func() Maybe[U] {
			for t, ok := next().Get(); ok; t, ok = next().Get() {
				var i any = t
				if u, is := i.(U); is {
					return Some(u)
				}
			}
			return No[U]()
		}
	}, FastCountIfEmptyOption[U](q.fastCount()))
}
