package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/linqgo/linq"
)

func TestIota(t *testing.T) {
	t.Parallel()

	ie := linq.Iota[int]().Enumerator()
	for i := 0; i < 10; i++ {
		assertSome(t, i, ie())
	}

	assertOneShot(t, false, linq.Iota[int]())

	assertNo(t, linq.Iota[int]().FastCount())
}

func TestIota12(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t, []int{}, linq.Iota1(0))
	assertQueryEqual(t, []int{0, 1, 2}, linq.Iota1(3))
	assertQueryEqual(t, []int{4, 5, 6}, linq.Iota2(4, 7))

	assertOneShot(t, false, linq.Iota1(10))
	assertOneShot(t, false, linq.Iota2(0, 10))

	assertSome(t, 10, linq.Iota1(10).FastCount())
	assertSome(t, 10, linq.Iota2(0, 10).FastCount())
}

func TestIota3(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t, []int{3, 5, 7}, linq.Iota3(3, 8, 2))
	assertQueryEqual(t, []int{8, 6, 4}, linq.Iota3(8, 3, -2))
	assertQueryEqual(t, []int{}, linq.Iota3(0, 0, 0))
	assert.Panics(t, func() { linq.Iota3(0, 1, 0) })

	assertOneShot(t, false, linq.Iota3(0, 10, 2))

	assertSome(t, 5, linq.Iota3(0, 10, 2).FastCount())
	assertSome(t, 4, linq.Iota3(0, 10, 3).FastCount())
}

func TestIotaFastElementAt(t *testing.T) {
	t.Parallel()

	q := linq.Iota[int]()
	assertSome(t, 0, q.FastElementAt(0))
	assertSome(t, 3, q.FastElementAt(3))
	assertSome(t, 9999, q.FastElementAt(9999))
	assertNo(t, q.FastElementAt(-1))
}

func TestIota3FastElementAt(t *testing.T) {
	t.Parallel()

	q := linq.Iota3(10, 20, 3)
	assertSome(t, 10, q.FastElementAt(0))
	assertSome(t, 19, q.FastElementAt(3))
	assertNo(t, q.FastElementAt(4))
	assertNo(t, q.FastElementAt(-1))
}

func TestIota3BackwardsFastElementAt(t *testing.T) {
	t.Parallel()

	q := linq.Iota3(20, 10, -3)
	assertSome(t, 20, q.FastElementAt(0))
	assertSome(t, 11, q.FastElementAt(3))
	assertNo(t, q.FastElementAt(4))
	assertNo(t, q.FastElementAt(-1))
}
