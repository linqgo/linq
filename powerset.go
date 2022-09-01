package linq

import "math/bits"

func PowerSet[T any](q Query[T]) Query[Query[T]] {
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
	})
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
	})
}
