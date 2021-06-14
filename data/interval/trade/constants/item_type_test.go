package constants

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestItemType(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		require.Equal(t, USD.String(), usdItemTypeStr)
		require.Equal(t, Stock.String(), stockItemTypeStr)
		require.Equal(t, Option.String(), optionItemTypeStr)
		require.Equal(t, Crypto.String(), cryptoItemTypeStr)
	})
}
