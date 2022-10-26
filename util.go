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

import "golang.org/x/exp/constraints"

// Identity returns t unmodified.
func Identity[T any](t T) T {
	return t
}

// Deref returns *t.
func Deref[T any](t *T) T {
	return *t
}

// Equal returns a == b.
func Equal[T comparable](a, b T) bool {
	return a == b
}

// False returns false, ignoring the input value.
func False[T any](T) bool {
	return false
}

// Greater returns a > b.
func Greater[T constraints.Ordered](a, b T) bool {
	return a > b
}

// Less returns a < b.
func Less[T constraints.Ordered](a, b T) bool {
	return a < b
}

// LongerSlice returns len(a) > len(b).
func LongerSlice[T any](a, b []T) bool {
	return len(a) > len(b)
}

// LongerMap returns len(a) > len(b).
func LongerMap[K comparable, V any](a, b map[K]V) bool {
	return len(a) > len(b)
}

// Not returns a func that returns !pred(t) when called with t.
func Not[T any](pred func(T) bool) func(T) bool {
	return func(t T) bool {
		return !pred(t)
	}
}

// NotEqual returns a != b.
func NotEqual[T comparable](a, b T) bool {
	return a != b
}

// Pointer returns &t.
func Pointer[T any](t T) *T {
	return &t
}

// ShorterSlice returns len(a) < len(b).
func ShorterSlice[T any](a, b []T) bool {
	return len(a) < len(b)
}

// ShorterMap returns len(a) < len(b).
func ShorterMap[K comparable, V any](a, b map[K]V) bool {
	return len(a) < len(b)
}

// SwapArgs returns a function that swaps the parameters of the specified
// function.
func SwapArgs[A, B, C any](f func(a A, b B) C) func(B, A) C {
	return func(b B, a A) C {
		return f(a, b)
	}
}

// True returns true, ignoring the input value.
func True[T any](T) bool {
	return true
}

// Zero returns the zero value for U, ignoring the input value.
func Zero[U, T any](T) U {
	var u U
	return u
}

// Drain consumes next and returns the number of elements consumed.
func Drain[T any](next Enumerator[T]) int {
	n := 0
	for t := next(); t.Valid(); t = next() {
		n++
	}
	return n
}
