// Copyright 2022 Marcelo Cantos
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0//
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package linq

func Index[T any](q Query[T]) Query[KV[int, T]] {
	return IndexFrom(q, 0)
}

func IndexFrom[T any](q Query[T], start int) Query[KV[int, T]] {
	var get Getter[KV[int, T]]
	if qget := q.getter(); qget != nil {
		get = func(i int) Maybe[KV[int, T]] {
			if t, ok := qget(i).Get(); ok {
				return Some(NewKV(start+i, t))
			}
			return No[KV[int, T]]()
		}
	}
	return PipeOneToOne(q,
		func() func(t T) KV[int, T] {
			i := start - 1
			return func(t T) KV[int, T] {
				i++
				return NewKV(i, t)
			}
		},
		FastGetOption(get),
	)
}
