package linq_test

import (
	"testing"

	"github.com/linqgo/linq"
)

func TestElementAt(t *testing.T) {
	t.Parallel()

	assertNo(t, linq.From[int]().ElementAt(0))
	assertNo(t, linq.Iota2(1, 6).ElementAt(42))
	assertSome(t, 3, linq.Iota2(1, 6).ElementAt(2))
}

func TestFastElementAt(t *testing.T) {
	t.Parallel()

	assertSome(t, 4, linq.Iota2(1, 6).FastElementAt(3))
	assertNo(t, linq.Iota2(1, 6).FastElementAt(6))
	assertNo(t, linq.Iota2(1, 6).FastElementAt(-1))
}

func TestFirst(t *testing.T) {
	t.Parallel()

	assertNo(t, linq.Iota1(0).First())
	assertSome(t, 1, linq.Iota2(1, 6).First())
}

func TestLast(t *testing.T) {
	t.Parallel()

	assertNo(t, linq.Iota1(0).Last())
	assertSome(t, 5, linq.Iota2(1, 6).Last())
	assertSome(t, 3, oneshotN(1, 2, 3).Last())
}
