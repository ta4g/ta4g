package constants

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOrderType(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		require.Equal(t, EnterOrderType.String(), enterOrderTypeStr)
		require.Equal(t, ExitOrderType.String(), exitOrderTypeStr)
		require.Equal(t, AdjustmentOrderType.String(), adjustmentOrderTypeStr)
	})
}
