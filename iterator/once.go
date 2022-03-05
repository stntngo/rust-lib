package iterator

import "github.com/stntngo/rust-lib/option"

// Once creates an iterator that yields an element
// exactly once.
func Once[T any](item T) *OnceIterator[T] {
	return &OnceIterator[T]{
		done: false,
		item: item,
	}
}

type OnceIterator[T any] struct {
	done bool
	item T
}

func (o *OnceIterator[T]) Next() option.Option[T] {
	if !o.done {
		o.done = true
		return option.Some(o.item)
	}

	return option.None[T]()
}
