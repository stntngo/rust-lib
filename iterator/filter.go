package iterator

import "github.com/stntngo/rust-lib/option"

// Filter creates an iterator which uses a predicate function to
// deterimine if an iterator should be yielded.
func Filter[T any](it Iterator[T], pred func(T) bool) *FilterIterator[T] {
	return &FilterIterator[T]{
		it:   it,
		pred: pred,
	}
}

type FilterIterator[T any] struct {
	it   Iterator[T]
	pred func(T) bool
}

func (f *FilterIterator[T]) Next() option.Option[T] {
	for {
		next := f.it.Next()
		if !next.Valid() {
			return next
		}

		if f.pred(next.Unwrap()) {
			return next
		}
	}
}

func (f *FilterIterator[T]) Clone() Iterator[T] {
	original := f.it.(Clone[T, Iterator[T]])

	return &FilterIterator[T]{
		it:   original.Clone(),
		pred: f.pred,
	}

}
