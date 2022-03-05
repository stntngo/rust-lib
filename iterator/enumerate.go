package iterator

import "github.com/stntngo/rust-lib/option"

func Enumerate[T any](it Iterator[T]) *EnumerateIterator[T] {
	return &EnumerateIterator[T]{
		it:    it,
		index: -1,
	}
}

type EnumerateIterator[T any] struct {
	it    Iterator[T]
	index int
}

func (e *EnumerateIterator[T]) Next() option.Option[Pair[int, T]] {
	opt := e.it.Next()
	if !opt.Valid() {
		return option.None[Pair[int, T]]()
	}

	e.index++
	return option.Some(Pair[int, T]{
		First:  e.index,
		Second: opt.Unwrap(),
	})
}
