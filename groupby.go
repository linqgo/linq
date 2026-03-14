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

// GroupBySlices returns an iter.Seq of KV pairs with elements from seq grouped
// using the specified key function, with values as slices.
func GroupBySlices[T any, K comparable](
	seq iter.Seq[T],
	key func(t T) K,
) iter.Seq[KV[K, []T]] {
	return GroupBySelectSlices(seq, keyIdentity(key))
}

// GroupBySelectSlices returns an iter.Seq of KV pairs with elements from seq
// grouped using the specified sel function.
func GroupBySelectSlices[T, U any, K comparable](
	seq iter.Seq[T],
	sel func(t T) KV[K, U],
) iter.Seq[KV[K, []U]] {
	return func(yield func(KV[K, []U]) bool) {
		m := map[K][]U{}
		var keys []K
		for t := range seq {
			k, v := sel(t).KV()
			if _, exists := m[k]; !exists {
				keys = append(keys, k)
			}
			m[k] = append(m[k], v)
		}
		for _, k := range keys {
			if !yield(NewKV(k, m[k])) {
				return
			}
		}
	}
}

// GroupBy returns a Query[KV[K, Query[T]]] with elements from q grouped using
// the specified key function.
func GroupBy[T any, K comparable](
	q Query[T],
	key func(t T) K,
) Query[KV[K, Query[T]]] {
	return GroupBySelect(q, keyIdentity(key))
}

// GroupBySlicesQuery returns a Query[KV[K, []T]] with elements from q grouped
// using the specified key function.
func GroupBySlicesQuery[T any, K comparable](
	q Query[T],
	key func(t T) K,
) Query[KV[K, []T]] {
	return GroupBySelectSlicesQuery(q, keyIdentity(key))
}

// GroupBySelect returns a Query[KV[K, Query[T]]] with elements from q grouped
// using the specified sel function.
func GroupBySelect[T, U any, K comparable](
	q Query[T],
	sel func(t T) KV[K, U],
) Query[KV[K, Query[U]]] {
	inner := GroupBySelectSlicesQuery(q, sel)
	return Pipe(inner, Select(
		inner.Seq(),
		func(kv KV[K, []U]) KV[K, Query[U]] {
			return NewKV(kv.Key, From(kv.Value...))
		},
	),
		FastCountIfEmptyOption[KV[K, Query[U]]](inner.fastCount()),
	)
}

// GroupBySelectSlicesQuery returns a Query[KV[K, []T]] with elements from q
// grouped using the specified sel function.
func GroupBySelectSlicesQuery[T, U any, K comparable](
	q Query[T],
	sel func(t T) KV[K, U],
) Query[KV[K, []U]] {
	return Pipe(q,
		GroupBySelectSlices(q.Seq(), sel),
		FastCountIfEmptyOption[KV[K, []U]](q.fastCount()))
}

func keyIdentity[T any, K comparable](key func(t T) K) func(t T) KV[K, T] {
	return func(t T) KV[K, T] {
		return NewKV(key(t), t)
	}
}
