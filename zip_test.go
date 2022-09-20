package linq_test

import (
	"fmt"
	"testing"

	"github.com/linqgo/linq"
)

func TestZip(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t,
		[]string{"A1", "B2", "C3"},
		linq.Zip(
			linq.From("A", "B", "C"),
			linq.From(1, 2, 3, 4),
			func(a string, b int) string {
				return fmt.Sprintf("%s%d", a, b)
			},
		),
	)

	q := linq.Zip(
		linq.From("A", "B", "C", "D"),
		linq.From(1, 2, 3),
		func(a string, b int) string {
			return fmt.Sprintf("%s%d", a, b)
		},
	)
	assertQueryEqual(t, []string{"A1", "B2", "C3"}, q)

	assertOneShot(t, false, q)
	assertOneShot(t, true, linq.Zip(
		oneshot(),
		linq.From(1, 2, 3),
		func(a, b int) int { return a + b },
	))
	assertOneShot(t, true, linq.Zip(
		linq.From(1, 2, 3),
		oneshot(),
		func(a, b int) int { return a + b },
	))
}

func TestZipKV(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t,
		[]linq.KV[string, int]{{"A", 1}, {"B", 2}, {"C", 3}},
		linq.ZipKV(linq.From("A", "B", "C"), linq.From(1, 2, 3, 4)),
	)
}

func TestUnzip(t *testing.T) {
	t.Parallel()

	a, b := linq.Unzip(linq.Iota1(10), func(i int) (int, int) { return i / 3, i % 3 })

	assertQueryEqual(t, []int{0, 0, 0, 1, 1, 1, 2, 2, 2, 3}, a)
	assertQueryEqual(t, []int{0, 1, 2, 0, 1, 2, 0, 1, 2, 0}, b)
}

func TestUnzipKV(t *testing.T) {
	t.Parallel()

	k, v := linq.UnzipKV(linq.FromMap(map[string]int{"A": 1, "B": 2, "C": 3}))

	assertQueryElementsMatch(t, []string{"A", "B", "C"}, k)
	assertQueryElementsMatch(t, []int{1, 2, 3}, v)
}

func TestUnzipKVBuffered(t *testing.T) {
	t.Parallel()

	c := make(chan linq.KV[string, int], 3)
	c <- linq.NewKV("A", 1)
	c <- linq.NewKV("B", 2)
	c <- linq.NewKV("C", 3)
	close(c)
	k, v := linq.UnzipKV(linq.FromChannel(c))

	assertQueryElementsMatch(t, []string{"A", "B", "C"}, k)
	assertQueryElementsMatch(t, []int{1, 2, 3}, v)
	assertQueryElementsMatch(t, []string{"A", "B", "C"}, k)
	assertQueryElementsMatch(t, []int{1, 2, 3}, v)
}
