package cost_model

import (
	"github.com/stretchr/testify/require"
	"github.com/ta4g/ta4g/data/interval/trade/constants"
	"github.com/ta4g/ta4g/data/interval/trade/postion"
	"testing"
	"time"
)

func TestNoCostModel(t *testing.T) {
	costModel := NewNoCostModel()
	require.NotNil(t, costModel)

	buyStockOrder := postion.NewOrder(
		time.Now(),
		postion.NewStockOrderItem(constants.Buy, "ABC", 100, 10.01),
	)
	buyCoveredCallOrder := postion.NewOrder(
		time.Now(),
		postion.NewStockOrderItem(constants.Buy, "ABC", 100, 10.01),
		postion.NewOptionOrderItem(constants.Sell, "ABC CALL @ 10.0", 1, 1.01),
	)

	sellStockOrder := postion.NewOrder(
		time.Now(),
		postion.NewStockOrderItem(constants.Sell, "ABC", 100, 10.01),
	)
	sellCoveredCallOrder := postion.NewOrder(
		time.Now(),
		postion.NewStockOrderItem(constants.Sell, "ABC", 100, 10.01),
		postion.NewOptionOrderItem(constants.Buy, "ABC CALL @ 10.0", 1, 1.01),
	)

	t.Run("BuyStockOrder", func(t *testing.T) {
		orderCost, marginRequirement, err := costModel.BalanceChangeOnOpen(buyStockOrder)
		require.NoError(t, err)
		require.Equal(t, orderCost, -10.01*100)
		require.Equal(t, marginRequirement, 0.0)

		orderCost, marginRequirement, err = costModel.BalanceChangeOnClose(buyStockOrder)
		require.NoError(t, err)
		require.Equal(t, orderCost, 10.01*100)
		require.Equal(t, marginRequirement, 0.0)
	})

	t.Run("BuyCoveredCallOrder", func(t *testing.T) {
		orderCost, marginRequirement, err := costModel.BalanceChangeOnOpen(buyCoveredCallOrder)
		require.NoError(t, err)
		require.Equal(t, orderCost, -10.01*100+1.01*100)
		require.Equal(t, marginRequirement, 0.0)

		orderCost, marginRequirement, err = costModel.BalanceChangeOnClose(buyCoveredCallOrder)
		require.NoError(t, err)
		require.Equal(t, orderCost, 10.01*100-1.01*100)
		require.Equal(t, marginRequirement, 0.0)
	})

	t.Run("SellStockOrder", func(t *testing.T) {
		orderCost, marginRequirement, err := costModel.BalanceChangeOnOpen(sellStockOrder)
		require.NoError(t, err)
		require.Equal(t, orderCost, 10.01*100)
		require.Equal(t, marginRequirement, 0.0)

		orderCost, marginRequirement, err = costModel.BalanceChangeOnClose(sellStockOrder)
		require.NoError(t, err)
		require.Equal(t, orderCost, -10.01*100)
		require.Equal(t, marginRequirement, 0.0)
	})

	t.Run("SellCoveredCallOrder", func(t *testing.T) {
		orderCost, marginRequirement, err := costModel.BalanceChangeOnOpen(sellCoveredCallOrder)
		require.NoError(t, err)
		require.Equal(t, orderCost, 10.01*100-1.01*100)
		require.Equal(t, marginRequirement, 0.0)

		orderCost, marginRequirement, err = costModel.BalanceChangeOnClose(sellCoveredCallOrder)
		require.NoError(t, err)
		require.Equal(t, orderCost, -10.01*100+1.01*100)
		require.Equal(t, marginRequirement, 0.0)
	})
}
