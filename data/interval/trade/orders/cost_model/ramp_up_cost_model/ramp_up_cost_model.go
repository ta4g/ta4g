package ramp_up_cost_model

import (
	"github.com/ta4g/ta4g/data/interval/trade/orders"
	"github.com/ta4g/ta4g/data/interval/trade/orders/cost_model/standard_cost_model"
)

// RampUpCostModel is a cost model that is useful for back testing volatile markets
// The price to place open goes up with each request to BalanceChangeOnOpen, 
// and goes down on each request to BalanceChangeOnClose.
//
// However it is a percentage change so it results in ever increasing costs.
//
type RampUpCostModel struct {
	IncreasePct float64 `csv:"increase_pct" avro:"increase_pct" json:"increase_pct"`
	*standard_cost_model.StandardCostModel
}
// Compile type type enforcement
var _ orders.CostModel = &RampUpCostModel{}

func NewRampUpCostModel(increasePct float64, standardCostModel *standard_cost_model.StandardCostModel) orders.CostModel {
	return &RampUpCostModel{
		IncreasePct:       increasePct,
		StandardCostModel: standardCostModel,
	}
}

func (n *RampUpCostModel) BalanceChangeOnOpen(order *orders.Order) (float64, float64, error) {
	n.increase()
	return n.StandardCostModel.BalanceChangeOnOpen(order)
}

func (n *RampUpCostModel) BalanceChangeOnClose(order *orders.Order) (float64, float64, error) {
	n.decrease()
	return n.StandardCostModel.BalanceChangeOnClose(order)
}

func (n *RampUpCostModel) increase() {
	for _, fee := range []*standard_cost_model.Fees{n.USD, n.Stock, n.Option, n.Crypto} {
		fee.Exchange += fee.Exchange * n.IncreasePct
		fee.Order += fee.Order * n.IncreasePct
		fee.Amount += fee.Amount * n.IncreasePct
	}
}

func (n *RampUpCostModel) decrease() {
	for _, fee := range []*standard_cost_model.Fees{n.USD, n.Stock, n.Option, n.Crypto} {
		fee.Exchange -= fee.Exchange * n.IncreasePct
		fee.Order -= fee.Order * n.IncreasePct
		fee.Amount -= fee.Amount * n.IncreasePct
	}
}