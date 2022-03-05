package containers

import (
	"unicode/utf8"

	"github.com/stntngo/rust-lib/iterator"
	"github.com/stntngo/rust-lib/option"
)

type StringLike interface {
	~string | ~rune | ~byte
}

func String(s string) *StringIterator {
	it := StringIterator(s)
	return &it
}

type StringIterator string

func (s *StringIterator) Next() option.Option[rune] {
	r, w := utf8.DecodeRuneInString(string(*s))

	if r == utf8.RuneError {
		return option.None[rune]()
	}

	*s = (*s)[w:]

	return option.Some(r)
}

func (*StringIterator) Collect(it iterator.Iterator[string]) *StringIterator {
	var s StringIterator
	for opt := it.Next(); opt.Valid(); opt = it.Next() {
		s += StringIterator(opt.Unwrap())
	}

	return &s
}
