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

func (q Query[T]) Single() (T, bool) {
	return Single(q)
}

func Single[T any](q Query[T]) (T, bool) {
	i := -1
	var t T
	for i, t = range q.ISeq() {
		if i == 1 {
			return no[T]()
		}
	}
	return t, i == 0
}
