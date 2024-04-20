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

// ElementAt returns the element at position i or !ok if there is no element i.
func (q Query[T]) ElementAt(i int) (T, bool) {
	return ElementAt(q, i)
}

// FastElementAt returns the element at position i or !ok if there is no element
// i or the element cannot be accessed in O(1) time.
func (q Query[T]) FastElementAt(i int) (T, bool) {
	return FastElementAt(q, i)
}

// First returns the first element or !ok if q is empty.
func (q Query[T]) First() (T, bool) {
	return First(q)
}

// Last returns the last element or !ok if q is empty.
func (q Query[T]) Last() (T, bool) {
	return Last(q)
}

// ElementAt returns the element at position i or !ok if there is no element i.
func ElementAt[T any](q Query[T], at int) (T, bool) {
	if at < 0 {
		var zero T
		return zero, false
	}
	if t, ok := FastElementAt(q, at); ok {
		return t, ok
	}
	for i, t := range q.ISeq() {
		if i == at {
			return t, true
		}
	}
	var zero T
	return zero, false
}

// FastElementAt returns the element at position i or !ok if there is no element
// i or element i cannot be accessed in O(1) time.
func FastElementAt[T any](q Query[T], i int) (T, bool) {
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

// FastLast returns the last element or !ok if q is empty or the last element
// cannot be accessed in O(1) time.
func FastLast[T any](q Query[T]) (T, bool) {
	if q.count > 0 {
		if get := q.getter(); get != nil {
			return get(q.count - 1)
		}
	}
	return no[T]()
}

// First returns the first element or !ok if q is empty.
func First[T any](q Query[T]) (T, bool) {
	for t := range q.Seq() {
		return t, true
	}
	return no[T]()
}

// Last returns the last element or !ok if q is empty.
func Last[T any](q Query[T]) (T, bool) {
	if t, ok := FastLast(q); ok {
		return t, ok
	}
	var t T
	ok := false
	for t = range q.Seq() {
		ok = true
	}
	return t, ok
}
