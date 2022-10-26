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

type Maybe[T any] struct {
	t     T
	valid bool
}

func NewMaybe[T any](t T, valid bool) Maybe[T] {
	return Maybe[T]{t: t, valid: valid}
}

func No[T any]() Maybe[T] {
	return Maybe[T]{}
}

func Some[T any](t T) Maybe[T] {
	return Maybe[T]{t: t, valid: true}
}

func (m Maybe[T]) Else(alt T) T {
	if m.valid {
		return m.t
	}
	return alt
}

func (m Maybe[T]) Get() (T, bool) {
	return m.t, m.valid
}

func (m Maybe[T]) Must() T {
	if m.valid {
		return m.t
	}
	panic(NoValueError)
}

func (m Maybe[T]) Valid() bool {
	return m.valid
}

func ElseNaN[R realNumber](r Maybe[R]) R {
	return r.Else(R(nan))
}
