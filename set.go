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

type set[T comparable] map[T]struct{}

func (s set[T]) Has(t T) bool {
	_, ok := s[t]
	return ok
}

func (s set[T]) Add(t T) {
	s[t] = struct{}{}
}

func setFrom[T comparable](i Enumerator[T]) set[T] {
	m := set[T]{}
	for t, ok := i().Get(); ok; t, ok = i().Get() {
		m.Add(t)
	}
	return m
}
