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

// FromMap returns a query with KVs sourced from m.
func FromMap[K comparable, V any](m map[K]V) Query[KV[K, V]] {
	if len(m) == 0 {
		return None[KV[K, V]]()
	}
	return NewQuery(func() Enumerator[KV[K, V]] {
		keys := make([]K, 0, len(m))
		for k := range m {
			keys = append(keys, k)
		}
		ki := From(keys...).Enumerator()
		return func() Maybe[KV[K, V]] {
			if key, ok := ki().Get(); ok {
				return Some(NewKV(key, m[key]))
			}
			return No[KV[K, V]]()
		}
	}, FastCountOption[KV[K, V]](len(m)))
}

// MustToMap converts a query to a map, with sel providing key/value pairs. If
// any keys are duplicated, MustToMap will panic.
func MustToMap[T, U any, K comparable](q Query[T], sel func(t T) KV[K, U]) map[K]U {
	m, err := ToMap(q, sel)
	if err != nil {
		panic(err)
	}
	return m
}

// MustToMapKV converts a Query[KV[...]] to a map. If any keys are duplicated,
// MustToMapKV will panic.
func MustToMapKV[K comparable, V any](q Query[KV[K, V]]) map[K]V {
	m, err := ToMapKV(q)
	if err != nil {
		panic(err)
	}
	return m
}

// SelectKeys returns a query containing the keys from a query of KVs.
func SelectKeys[K, V any](q Query[KV[K, V]]) Query[K] {
	return Select(q, func(kv KV[K, V]) K { return kv.Key })
}

// SelectValues returns a query containing the values from a query of KVs.
func SelectValues[K, V any](q Query[KV[K, V]]) Query[V] {
	return Select(q, func(kv KV[K, V]) V { return kv.Value })
}

// ToMap converts a query to a map, with sel providing key/value pairs. If any
// keys are duplicated, ToMap will return an error.
func ToMap[T, U any, K comparable](q Query[T], sel func(t T) KV[K, U]) (map[K]U, error) {
	next := q.Enumerator()
	ret := map[K]U{}
	for t, ok := next().Get(); ok; t, ok = next().Get() {
		kv := sel(t)
		if _, ok := ret[kv.Key]; ok {
			return nil, errorf("duplicate key %v", kv.Key)
		}
		ret[kv.Key] = kv.Value
	}
	return ret, nil
}

// ToMapKV converts a Query[KV[...]] to a map. If any keys are duplicated, ToMap
// will return an error.
func ToMapKV[K comparable, V any](q Query[KV[K, V]]) (map[K]V, error) {
	return ToMap(q, Identity[KV[K, V]])
}
