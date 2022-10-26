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
