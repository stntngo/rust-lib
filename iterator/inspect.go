package iterator

import "github.com/stntngo/rust-lib/option"

// Inspect does something with each element of an iterator
// and then passes the value on.
func Inspect[T any](it Iterator[T], fn func(T)) *InspectIterator[T] {
	return &InspectIterator[T]{
		it: it,
		fn: fn,
	}
}

type InspectIterator[T any] struct {
	it Iterator[T]
	fn func(T)
}

func (i *InspectIterator[T]) Next() option.Option[T] {
	item := i.it.Next()
	if item.Valid() {
		i.fn(item.Unwrap())
	}

	return item
}
