package iterator_test

import (
	"testing"

	"github.com/stntngo/rust-lib/containers"
	"github.com/stntngo/rust-lib/iterator"
	"github.com/stretchr/testify/assert"
)

func Test_Partition(t *testing.T) {
	first, second := iterator.Partition[int, *containers.VecIterator[int]](
		iterator.Range(1, 4, 1),
		func(i int) bool { return i%2 == 0 },
	)

	assert.Equal(t, containers.Vec(2), first)
	assert.Equal(t, containers.Vec(1, 3), second)
}
