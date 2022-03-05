package iterator

import (
	"github.com/stntngo/rust-lib/option"
	"golang.org/x/exp/constraints"
)

// Max returns the maximum element of an iterator.
func Max[T constraints.Ordered](it Iterator[T]) option.Option[T] {
	max := it.Next()
	for next := it.Next(); next.Valid(); next = it.Next() {
		if max.Unwrap() < next.Unwrap() {
			max = next
		}
	}

	return max
}

// Min returns the minimum element of an iterator.
func Min[T constraints.Ordered](it Iterator[T]) option.Option[T] {
	min := it.Next()
	for next := it.Next(); next.Valid(); next = it.Next() {
		if min.Unwrap() > next.Unwrap() {
			min = next
		}
	}

	return min
}
