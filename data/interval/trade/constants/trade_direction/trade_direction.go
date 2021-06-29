package trade_direction

import "github.com/ta4g/ta4g/gen/proto/interval/trade_direction"

// TradeDirection indicates if we are purchasing or selling some entity
type TradeDirection = trade_direction.TradeDirection

const (
	min     = trade_direction.TradeDirection_UNKNOWN
	Buy     = trade_direction.TradeDirection_BUY
	Neutral = trade_direction.TradeDirection_NEUTRAL
	Sell    = trade_direction.TradeDirection_SELL
	max     = trade_direction.TradeDirection_SELL + 1
)

var opposites = map[TradeDirection]TradeDirection{
	Buy:  Sell,
	Sell: Buy,
}

func Opposite(o TradeDirection) TradeDirection {
	output, ok := opposites[o]
	if !ok {
		return o
	}
	return output
}

func FromProto(o trade_direction.TradeDirection) TradeDirection {
	return o
}

func ToProto(o TradeDirection) trade_direction.TradeDirection {
	return o
}
