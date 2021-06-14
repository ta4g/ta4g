package orders

import "github.com/ta4g/ta4g/data/interval/trade/cost_model"

type Position struct {
	Entry       *Order
	Adjustments []*Order
	Exit        *Order
	cost_model.CostModel
}
