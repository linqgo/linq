package linq

import (
	"math/bits"
	"unsafe"
)

func PowerSet[T any](q Query[T]) Query[Query[T]] {
	// Calculate the number of bits available in a positive int (minus one
	// because the high bit is reserved for negative ints).
	const positiveIntBits = int(8*unsafe.Sizeof(int(0))) - 1

	// Slight problem: If FastCount(q) >= 64, then the actual count can't be
	// represented with an int.
	count := -1
	if c, ok := FastCount(q).Get(); ok && count < positiveIntBits {
		count = 1 << c
	}

	return Pipe(q, func(next Enumerator[T]) Enumerator[Query[T]] {
		var cache []T
		var mask uint64 = 0
		mask--
		return func() Maybe[Query[T]] {
			mask++
			if mask > 0 && mask&(mask-1) == 0 {
				// New bit
				t, ok := next().Get()
				if !ok {
					mask--
					return No[Query[T]]()
				}
				cache = append(cache, t)
			}
			return Some(powerSubSet(cache, mask))
		}
	}, FastCountOption[Query[T]](count))
}

func powerSubSet[T any](cache []T, mask uint64) Query[T] {
	return NewQuery(func() Enumerator[T] {
		return func() Maybe[T] {
			if mask == 0 {
				return No[T]()
			}
			i := bits.TrailingZeros64(mask)
			mask &= mask - 1
			return Some(cache[i])
		}
	}, FastCountOption[T](bits.OnesCount64(mask)))
}
