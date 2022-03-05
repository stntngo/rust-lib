package iterator

import "github.com/stntngo/rust-lib/option"

// Zip returns an iterator that "zips up" two iterators
// into a single iterator of pairs.
func Zip[T, U any](left Iterator[T], right Iterator[U]) *ZipIterator[T, U] {
	return &ZipIterator[T, U]{
		left:  left,
		right: right,
	}
}

type ZipIterator[T, U any] struct {
	left  Iterator[T]
	right Iterator[U]
}

func (z *ZipIterator[T, U]) Next() option.Option[Pair[T, U]] {
	left := z.left.Next()
	if !left.Valid() {
		return option.None[Pair[T, U]]()
	}

	right := z.right.Next()
	if !right.Valid() {
		return option.None[Pair[T, U]]()
	}

	return option.Some(Pair[T, U]{
		First:  left.Unwrap(),
		Second: right.Unwrap(),
	})
}
