package orders

import (
	"github.com/ta4g/ta4g/data/interval/trade/constants"
)

type Position struct {
	Orders
	CostModel
}

func (p *Position) GetFundsAdded() Orders {
	return p.filterOrdersByType(constants.AddFundsOrderType)
}

func (p *Position) GetFundsRemoved() Orders {
	return p.filterOrdersByType(constants.RemoveFundsOrderType)
}

func (p *Position) GetEntries() Orders {
	return p.filterOrdersByType(constants.EnterOrderType)
}

func (p *Position) GetExits() Orders {
	return p.filterOrdersByType(constants.ExitOrderType)
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
