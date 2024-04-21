// Copyright 2022-2024 Marcelo Cantos
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
	"bufio"
	"io"
	"iter"
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
	return FromSeq(
		readerSeq(func() ([]byte, error) {
			switch {
			case s.Scan():
				return s.Bytes(), nil
			case s.Err() == nil:
				return nil, io.EOF
			default:
				return nil, s.Err()
			}
		}),
		OneShotOption[[]byte](true),
	)
}

// FromScannerString reads a query containing string tokens read from s. Hint: use
// (Scanner).Split() to control how the input stream is tokenized.
func FromScannerString(s *bufio.Scanner) Query[string] {
	return FromSeq(
		readerSeq(func() (string, error) {
			switch {
			case s.Scan():
				return s.Text(), nil
			case s.Err() == nil:
				return "", io.EOF
			default:
				return "", s.Err()
			}
		}),
		OneShotOption[string](true),
	)
}

func readerQuery[T any](read func() (T, error)) Query[T] {
	return FromSeq(readerSeq(read), OneShotOption[T](true))
}

func readerSeq[T any](read func() (T, error)) iter.Seq[T] {
	return func(yield func(t T) bool) {
		for {
			c, err := read()
			if err != nil {
				if err != io.EOF {
					panic(err)
				}
				return
			}
			if !yield(c) {
				return
			}
		}
	}
}
