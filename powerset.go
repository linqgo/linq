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
	c, ok := FastCount(q)
	if ok && count < positiveIntBits {
		count = 1 << c
	}

	return Pipe(q, func(next Enumerator[T]) Enumerator[Query[T]] {
		var cache []T
		var mask uint64 = 0
		mask--
		return func() (q Query[T], ok bool) {
			mask++
			if mask > 0 && mask&(mask-1) == 0 {
				// New bit
				t, ok := next()
				if !ok {
					mask--
					return q, false
				}
				cache = append(cache, t)
			}
			return powerSubSet(cache, mask), true
		}
	}, FastCountOption[Query[T]](count))
}

func powerSubSet[T any](cache []T, mask uint64) Query[T] {
	return NewQuery(func() Enumerator[T] {
		return func() (t T, ok bool) {
			if mask == 0 {
				return t, false
			}
			i := bits.TrailingZeros64(mask)
			mask &= mask - 1
			return cache[i], true
		}
	}, FastCountOption[T](bits.OnesCount64(mask)))
}
