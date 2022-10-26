package linq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKVKV(t *testing.T) {
	t.Parallel()

	k, v := NewKV(42, "hello").KV()
	assert.Equal(t, 42, k)
	assert.Equal(t, "hello", v)
}
