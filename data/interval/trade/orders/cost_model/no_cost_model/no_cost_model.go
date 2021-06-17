package no_cost_model

import "github.com/ta4g/ta4g/data/interval/trade/orders"

// NoCostModel is used when there are no trading fees, no maintenance fees, and no brokerage fees
// This is useful for a theoretical back test on the 100% most optimal performance,
// and should only be used as a benchmark for testing
type NoCostModel struct{}

// Compile type type enforcement
var _ orders.CostModel = &NoCostModel{}

func NewNoCostModel() orders.CostModel {
	return &NoCostModel{}
}

func (n NoCostModel) BalanceChangeOnOpen(order *orders.Order) (float64, float64, error) {
	total := 0.0
	for _, item := range order.OrderItems {
		total += item.CalculatePrice(0.0, 0.0, 0.0)
	}
	return total, 0.0, nil
}

func (n NoCostModel) BalanceChangeOnClose(order *orders.Order) (float64, float64, error) {
	total, _, err := n.BalanceChangeOnOpen(order)
	if nil != err {
		return 0, 0, err
	}
	return -1 * total, 0.0, nil
}
