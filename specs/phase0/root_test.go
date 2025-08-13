package phase0_test

import (
	"testing"

	"github.com/ethpandaops/go-lean-types/specs/phase0"
	"github.com/stretchr/testify/require"
)

func TestZeroRoot(t *testing.T) {
	zeroRoot := &phase0.Root{}
	require.True(t, zeroRoot.IsZero())

	nonZeroRoot := &phase0.Root{0x01}
	require.False(t, nonZeroRoot.IsZero())
}
