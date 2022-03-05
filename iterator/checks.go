package iterator

// All tests if every element of the iterator matches a
// predicate.
//
// An empty iterator returns true.
func All[T any](it Iterator[T], pred func(T) bool) bool {
	for item := it.Next(); item.Valid(); item = it.Next() {
		if !pred(item.Unwrap()) {
			return false
		}
	}

	return true
}

// Any tests if any element of the iterator matches a
// predicate.
//
// An empty iterator returns false.
func Any[T any](it Iterator[T], pred func(T) bool) bool {
	for item := it.Next(); item.Valid(); item = it.Next() {
		if pred(item.Unwrap()) {
			return true
		}
	}

	return false
}
