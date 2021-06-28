package portfolio

import (
	"github.com/ta4g/ta4g/data/interval/trade/cost_model"
	"github.com/ta4g/ta4g/data/interval/trade/postion"
)

type Portfolio struct {
	Positions map[string]postion.Trades
	cost_model.CostModel
}
