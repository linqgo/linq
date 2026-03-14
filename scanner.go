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

func (q Query[T]) Scanner() (pull func(p *T) bool, stop func()) {
	return Scanner(q)
}

// Scanner returns a function that scans entries from q, setting the passed
// pointer parameter to each element and returning true until it runs out of
// elements. Then it returns false.
func Scanner[T any](q Query[T]) (scan func(p *T) bool, stop func()) {
	next, stop := iter.Pull(q.Seq())
	return func(p *T) bool {
		if t, ok := next(); ok {
			*p = t
			return true
		}
		return false
	}, stop
}
