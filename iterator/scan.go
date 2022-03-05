package iterator

import "github.com/stntngo/rust-lib/option"

// An iterator adapter similar to `Fold` that holds internal
// state and produces a new iterator.
//
// Scan takes two arguments in addition to the underlying
// iterator: an initital value which seeds the internal state
// and a function that takes two arguments with the first
// being a pointer to the internal state and the second
// an iterator element.
//
// On iteration, the function is applied to each element of the
// iterator and the return value from the function is yielded
// by the ScanIterator.
func Scan[T, U, V any](it Iterator[T], state U, fn func(*U, T) option.Option[V]) *ScanIterator[T, U, V] {
	return &ScanIterator[T, U, V]{
		it:    it,
		state: state,
		fn:    fn,
	}
}

type ScanIterator[T, U, V any] struct {
	it    Iterator[T]
	state U
	fn    func(*U, T) option.Option[V]
}

func (s *ScanIterator[T, U, V]) Next() option.Option[V] {
	next := s.it.Next()
	if !next.Valid() {
		return option.None[V]()
	}

	return s.fn(&s.state, next.Unwrap())
}
