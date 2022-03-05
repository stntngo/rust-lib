package iterator_test

import (
	"testing"

	"github.com/stntngo/rust-lib/containers"
	"github.com/stntngo/rust-lib/iterator"
	"github.com/stretchr/testify/assert"
)

func Test_Flatten(t *testing.T) {
	var vec iterator.Iterator[*containers.VecIterator[int]]
	vec = containers.Vec(
		containers.Vec(1, 2, 3, 4),
		containers.Vec(5, 6),
	)

	flatten := iterator.Flatten[int](vec)

	out := iterator.Collect[int, *containers.VecIterator[int]](flatten)

	assert.Equal(t, containers.Vec(1, 2, 3, 4, 5, 6), out)
}
