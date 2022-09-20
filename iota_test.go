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
		assertResultEqual(t, i, maybe(ie()))
	}

	assertOneShot(t, false, linq.Iota[int]())

	assertNoFastCount(t, linq.Iota[int]())
}

func TestIota12(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t, []int{}, linq.Iota1(0))
	assertQueryEqual(t, []int{0, 1, 2}, linq.Iota1(3))
	assertQueryEqual(t, []int{4, 5, 6}, linq.Iota2(4, 7))

	assertOneShot(t, false, linq.Iota1(10))
	assertOneShot(t, false, linq.Iota2(0, 10))

	assertFastCountEqual(t, 10, linq.Iota1(10))
	assertFastCountEqual(t, 10, linq.Iota2(0, 10))
}

func TestIota3(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t, []int{3, 5, 7}, linq.Iota3(3, 8, 2))
	assertQueryEqual(t, []int{8, 6, 4}, linq.Iota3(8, 3, -2))
	assertQueryEqual(t, []int{}, linq.Iota3(0, 0, 0))
	assert.Panics(t, func() { linq.Iota3(0, 1, 0) })

	assertOneShot(t, false, linq.Iota3(0, 10, 2))

	assertFastCountEqual(t, 5, linq.Iota3(0, 10, 2))
	assertFastCountEqual(t, 4, linq.Iota3(0, 10, 3))
}
