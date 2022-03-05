package option

import "fmt"

func Some[T any](value T) Option[T] {
	return Option[T]{
		value: value,
		valid: true,
	}
}

func None[T any]() Option[T] {
	return Option[T]{}
}

type Option[T any] struct {
	value T
	valid bool
}

func (o Option[T]) String() string {
	if !o.valid {
		return "None"
	}

	return fmt.Sprintf("Some[%v]", o.value)
}

func (o Option[T]) Valid() bool {
	return o.valid
}

func (o Option[T]) Unwrap() T {
	if !o.Valid() {
		panic("called `Option.Unwrap` on a `None` value")
	}

	return o.value
}

func (o Option[T]) UnwrapOr(alt T) T {
	if o.Valid() {
		return o.Unwrap()
	}

	return alt
}

func (o Option[T]) UnwrapOrError(err error) (T, error) {
	if o.Valid() {
		return o.Unwrap(), nil
	}

	var empty T
	return empty, err
}

func Map[T, U any](opt Option[T], fn func(T) U) Option[U] {
	if !opt.Valid() {
		return None[U]()
	}

	return Some(fn(opt.Unwrap()))
}
