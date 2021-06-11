package trade

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOrderDirection(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		require.Equal(t, SellOrderDirection.String(), sellOrderDirectionStr)
		require.Equal(t, BuyOrderDirection.String(), buyOrderDirectionStr)
	})
}
