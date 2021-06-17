package trade

import (
	"github.com/ta4g/ta4g/data/interval/trade/orders"
)

type Trade interface {
	GetEntry() orders.Order
	GetAdjustments() orders.Order
	GetExit() orders.Order
	GetOrderCostModel() orders.CostModel
	GetHoldingCostModel() orders.CostModel
}
