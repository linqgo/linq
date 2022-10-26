package linq

// KV represents a key/value pair.
type KV[K, V any] struct {
	Key   K
	Value V
}

// NewKV returns a new KV.
func NewKV[K, V any](key K, value V) KV[K, V] {
	return KV[K, V]{
		Key:   key,
		Value: value,
	}
}

func (kv KV[K, V]) KV() (K, V) {
	return kv.Key, kv.Value
}

func Key[KeyVal KV[K, V], K, V any](kv KV[K, V]) K {
	return kv.Key
}

func Value[KeyVal KV[K, V], K, V any](kv KV[K, V]) V {
	return kv.Value
}
