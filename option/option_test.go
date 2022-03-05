package option_test

import (
	"testing"

	"github.com/stntngo/rust-lib/option"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOption(t *testing.T) {
	none := option.None[int]()
	assert.False(t, none.Valid())
	assert.Panics(t, func() { none.Unwrap() })

	some := option.Some(32891)
	assert.True(t, some.Valid())
	require.NotPanics(t, func() { some.Unwrap() })
	assert.Equal(t, 32891, some.Unwrap())
}
