package containers

import (
	"github.com/stntngo/rust-lib/iterator"
	"github.com/stntngo/rust-lib/option"
)

type HashMap[K comparable, V any] map[K]V

func (h *HashMap[K, V]) Next() option.Option[iterator.Pair[K, V]] {
	for k, v := range *h {
		delete(*h, k)

		return option.Some(iterator.Pair[K, V]{
			First:  k,
			Second: v,
		})
	}

	return option.None[iterator.Pair[K, V]]()
}

func (h *HashMap[K, V]) Clone() *HashMap[K, V] {
	clone := make(HashMap[K, V], len(*h))

	for k, v := range *h {
		clone[k] = v
	}

	return &clone
}

func (h *HashMap[K, V]) Collect(it iterator.Iterator[iterator.Pair[K, V]]) {
	for opt := it.Next(); opt.Valid(); opt = it.Next() {
		pair := opt.Unwrap()

		(*h)[pair.First] = pair.Second
	}
}
