package order_type

import "github.com/ta4g/ta4g/gen/interval/order_type"

type OrderType order_type.OrderType

const (
	min                = OrderType(order_type.OrderType_UNKNOWN)
	PortfolioOpen      = OrderType(order_type.OrderType_PORTFOLIO_OPEN)
	PortfolioClose     = OrderType(order_type.OrderType_PORTFOLIO_CLOSE)
	PositionOpen       = OrderType(order_type.OrderType_POSITION_OPEN)
	PositionClose      = OrderType(order_type.OrderType_POSITION_CLOSE)
	PositionAdjustment = OrderType(order_type.OrderType_POSITION_ADJUSTMENT)
	max                = OrderType(order_type.OrderType_POSITION_ADJUSTMENT + 1)
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
