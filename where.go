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

// Where returns a query with elements from q for which pred returns true.
func (q Query[T]) Where(pred func(t T) bool) Query[T] {
	return Pipe(q, Where(q.Seq(), pred),
		FastCountIfEmptyOption[T](q.fastCount()),
	)
}

// Where returns a seq with elements from seq for which pred returns true.
func Where[T any](seq iter.Seq[T], pred func(t T) bool) iter.Seq[T] {
	return func(yield func(t T) bool) {
		seq(func(t T) bool { return !pred(t) || yield(t) })
	}
}
