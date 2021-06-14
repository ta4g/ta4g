package constants

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOrderDirection(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		require.Equal(t, Sell.String(), sellOrderDirectionStr)
		require.Equal(t, Buy.String(), buyOrderDirectionStr)
	})
}
