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

// ElementAt returns the element at position i or !ok if there is no element i.
func (q Query[T]) ElementAt(i int) (T, bool) {
	if i < 0 {
		var zero T
		return zero, false
	}
	if t, ok := q.FastElementAt(i); ok {
		return t, ok
	}
	return ElementAt(q.Seq(), i)
}

// FastElementAt returns the element at position i or !ok if there is no element
// i or the element cannot be accessed in O(1) time.
func (q Query[T]) FastElementAt(i int) (T, bool) {
	if i < 0 {
		var zero T
		return zero, false
	}
	if q.fastCount() != 0 {
		if get := q.getter(); get != nil {
			return get(i)
		}
	}
	return no[T]()
}

// First returns the first element or !ok if q is empty.
func (q Query[T]) First() (T, bool) {
	return First(q.Seq())
}

// Last returns the last element or !ok if q is empty.
func (q Query[T]) Last() (T, bool) {
	if t, ok := FastLast(q); ok {
		return t, ok
	}
	return Last(q.Seq())
}

// ElementAt returns the element at position i or !ok if there is no element i.
func ElementAt[T any](seq iter.Seq[T], at int) (T, bool) {
	if at < 0 {
		var zero T
		return zero, false
	}
	i := 0
	for t := range seq {
		if i == at {
			return t, true
		}
		i++
	}
	var zero T
	return zero, false
}

// FastLast returns the last element or !ok if q is empty or the last element
// cannot be accessed in O(1) time.
func FastLast[T any](q Query[T]) (T, bool) {
	if c, ok := q.FastCount(); ok && c > 0 {
		if get := q.getter(); get != nil {
			return get(c - 1)
		}
	}
	return no[T]()
}

// First returns the first element or !ok if seq is empty.
func First[T any](seq iter.Seq[T]) (T, bool) {
	for t := range seq {
		return t, true
	}
	return no[T]()
}

// Last returns the last element or !ok if seq is empty.
func Last[T any](seq iter.Seq[T]) (T, bool) {
	var t T
	ok := false
	for t = range seq {
		ok = true
	}
	return t, ok
}
