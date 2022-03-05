package iterator

import "github.com/stntngo/rust-lib/option"

// Cycle repeats an iterator endlessly
func Cycle[T any, U Clone[T, U]](original U) *CycleIterator[T, U] {
	return &CycleIterator[T, U]{
		original: original,
		current:  original.Clone(),
	}
}

type CycleIterator[T any, U Clone[T, U]] struct {
	original Clone[T, U]
	current  Iterator[T]
}

func (c *CycleIterator[T, U]) Next() option.Option[T] {
	next := c.current.Next()

	if !next.Valid() {
		c.current = c.original.Clone()

		next = c.Next()
	}

	return next
}
