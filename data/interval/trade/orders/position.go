package orders

import (
	"github.com/ta4g/ta4g/data/interval/trade/constants"
	"github.com/ta4g/ta4g/data/interval/trade/cost_model"
)

type Position struct {
	Orders
	cost_model.CostModel
}

func (p *Position) GetEntry() *Order {
	orders := p.filterOrdersByType(constants.EnterOrderType)
	if len(orders) == 0 {
		return nil
	}
	return orders[0]
}

func (p *Position) GetExit() *Order {
	orders := p.filterOrdersByType(constants.ExitOrderType)
	if len(orders) == 0 {
		return nil
	}
	return orders[0]
}

func (p *Position) GetAdjustments() Orders {
	return p.filterOrdersByType(constants.AdjustmentOrderType)
}

func (p *Position) filterOrdersByType(orderType constants.OrderType) Orders {
	output := make(Orders, 0, len(p.Orders))
	for _, value := range p.Orders {
		if value.OrderType == orderType {
			output = append(output, value)
		}
	}
	return output
}
