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

package ring

type Enumerator[T any] func() (T, bool)

func No[T any]() (T, bool) {
	var zero T
	return zero, false
}

type Ring[T any] struct {
	buffer     []T
	head, size int
}

func New[T any](cap int) Ring[T] {
	return Ring[T]{
		buffer: make([]T, cap),
	}
}

func (r *Ring[T]) Cap() int {
	return len(r.buffer)
}

func (r *Ring[T]) Enumerator() func() (T, bool) {
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

func (r *Ring[T]) Empty() bool {
	return r.size == 0
}

func (r *Ring[T]) Head() T {
	if r.Empty() {
		panic("buffer empty")
	}
	return r.buffer[r.head]
}

func (r *Ring[T]) Push(t T) {
	if r.full() {
		if r.head == 0 {
			r.buffer = append(r.buffer, t)
			r.size++
			return
		} else {
			buf := make([]T, 2*len(r.buffer))
			copy(buf, r.buffer[r.head:])
			copy(buf[r.Cap()-r.head:], r.buffer[:r.head])
			r.buffer = buf
			r.head = 0
		}
	}
	r.buffer[r.tail()] = t
	r.size++
}

func (r *Ring[T]) Pop() T {
	t := r.Head()
	r.head = (r.head + 1) % r.Cap()
	r.size--
	return t
}

func (r *Ring[T]) full() bool {
	return r.size == r.Cap()
}

func (r *Ring[T]) tail() int {
	return (r.head + r.size) % r.Cap()
}

func sliceEnumerator[T any](s []T) func() (T, bool) {
	i := 0
	return func() (T, bool) {
		if i == len(s) {
			var zero T
			return zero, false
		}
		t := s[i]
		i++
		return t, true
	}
}

func concatEnumerators[T any](nexts ...Enumerator[T]) Enumerator[T] {
	next := No[T]
	return func() (T, bool) {
		for {
			if t, ok := next(); ok {
				return t, true
			}
			if len(nexts) == 0 {
				return No[T]()
			}
			next, nexts = nexts[0], nexts[1:]
		}
	}
}
