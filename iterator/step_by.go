package iterator

import "github.com/stntngo/rust-lib/option"

// StepBy creates an iterator starting at the same point, but stepping
// by the given amount at each iteration.
func StepBy[T any](it Iterator[T], step int) *StepByIterator[T] {
	return &StepByIterator[T]{
		it:    it,
		step:  step,
		taken: false,
	}
}

type StepByIterator[T any] struct {
	it    Iterator[T]
	step  int
	taken bool
}

func (s *StepByIterator[T]) Next() option.Option[T] {
	if !s.taken {
		s.taken = true

		return s.it.Next()
	}

	return Nth(s.it, s.step-1)
}
