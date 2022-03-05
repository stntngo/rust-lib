package iterator

// ForEach calls a function on each element of an Iterator.
func ForEach[T any](it Iterator[T], fn func(T)) {
	Fold(it, nil, func(_ interface{}, item T) interface{} {
		fn(item)
		return nil
	})
}
