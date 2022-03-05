package iterator

import (
	"sync"

	"github.com/stntngo/rust-lib/option"
)

// SkipWhile creates an iterator that skips elements so
// long as an iterator returns true.
func SkipWhile[T any](it Iterator[T], pred func(T) bool) *SkipWhileIterator[T] {
	return &SkipWhileIterator[T]{
		it: Peekable(it),
		pred: func(opt option.Option[T]) bool {
			if opt.Valid() {
				return pred(opt.Unwrap())
			}

			return false
		},
	}
}

type SkipWhileIterator[T any] struct {
	it   *PeekableIterator[T]
	pred func(option.Option[T]) bool

	once sync.Once
}

func (s *SkipWhileIterator[T]) Next() option.Option[T] {
	s.once.Do(func() {
		for s.pred(s.it.Peek()) {
			_ = s.it.Next()
		}

	})

	return s.it.Next()
}
