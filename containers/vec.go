package containers

import (
	"github.com/stntngo/rust-lib/iterator"
	"github.com/stntngo/rust-lib/option"
)

func Vec[T any](vec ...T) *VecIterator[T] {
	it := VecIterator[T](vec)
	return &it
}

type VecIterator[T any] []T

func (*VecIterator[T]) Collect(it iterator.Iterator[T]) *VecIterator[T] {
	var sink VecIterator[T]

	for opt := it.Next(); opt.Valid(); opt = it.Next() {
		sink = append(sink, opt.Unwrap())
	}

	return &sink
}

func (v *VecIterator[T]) Extend(item T) *VecIterator[T] {
	if v == nil {
		return Vec(item)
	}

	*v = append(*v, item)

	return v
}

func (v *VecIterator[T]) Next() option.Option[T] {
	if len(*v) == 0 {
		return option.None[T]()
	}

	var head T
	head, *v = (*v)[0], (*v)[1:]

	return option.Some(head)
}

func (v *VecIterator[T]) Clone() *VecIterator[T] {
	clone := make(VecIterator[T], len(*v))
	copy(clone, *v)

	return &clone
}
