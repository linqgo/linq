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

type ring[T any] struct {
	buffer     []T
	head, size int
}

func newRing[T any](cap int) ring[T] {
	return ring[T]{
		buffer: make([]T, cap),
	}
}

func (r *ring[T]) Cap() int {
	return len(r.buffer)
}

func (r *ring[T]) Enumerator() Enumerator[T] {
	if r.size == 0 {
		return No[T]
	}
	tail := r.tail()
	if r.head < tail {
		// No wrap-around
		return sliceEnumerator(r.buffer[r.head:tail])
	}
	// Wrap-around
	return concatEnumerators(
		sliceEnumerator(r.buffer[r.head:]),
		sliceEnumerator(r.buffer[:tail]),
	)
}

func (r *ring[T]) Empty() bool {
	return r.size == 0
}

func (r *ring[T]) Full() bool {
	return r.size == r.Cap()
}

func (r *ring[T]) Push(t T) {
	if r.Full() {
		panic("buffer full")
	}
	r.buffer[r.tail()] = t
	r.size++
}

func (r *ring[T]) Pop() T {
	if r.Empty() {
		panic("buffer empty")
	}
	t := r.buffer[r.head]
	r.head = (r.head + 1) % r.Cap()
	r.size--
	return t
}

func (r *ring[T]) tail() int {
	return (r.head + r.size) % r.Cap()
}
