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

// FromMap returns a query with KVs sourced from m.
func FromMap[K comparable, V any, M ~map[K]V](m M) Query[KV[K, V]] {
	if len(m) == 0 {
		return None[KV[K, V]]()
	}
	return FromSeq(
		func(yield func(KV[K, V]) bool) {
			seqMap(m)(func(k K, v V) bool {
				return yield(NewKV(k, v))
			})
		},
		FastCountOption[KV[K, V]](len(m)))
}

// ToMap converts a seq to a map, with sel providing key/value pairs. If any
// keys are duplicated, ToMap will return an error.
func ToMap[T, U any, K comparable](seq iter.Seq[T], sel func(t T) KV[K, U]) (map[K]U, error) {
	ret := map[K]U{}
	for t := range seq {
		k, v := sel(t).KV()
		if _, ok := ret[k]; ok {
			return nil, errorf("duplicate key %v", k)
		}
		ret[k] = v
	}
	return ret, nil
}

// ToMapKV converts a seq of KV pairs to a map. If any keys are duplicated,
// ToMapKV will return an error.
func ToMapKV[K comparable, V any](seq iter.Seq[KV[K, V]]) (map[K]V, error) {
	return ToMap(seq, Identity[KV[K, V]])
}

// MustToMap converts a seq to a map, with sel providing key/value pairs. If
// any keys are duplicated, MustToMap will panic.
func MustToMap[T, U any, K comparable](seq iter.Seq[T], sel func(t T) KV[K, U]) map[K]U {
	m, err := ToMap(seq, sel)
	if err != nil {
		panic(err)
	}
	return m
}

// MustToMapKV converts a seq of KV pairs to a map. If any keys are duplicated,
// MustToMapKV will panic.
func MustToMapKV[K comparable, V any](seq iter.Seq[KV[K, V]]) map[K]V {
	m, err := ToMapKV(seq)
	if err != nil {
		panic(err)
	}
	return m
}

// SelectKeys returns a query containing the keys from a query of KVs.
func SelectKeys[K, V any](q Query[KV[K, V]]) Query[K] {
	return Pipe(q, Select(q.Seq(), func(kv KV[K, V]) K { return kv.Key }))
}

// SelectValues returns a query containing the values from a query of KVs.
func SelectValues[K, V any](q Query[KV[K, V]]) Query[V] {
	return Pipe(q, Select(q.Seq(), func(kv KV[K, V]) V { return kv.Value }))
}
