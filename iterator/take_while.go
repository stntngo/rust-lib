package iterator

import "github.com/stntngo/rust-lib/option"

// TakeWhile creates an iterator that yields elements
// based on a predicate. After the first time that
// predicate returns false all future calls to
// the TakeWhileIterator will return None.
func TakeWhile[T any](it Iterator[T], pred func(T) bool) *TakeWhileIterator[T] {
	return &TakeWhileIterator[T]{
		it:   option.Some(it),
		pred: pred,
	}
}

type TakeWhileIterator[T any] struct {
	it   option.Option[Iterator[T]]
	pred func(T) bool
}

func (t *TakeWhileIterator[T]) Next() option.Option[T] {
	if !t.it.Valid() {
		return option.None[T]()
	}

	opt := t.it.Unwrap().Next()
	if !opt.Valid() {
		t.it = option.None[Iterator[T]]()

		return option.None[T]()
	}

	if t.pred(opt.Unwrap()) {
		return opt
	}

	t.it = option.None[Iterator[T]]()

	return option.None[T]()
}
