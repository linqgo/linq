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

// Union returns the set union of a and b.
func Union[T comparable](a, b iter.Seq[T]) iter.Seq[T] {
	return Concat(a, Except(b, a))
}

// UnionQuery returns the set union of a and b as a Query.
func UnionQuery[T comparable](a, b Query[T]) Query[T] {
	return a.Concat(ExceptQuery(b, a))
}
