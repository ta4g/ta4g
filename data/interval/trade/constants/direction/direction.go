package direction

// Direction indicates if we are purchasing or selling some entity
type Direction int

const (
	_ Direction = iota
	Buy
	Sell
)

const buyOrderDirectionStr = "Buy"
const sellOrderDirectionStr = "Sell"

var orderDirections = map[Direction]string{
	Buy:  buyOrderDirectionStr,
	Sell: sellOrderDirectionStr,
}

func (o Direction) String() string {
	str, _ := orderDirections[o]
	return str
}
