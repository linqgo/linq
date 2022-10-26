package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/linqgo/linq"
)

func TestRandomSample(t *testing.T) {
	t.Parallel()

	assert.InEpsilon(t, 1000, linq.Iota1(10000).Sample(0.1).Count(), 1.1)

	assertQueryEqual(t, []int{1, 3, 4, 5, 6, 9}, linq.Iota1(10).SampleSeed(0.6, 0))
	assertQueryEqual(t, []int{0, 8, 9, 12}, linq.Iota1(20).SampleSeed(0.3, 1234))

	assertOneShot(t, false, linq.Iota1(10000).Sample(0.1))
	assertOneShot(t, true, oneshot().Sample(0.1))
	assertOneShot(t, false, linq.Iota1(10).SampleSeed(0.6, 0))
	assertOneShot(t, true, oneshot().SampleSeed(0.6, 0))
}
