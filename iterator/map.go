package iterator

import "github.com/stntngo/rust-lib/option"

// Map creates a new iterator which calls a provided
// a function on each element of the original
// iterator
func Map[T, U any](it Iterator[T], fn func(T) U) *MapIterator[T, U] {
	return &MapIterator[T, U]{
		it: it,
		fn: fn,
	}
}

type MapIterator[T, U any] struct {
	it Iterator[T]
	fn func(T) U
}

func (m *MapIterator[T, U]) Next() option.Option[U] {
	return option.Map(m.it.Next(), m.fn)
}
