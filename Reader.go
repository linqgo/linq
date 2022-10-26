package linq

import (
	"bufio"
	"io"
)

// FromByteReader returns a query containing bytes read from r. Hint: use a
// bufio.Reader to wrap an io.Reader in an io.ByteReader.
//
// The returned query is not replayable. Use (Query).Memoize() if you need a
// replayable query.
func FromByteReader(r io.ByteReader) Query[byte] {
	return readerQuery(r.ReadByte)
}

// FromRuneReader returns a query containing runes read from r. Hint: use a
// bufio.Reader to wrap an io.Reader in an io.RuneReader.
//
// The returned query is not replayable. Use (Query).Memoize() if you need a
// replayable query.
func FromRuneReader(r io.RuneReader) Query[rune] {
	return readerQuery(func() (rune, error) {
		c, _, err := r.ReadRune()
		return c, err
	})
}

// FromScanner reads a query containing []byte tokens read from s. Hint: use
// (Scanner).Split() to control how the input stream is tokenized.
func FromScanner(s *bufio.Scanner) Query[[]byte] {
	return NewQuery(func() Enumerator[[]byte] {
		return readerEnumerator(func() ([]byte, error) {
			switch {
			case s.Scan():
				return s.Bytes(), nil
			case s.Err() == nil:
				return nil, io.EOF
			default:
				return nil, s.Err()
			}
		})
	}, OneShotOption[[]byte](true))
}

// FromScannerString reads a query containing string tokens read from s. Hint: use
// (Scanner).Split() to control how the input stream is tokenized.
func FromScannerString(r *bufio.Scanner) Query[string] {
	return NewQuery(func() Enumerator[string] {
		return readerEnumerator(func() (string, error) {
			switch {
			case r.Scan():
				return r.Text(), nil
			case r.Err() == nil:
				return "", io.EOF
			default:
				return "", r.Err()
			}
		})
	}, OneShotOption[string](true))
}

func readerQuery[T any](read func() (T, error)) Query[T] {
	return NewQuery(func() Enumerator[T] {
		return readerEnumerator(read)
	}, OneShotOption[T](true))
}

func readerEnumerator[T any](read func() (T, error)) Enumerator[T] {
	return func() Maybe[T] {
		c, err := read()
		if err == nil {
			return Some(c)
		}
		if err != io.EOF {
			panic(err)
		}
		return No[T]()
	}
}
