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

// Longer returns true if and only if q has more elements than r.
func (q Query[T]) Longer(r Query[T]) bool {
	return Shorter(r, q)
}

// Shorter returns true if and only if q has fewer elements than r.
func (q Query[T]) Shorter(r Query[T]) bool {
	return Shorter(q, r)
}

// FastLonger returns true if and only if a has more elements than b and this
// can be determined in O(1) time, otherwise returns ok = false.
func (a Query[T]) FastLonger(b Query[T]) Maybe[bool] {
	return FastLonger(a, b)
}

// FastShorter returns true if and only if a has fewer elements than b and this
// can be determined in O(1) time, otherwise returns ok = false.
func (a Query[T]) FastShorter(b Query[T]) Maybe[bool] {
	return FastShorter(a, b)
}

// FastLonger returns true if and only if a has more elements than b and this
// can be determined in O(1) time, otherwise returns ok = false.
func FastLonger[A, B any](a Query[A], b Query[B]) Maybe[bool] {
	return FastShorter(b, a)
}

// FastShorter returns true if and only if a has fewer elements than b and this
// can be determined in O(1) time, otherwise returns ok = false.
func FastShorter[A, B any](a Query[A], b Query[B]) Maybe[bool] {
	diff, ok := fastLenDiff(a, b).Get()
	return NewMaybe(diff < 0, ok)
}

// Longer returns true if and only if a has more elements than b.
func Longer[A, B any](a Query[A], b Query[B]) bool {
	return Shorter(b, a)
}

// Shorter returns true if and only if a has fewer elements than b.
func Shorter[A, B any](a Query[A], b Query[B]) bool {
	if shorter, ok := FastShorter(a, b).Get(); ok {
		return shorter
	}

	var aok, bok bool
	Drain(zipEnumerator(a.Enumerator(), b.Enumerator(), &aok, &bok))
	return !aok && bok
}

func fastLenDiff[A, B any](a Query[A], b Query[B]) Maybe[int] {
	alen, alenok := a.FastCount().Get()
	blen, blenok := b.FastCount().Get()
	return NewMaybe(alen-blen, alenok && blenok)
}
