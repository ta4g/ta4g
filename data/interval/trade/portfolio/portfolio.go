package portfolio

import (
	"github.com/ta4g/ta4g/data/interval/trade/cost_model"
	"github.com/ta4g/ta4g/data/interval/trade/orders"
)

type Portfolio struct {
	Positions map[string]orders.Orders
	cost_model.CostModel
}
