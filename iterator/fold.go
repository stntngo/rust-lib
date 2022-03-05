package iterator

// Fold folds every element into an accumulator by applying
// an operation to the accumulated element and each
// element of the iterator and then returning that accumulated
// value.
func Fold[T, U any](it Iterator[T], initial U, fn func(U, T) U) U {
	accumulator := initial
	for opt := it.Next(); opt.Valid(); opt = it.Next() {
		accumulator = fn(accumulator, opt.Unwrap())
	}

	return accumulator
}
