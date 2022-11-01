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

func Pairwise[T any](q Query[T]) Query[KV[T, T]] {
	return Pipe(q, func(next Enumerator[T]) Enumerator[KV[T, T]] {
		if a, ok := next().Get(); ok {
			return func() Maybe[KV[T, T]] {
				if b, ok := next().Get(); ok {
					kv := NewKV(a, b)
					a = b
					return Some(kv)
				}
				return No[KV[T, T]]()
			}
		}
		return No[KV[T, T]]
	})
}
