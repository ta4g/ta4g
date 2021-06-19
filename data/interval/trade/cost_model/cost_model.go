package cost_model

import "github.com/ta4g/ta4g/data/interval/trade/postion"

// CostModel is the pricing engine to compute the final cost of a given order.
//
// There are multiple types of pricing model listed out in the type assertions below:
// 1. NoCostModel - there are no fees and no margin interest
// 2. StandardCostModel - the normal fees and margin interest apply
// 3. RampUpCostModel - the price to place open goes up with each request to BalanceChangeOnOpen,
//    and goes down on each request to BalanceChangeOnClose however it is a percentage change so it results in ever increasing costs.
//
type CostModel interface {
	// BalanceChangeOnOpen returns the trading cost of a single order, this is the cost (or profit) of an opening the position
	BalanceChangeOnOpen(*postion.Order) (float64, float64, error)
	// BalanceChangeOnClose returns the trading cost of a single order, this is the cost (or profit) of a closing the position
	BalanceChangeOnClose(*postion.Order) (float64, float64, error)
}
