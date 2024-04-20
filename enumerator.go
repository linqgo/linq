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

import (
	"iter"
	"runtime"
)

// Enumerator is a function type that enumerates values. To produce a value, it
// returns the value and true. When there are no more values to produce, it
// returns an indeterminate value and false.
type Enumerator[T any] func() Maybe[T]

type Enumerable[T any] func() Enumerator[T]

func newEnumerable[T any](seq iter.Seq[T]) Enumerable[T] {
	return func() Enumerator[T] {
		next := pull(seq)
		return func() Maybe[T] { return NewMaybe(next()) }
	}
}

func (e Enumerable[T]) Seq() iter.Seq[T] {
	return func(yield func(t T) bool) {
		next := e()
		for {
			t, has := next().Get()
			if !has || !yield(t) {
				return
			}
		}
	}
}

func pull[T any](seq iter.Seq[T]) func() (T, bool) {
	next, stop := iter.Pull(seq)
	runtime.SetFinalizer(&next, func(*func() (T, bool)) { stop() })
	return next
}
