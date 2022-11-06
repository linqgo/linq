package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/linqgo/linq"
)

func TestScan(t *testing.T) {
	t.Parallel()

	i := 0
	scan := linq.From(1, 2, 3, 4, 5).Scan()
	var v int
	for scan(&v) {
		i++
		if !assert.Equal(t, i, v) {
			break
		}
	}
}
