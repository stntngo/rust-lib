package containers_test

import (
	"testing"

	"github.com/stntngo/rust-lib/containers"
	"github.com/stntngo/rust-lib/iterator"
	"github.com/stntngo/rust-lib/option"

	"github.com/stretchr/testify/assert"
)

func TestMaximumNestingDepth(t *testing.T) {
	var ints iterator.Iterator[int]
	ints = iterator.Map[rune](
		containers.String("(1+(2*3)+((8)/4))+1"),
		func(r rune) int {
			switch r {
			case '(':
				return 1
			case ')':
				return -1
			default:
				return 0
			}
		},
	)

	ints = iterator.Scan(
		ints,
		0,
		func(depth *int, current int) option.Option[int] {
			*depth += current

			return option.Some(*depth)
		},
	)

	max := iterator.Max(ints)
	assert.Equal(t, 3, max.UnwrapOr(0))
}
