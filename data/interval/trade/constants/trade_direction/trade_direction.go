package trade_direction

import "github.com/ta4g/ta4g/gen/interval/trade"

// TradeDirection indicates if we are purchasing or selling some entity
type TradeDirection trade.TradeDirection

const (
	min     TradeDirection = TradeDirection(trade.TradeDirection_UNKNOWN_TRADE_DIRECTION)
	Buy                    = TradeDirection(trade.TradeDirection_BUY_TRADE_DIRECTION)
	Neutral                = TradeDirection(trade.TradeDirection_NEUTRAL_TRADE_DIRECTION)
	Sell                   = TradeDirection(trade.TradeDirection_SELL_TRADE_DIRECTION)
	max                    = TradeDirection(trade.TradeDirection_SELL_TRADE_DIRECTION + 1)
)

var orderDirections = map[TradeDirection]string{
	Buy:     "buy",
	Neutral: "neutral",
	Sell:    "sell",
}

var opposites = map[TradeDirection]TradeDirection{
	Buy:  Sell,
	Sell: Buy,
}

func (o TradeDirection) Opposite() TradeDirection {
	output, ok := opposites[o]
	if !ok {
		return o
	}
	return output
}

func (o TradeDirection) String() string {
	str := orderDirections[o]
	return str
}
