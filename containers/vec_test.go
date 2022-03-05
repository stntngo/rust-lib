package containers_test

import (
	"strings"
	"testing"

	"github.com/stntngo/rust-lib/containers"
	"github.com/stntngo/rust-lib/iterator"
	"github.com/stretchr/testify/assert"
)

func Test_VecFilter(t *testing.T) {

	out := iterator.Collect[int, *containers.VecIterator[int]](
		iterator.Filter[int](
			containers.Vec(1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
			func(i int) bool {
				return i%3 == 0
			},
		),
	)

	assert.Equal(t, out, containers.Vec(3, 6, 9))
}

func Test_VecMap(t *testing.T) {
	var it iterator.Iterator[string]
	it = containers.Vec("hello", "world", "how", "are", "you")

	it = iterator.Map(it, strings.ToUpper)
	it = iterator.Intersperse(it, " ")

	s := iterator.Collect[string, *containers.StringIterator](it)

	assert.Equal(t, "HELLO WORLD HOW ARE YOU", string(*s))
}

func Test_VecFold(t *testing.T) {
	var it iterator.Iterator[int]
	it = containers.Vec(1, 2, 3, 4)

	assert.Equal(
		t,
		10,
		iterator.Fold(
			it,
			0,
			func(acc int, x int) int {
				return acc + x
			},
		),
	)

	it = containers.Vec(1, 2, 3, 4)

	assert.Equal(
		t,
		30,
		iterator.Fold(
			it,
			20,
			func(acc int, x int) int {
				return acc + x
			},
		),
	)
}
