package iterator

import "github.com/stntngo/rust-lib/option"

// Flatten creates an iterator that flattens nested structure.
//
// NOTE: Flatten only removes one level of nesting at a time
// and does not perfom "deep" flattening.
func Flatten[T any, U Iterator[T]](it Iterator[U]) *FlattenIterator[T, U] {
	return &FlattenIterator[T, U]{
		it: Peekable(it),
	}
}

type FlattenIterator[T any, U Iterator[T]] struct {
	it *PeekableIterator[U]
}

func (f *FlattenIterator[T, U]) Next() option.Option[T] {
	for {
		if !f.it.Peek().Valid() {
			return option.None[T]()
		}

		item := f.it.Peek().Unwrap().Next()
		if item.Valid() {
			return item
		}

		_ = f.it.Next()
	}
}
