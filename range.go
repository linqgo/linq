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

//go:build go1.22

package linq

import "iter"

func (q Query[T]) Range() iter.Seq[T] {
	return q.seq
}

func (q Query[T]) IRange() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		i := 0
		for t := range q.Range() {
			if !yield(i, t) {
				return
			}
			i++
		}
	}
}

func shunt[T any](seq iter.Seq[T], yield func(T) bool) {
	for t := range seq {
		if !yield(t) {
			return
		}
	}
}

func nextToSeq[T any](next func() (T, bool)) iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			if t, ok := next(); !ok || !yield(t) {
				return
			}
		}
	}
}
