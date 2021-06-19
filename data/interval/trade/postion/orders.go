package postion

import "github.com/ta4g/ta4g/data/interval/trade/constants"

type Orders []*Order

func (o Orders) IsClosed() bool {
	items := map[string]float64{}
	for _, order := range o {
		for _, item := range order.OrderItems {
			if item.Direction == constants.Buy {
				items[item.Symbol] += item.Amount
			} else {
				items[item.Symbol] -= item.Amount
			}
			if items[item.Symbol] == 0.0 {
				delete(items, item.Symbol)
			}
		}
	}
	return len(items) == 0
}

func (o Orders) GetFundsAdded() Orders {
	return o.filterOrdersByType(constants.AddFundsOrderType)
}

func (o Orders) GetFundsRemoved() Orders {
	return o.filterOrdersByType(constants.RemoveFundsOrderType)
}

func (o Orders) GetEntries() Orders {
	return o.filterOrdersByType(constants.EnterOrderType)
}

func (o Orders) GetExits() Orders {
	return o.filterOrdersByType(constants.ExitOrderType)
}

func (o Orders) GetAdjustments() Orders {
	return o.filterOrdersByType(constants.AdjustmentOrderType)
}

func (o Orders) filterOrdersByType(orderType constants.OrderType) Orders {
	output := make(Orders, 0, len(o))
	for _, value := range o {
		if value.OrderType == orderType {
			output = append(output, value)
		}
	}
	return output
}
