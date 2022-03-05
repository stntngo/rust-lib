package iterator

import "github.com/stntngo/rust-lib/option"

// Creates an iterator that supports a `Peek` method
// that lets callers look at the next element of the
// iterator without consuming it.
func Peekable[T any](it Iterator[T]) *PeekableIterator[T] {
	return &PeekableIterator[T]{
		it:   it,
		next: it.Next(),
	}
}

type PeekableIterator[T any] struct {
	it   Iterator[T]
	next option.Option[T]
}

func (p *PeekableIterator[T]) Peek() option.Option[T] {
	return p.next
}

func (p *PeekableIterator[T]) Next() option.Option[T] {
	defer func() {
		p.next = p.it.Next()
	}()

	return p.next
}
