package iterator

import "github.com/stntngo/rust-lib/option"

// Range creates a half-open range bounded inclusively
// below and exclusively above `(start -> end]` that
// when iterated over is stepped over by `step` each
// turn.
func Range(start, stop, step int) *RangeIterator {
	return &RangeIterator{
		current: start,
		target:  stop,
		step:    step,
	}
}

type RangeIterator struct {
	current, target, step int
}

func (r *RangeIterator) Next() option.Option[int] {
	if r.current >= r.target {
		return option.None[int]()
	}

	value := r.current
	r.current += r.step

	return option.Some(value)
}

func (r *RangeIterator) Clone() Iterator[int] {
	return &RangeIterator{
		current: r.current,
		target:  r.target,
		step:    r.step,
	}
}
