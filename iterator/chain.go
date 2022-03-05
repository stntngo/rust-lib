package iterator

import "github.com/stntngo/rust-lib/option"

// Chain will return a new iterator which will first
// iterate over values from the first iterator and
// then iterate over values from the second iterator.
func Chain[T any](first, second Iterator[T]) *ChainIterator[T] {
	return &ChainIterator[T]{
		first:  option.Some(first),
		second: option.Some(second),
	}
}

type ChainIterator[T any] struct {
	first, second option.Option[Iterator[T]]
}

func (c *ChainIterator[T]) Next() option.Option[T] {
	if c.first.Valid() {
		item := c.first.Unwrap().Next()
		if item.Valid() {
			return item
		}

		c.first = option.None[Iterator[T]]()
	}

	if c.second.Valid() {
		item := c.second.Unwrap().Next()
		if item.Valid() {
			return item
		}

		c.second = option.None[Iterator[T]]()
	}

	return option.None[T]()
}
