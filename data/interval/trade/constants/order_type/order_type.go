package order_type

import "github.com/ta4g/ta4g/gen/interval/trade"

type OrderType trade.OrderType

const (
	min                OrderType = OrderType(trade.OrderType_UNKNOWN_ORDER_TYPE)
	PortfolioOpen                = OrderType(trade.OrderType_PORTFOLIO_OPEN_ORDER_TYPE)
	PortfolioClose               = OrderType(trade.OrderType_PORTFOLIO_CLOSE_ORDER_TYPE)
	PositionOpen                 = OrderType(trade.OrderType_POSITION_OPEN_ORDER_TYPE)
	PositionClose                = OrderType(trade.OrderType_POSITION_CLOSE_ORDER_TYPE)
	PositionAdjustment           = OrderType(trade.OrderType_POSITION_ADJUSTMENT_ORDER_TYPE)
	max                          = OrderType(trade.OrderType_POSITION_ADJUSTMENT_ORDER_TYPE + 1)
)

var orderTypes = map[OrderType]string{
	PortfolioOpen:      "portfolio open",
	PortfolioClose:     "portfolio close",
	PositionOpen:       "position open",
	PositionClose:      "position close",
	PositionAdjustment: "position adjustment",
}

func (o OrderType) String() string {
	return orderTypes[o]
}
