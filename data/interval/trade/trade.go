package trade

import (
	"github.com/ta4g/ta4g/data/interval/trade/cost_model"
	"github.com/ta4g/ta4g/data/interval/trade/orders"
)

type Trade interface {
	GetEntry() orders.Order
	GetAdjustments() orders.Order
	GetExit() orders.Order
	GetOrderCostModel() cost_model.CostModel
	GetHoldingCostModel() cost_model.CostModel
}
