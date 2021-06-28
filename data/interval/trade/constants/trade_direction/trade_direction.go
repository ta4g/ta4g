package trade_direction

import "github.com/ta4g/ta4g/gen/proto/interval/trade_direction"

// TradeDirection indicates if we are purchasing or selling some entity
type TradeDirection trade_direction.TradeDirection

const (
	min     = TradeDirection(trade_direction.TradeDirection_UNKNOWN)
	Buy     = TradeDirection(trade_direction.TradeDirection_BUY)
	Neutral = TradeDirection(trade_direction.TradeDirection_NEUTRAL)
	Sell    = TradeDirection(trade_direction.TradeDirection_SELL)
	max     = TradeDirection(trade_direction.TradeDirection_SELL + 1)
)

var opposites = map[TradeDirection]TradeDirection{
	Buy:  Sell,
	Sell: Buy,
}

func (o TradeDirection) ToProto() trade_direction.TradeDirection {
	return trade_direction.TradeDirection(o)
}

func (o TradeDirection) Opposite() TradeDirection {
	output, ok := opposites[o]
	if !ok {
		return o
	}
	return output
}

func (o TradeDirection) String() string {
	return o.ToProto().String()
}
