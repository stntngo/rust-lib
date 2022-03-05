package iterator

type Collector[T, U any] interface {
	Collect(Iterator[T]) U
}

// Collect transforms an iterator into a Collection through
// the Collector interface.
func Collect[T any, U Collector[T, U]](it Iterator[T]) U {
	var collector U
	return collector.Collect(it)
}
