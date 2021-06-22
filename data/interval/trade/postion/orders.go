package postion

import (
	"github.com/ta4g/ta4g/data/interval/trade/constants/trade_direction"
	"github.com/ta4g/ta4g/data/interval/trade/constants/order_type"
)

type Orders []*Order

func (o Orders) IsClosed() bool {
	items := map[string]float64{}
	for _, order := range o {
		for _, item := range order.OrderItems {
			if item.TradeDirection == trade_direction.Buy {
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
	return o.filterOrdersByType(order_type.PortfolioOpen)
}

func (o Orders) GetFundsRemoved() Orders {
	return o.filterOrdersByType(order_type.PortfolioClose)
}

func (o Orders) GetEntries() Orders {
	return o.filterOrdersByType(order_type.PositionOpen)
}

func (o Orders) GetExits() Orders {
	return o.filterOrdersByType(order_type.PositionClose)
}

func (o Orders) GetAdjustments() Orders {
	return o.filterOrdersByType(order_type.PositionAdjustment)
}

func (o Orders) filterOrdersByType(orderType order_type.OrderType) Orders {
	output := make(Orders, 0, len(o))
	for _, value := range o {
		if value.OrderType == orderType {
			output = append(output, value)
		}
	}
	return output
}
