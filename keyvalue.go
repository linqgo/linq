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

// KV represents a key/value pair.
type KV[K, V any] struct {
	Key   K
	Value V
}

// NewKV returns a new KV.
func NewKV[K, V any](key K, value V) KV[K, V] {
	return KV[K, V]{
		Key:   key,
		Value: value,
	}
}

func (kv KV[K, V]) KV() (K, V) {
	return kv.Key, kv.Value
}

func Key[KeyVal KV[K, V], K, V any](kv KV[K, V]) K {
	return kv.Key
}

func Value[KeyVal KV[K, V], K, V any](kv KV[K, V]) V {
	return kv.Value
}
