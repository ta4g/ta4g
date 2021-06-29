package order_type

import "github.com/ta4g/ta4g/gen/proto/interval/order_type"

type OrderType = order_type.OrderType

const (
	min                = order_type.OrderType_UNKNOWN
	PortfolioOpen      = order_type.OrderType_PORTFOLIO_OPEN
	PortfolioClose     = order_type.OrderType_PORTFOLIO_CLOSE
	PositionOpen       = order_type.OrderType_POSITION_OPEN
	PositionClose      = order_type.OrderType_POSITION_CLOSE
	PositionAdjustment = order_type.OrderType_POSITION_ADJUSTMENT
	max                = order_type.OrderType_POSITION_ADJUSTMENT + 1
)

var opposites = map[OrderType]OrderType{
	PortfolioOpen:  PortfolioClose,
	PortfolioClose: PortfolioOpen,
	PositionOpen:   PositionClose,
	PositionClose:  PositionOpen,
}

func Opposite(o OrderType) OrderType {
	output, ok := opposites[o]
	if !ok {
		return o
	}
	return output
}

func FromProto(e order_type.OrderType) OrderType {
	return e
}

func ToProto(e OrderType) order_type.OrderType {
	return e
}
