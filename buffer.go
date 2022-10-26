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

// buffer supports buffering values before reading them. Each value read from a
// source must pass through a buffer of a given size before it can be read from
// the buffer. Once the source runs out of elements such that the buffer cannot
// be kept full, the buffer runs out. Elements remaining in the buffer may be
// consumed via the Enumerator method.
type buffer[T any] struct {
	next Enumerator[T]
	buf  ring[T]
}

// newBuffer creates a new buffer
func newBuffer[T any](next Enumerator[T], bufsize int) *buffer[T] {
	buf := newRing[T](bufsize)

	// Pre-fill the buffer.
	for i := 0; i < bufsize; i++ {
		t, ok := next().Get()
		if !ok {
			next = No[T]
			break
		}
		buf.Push(t)
	}

	return &buffer[T]{next: next, buf: buf}
}

// Next enumerates the buffer. It runs out when the source enumerator can't keep
// the buffer full. The residual elements remaining in the buffer may be
// accessed via Enumerator().
func (b *buffer[T]) Next() Maybe[T] {
	t2, ok := b.next().Get()
	if !ok {
		return No[T]()
	}
	t := b.buf.Pop()
	b.buf.Push(t2)
	return Some(t)
}

// Enumerator returns an enumerator for the elements in the buffer.
func (b *buffer[T]) Enumerator() Enumerator[T] {
	return b.buf.Enumerator()
}
