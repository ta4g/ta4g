package trade_direction

// TradeDirection indicates if we are purchasing or selling some entity
type TradeDirection int

const (
	min TradeDirection = iota
	Buy
	Neutral
	Sell
	max
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
