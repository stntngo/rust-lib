package iterator

type Extender[T, U any] interface {
	Extend(T) U
}

// Partition consumes an iterator returning two collections from
// it according to whether the provided predicate function returned
// either true or false.
func Partition[T any, U Extender[T, U]](it Iterator[T], pred func(T) bool) (U, U) {
	var trues, falses U

	for opt := it.Next(); opt.Valid(); opt = it.Next() {
		item := opt.Unwrap()
		if pred(item) {
			trues = trues.Extend(item)
		} else {
			falses = falses.Extend(item)
		}
	}

	return trues, falses
}
