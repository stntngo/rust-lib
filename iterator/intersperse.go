package iterator

import "github.com/stntngo/rust-lib/option"

// Intersperse creates an iterator which places a copy of
// `separator` between adjacent items of the original
// iterator.
func Intersperse[T any](it Iterator[T], separator T) *IntersperseIterator[T] {
	return &IntersperseIterator[T]{
		it:   Peekable(it),
		item: separator,
		odd:  false,
	}
}

type IntersperseIterator[T any] struct {
	it   *PeekableIterator[T]
	item T

	odd bool
}

func (i *IntersperseIterator[T]) Next() option.Option[T] {
	defer func() {
		i.odd = !i.odd
	}()

	if i.odd && i.it.Peek().Valid() {
		return option.Some(i.item)
	}

	return i.it.Next()
}

// Intersperse creates an iterator which places a the
// result of calling `separator` between adjacent items
// of the original iterator.
func IntersperseWith[T any](it Iterator[T], seperator func() T) *IntersperseWithIterator[T] {
	return &IntersperseWithIterator[T]{
		it:        Peekable(it),
		seperator: seperator,
		odd:       false,
	}
}

type IntersperseWithIterator[T any] struct {
	it        *PeekableIterator[T]
	seperator func() T

	odd bool
}

func (i *IntersperseWithIterator[T]) Next() option.Option[T] {
	defer func() {
		i.odd = !i.odd
	}()

	if i.odd && i.it.Peek().Valid() {
		return option.Some(i.seperator())
	}

	return i.it.Next()
}
