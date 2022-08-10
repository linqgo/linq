package linq_test

import (
	"fmt"
	"testing"

	"github.com/marcelocantos/linq"
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

	assertQueryEqual(t,
		[]string{"A1", "B2", "C3"},
		linq.Zip(
			linq.From("A", "B", "C", "D"),
			linq.From(1, 2, 3),
			func(a string, b int) string {
				return fmt.Sprintf("%s%d", a, b)
			},
		),
	)
}
