package linq

// KV represents a key/value pair.
type KV[K comparable, V any] struct {
	Key   K
	Value V
}

// NewKV returns a new KV.
func NewKV[K comparable, V any](key K, value V) KV[K, V] {
	return KV[K, V]{Key: key, Value: value}
}
