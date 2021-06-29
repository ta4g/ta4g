package postion

import (
	"github.com/ta4g/ta4g/data/interval/trade/constants/order_type"
	"github.com/ta4g/ta4g/data/interval/trade/constants/trade_direction"
)

type Trades []*Trade

func (o Trades) IsClosed() bool {
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

func (o Trades) GetFundsAdded() Trades {
	return o.filterOrdersByType(order_type.PortfolioOpen)
}

func (o Trades) GetFundsRemoved() Trades {
	return o.filterOrdersByType(order_type.PortfolioClose)
}

func (o Trades) GetEntries() Trades {
	return o.filterOrdersByType(order_type.PositionOpen)
}

func (o Trades) GetExits() Trades {
	return o.filterOrdersByType(order_type.PositionClose)
}

func (o Trades) GetAdjustments() Trades {
	return o.filterOrdersByType(order_type.PositionAdjustment)
}

func (o Trades) filterOrdersByType(orderType order_type.OrderType) Trades {
	output := make(Trades, 0, len(o))
	for _, value := range o {
		if value.OrderType == orderType {
			output = append(output, value)
		}
	}
	return output
}
