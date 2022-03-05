package iterator

import (
	"errors"

	"github.com/stntngo/rust-lib/option"
)

// Iterator is the main interface for dealing
// with iterable objects.
type Iterator[T any] interface {
	Next() option.Option[T]
}

type DoubleEndedIterator[T any] interface {
	Iterator[T]
	NextBack() option.Option[T]
}

// Count consumes the iterator, counting the number of iterations
// and returning it.
func Count[T any](it Iterator[T]) int {
	var count int

	for next := it.Next(); next.Valid(); next = it.Next() {
		count++
	}

	return count
}

// Last consumes the iterator, returning the last element.
func Last[T any](it Iterator[T]) option.Option[T] {
	last := it.Next()
	for item := it.Next(); item.Valid(); item = it.Next() {
		last = item
	}

	return last
}

// AdvanceBy advances the iterator by n elements.
func AdvanceBy[T any](it Iterator[T], n int) (int, error) {
	for i := 0; i < n; i++ {
		if !it.Next().Valid() {
			return i, errors.New("encountered None during AdvanceBy")
		}
	}

	return n, nil
}

// Nth returns the nth element of the iterator
func Nth[T any](it Iterator[T], n int) option.Option[T] {
	for i := 0; i < (n - 1); i++ {
		_ = it.Next()
	}

	return it.Next()
}

type Clone[T, U any] interface {
	Iterator[T]
	Clone() U
}
