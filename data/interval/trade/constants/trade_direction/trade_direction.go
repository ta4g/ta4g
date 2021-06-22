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

func (o TradeDirection) Opposite() TradeDirection {
	switch o {
	case Sell:
		return Buy
	case Buy:
		return Sell
	case min:
		fallthrough
	case max:
		fallthrough
	default:
		return o
	}
}

func (o TradeDirection) String() string {
	str, _ := orderDirections[o]
	return str
}
