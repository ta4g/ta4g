package trade

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNoCostModel(t *testing.T) {
	costModel := NewNoCostModel()
	require.NotNil(t, costModel)

	buyStockOrder := NewStandardOrder(
		time.Now(),
		NewOrderItem(BuyOrderDirection, "ABC", false, 100, 10.01),
	)
	buyCoveredCallOrder := NewStandardOrder(
		time.Now(),
		NewOrderItem(BuyOrderDirection, "ABC", false, 100, 10.01),
		NewOrderItem(SellOrderDirection, "ABC CALL @ 10.0", true, 1, 1.01*100),
	)

	sellStockOrder := NewStandardOrder(
		time.Now(),
		NewOrderItem(SellOrderDirection, "ABC", false, 100, 10.01),
	)
	sellCoveredCallOrder := NewStandardOrder(
		time.Now(),
		NewOrderItem(SellOrderDirection, "ABC", false, 100, 10.01),
		NewOrderItem(BuyOrderDirection, "ABC CALL @ 10.0", true, 1, 1.01*100),
	)

	t.Run("BuyStockOrder", func(t *testing.T) {
		output, err := costModel.BalanceChangeOnOpen(buyStockOrder)
		require.NoError(t, err)
		require.Equal(t, output, -10.01*100)

		output, err = costModel.BalanceChangeOnClose(buyStockOrder)
		require.NoError(t, err)
		require.Equal(t, output, 10.01*100)
	})

	t.Run("BuyCoveredCallOrder", func(t *testing.T) {
		output, err := costModel.BalanceChangeOnOpen(buyCoveredCallOrder)
		require.NoError(t, err)
		require.Equal(t, output, -10.01*100+1.01*100)

		output, err = costModel.BalanceChangeOnClose(buyCoveredCallOrder)
		require.NoError(t, err)
		require.Equal(t, output, 10.01*100-1.01*100)
	})

	t.Run("SellStockOrder", func(t *testing.T) {
		output, err := costModel.BalanceChangeOnOpen(sellStockOrder)
		require.NoError(t, err)
		require.Equal(t, output, 10.01*100)

		output, err = costModel.BalanceChangeOnClose(sellStockOrder)
		require.NoError(t, err)
		require.Equal(t, output, -10.01*100)
	})

	t.Run("SellCoveredCallOrder", func(t *testing.T) {
		output, err := costModel.BalanceChangeOnOpen(sellCoveredCallOrder)
		require.NoError(t, err)
		require.Equal(t, output, 10.01*100-1.01*100)

		output, err = costModel.BalanceChangeOnClose(sellCoveredCallOrder)
		require.NoError(t, err)
		require.Equal(t, output, -10.01*100+1.01*100)
	})
}

func TestStandardCostModel(t *testing.T) {
	costModel := DefaultStandardCostModel()
	require.NotNil(t, costModel)

	buyStockOrder := NewStandardOrder(
		time.Now(),
		NewOrderItem(BuyOrderDirection, "ABC", false, 100, 10.01),
	)
	buyCoveredCallOrder := NewStandardOrder(
		time.Now(),
		NewOrderItem(BuyOrderDirection, "ABC", false, 100, 10.01),
		NewOrderItem(SellOrderDirection, "ABC CALL @ 10.0", true, 1, 1.01*100),
	)

	sellStockOrder := NewStandardOrder(
		time.Now(),
		NewOrderItem(SellOrderDirection, "ABC", false, 100, 10.01),
	)
	sellCoveredCallOrder := NewStandardOrder(
		time.Now(),
		NewOrderItem(SellOrderDirection, "ABC", false, 100, 10.01),
		NewOrderItem(BuyOrderDirection, "ABC CALL @ 10.0", true, 1, 1.01*100),
	)

	t.Run("BuyStockOrder", func(t *testing.T) {
		output, err := costModel.BalanceChangeOnOpen(buyStockOrder)
		require.NoError(t, err)
		require.Equal(t, output, -1002.4)

		output, err = costModel.BalanceChangeOnClose(buyStockOrder)
		require.NoError(t, err)
		require.Equal(t, output, 1002.4)
	})

	t.Run("BuyCoveredCallOrder", func(t *testing.T) {
		output, err := costModel.BalanceChangeOnOpen(buyCoveredCallOrder)
		require.NoError(t, err)
		require.Equal(t, output, -900.65)

		output, err = costModel.BalanceChangeOnClose(buyCoveredCallOrder)
		require.NoError(t, err)
		require.Equal(t, output, 900.65)
	})

	t.Run("SellStockOrder", func(t *testing.T) {
		output, err := costModel.BalanceChangeOnOpen(sellStockOrder)
		require.NoError(t, err)
		require.Equal(t, output, 1002.4)

		output, err = costModel.BalanceChangeOnClose(sellStockOrder)
		require.NoError(t, err)
		require.Equal(t, output, -1002.4)
	})

	t.Run("SellCoveredCallOrder", func(t *testing.T) {
		output, err := costModel.BalanceChangeOnOpen(sellCoveredCallOrder)
		require.NoError(t, err)
		require.Equal(t, output, 900.65)

		output, err = costModel.BalanceChangeOnClose(sellCoveredCallOrder)
		require.NoError(t, err)
		require.Equal(t, output, -900.65)
	})
}

func TestEvilGeniusCostModel(t *testing.T) {
	t.Run("New", func(t *testing.T) {})
	t.Run("BalanceChangeOnOpen", func(t *testing.T) {})
	t.Run("BalanceChangeOnClose", func(t *testing.T) {})
	t.Run("increase", func(t *testing.T) {})
}
