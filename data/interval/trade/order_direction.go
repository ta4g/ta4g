package trade

// OrderDirection indicates if we are purchasing or selling some entity
type OrderDirection int

const (
	_ OrderDirection = iota
	BuyOrderDirection
	SellOrderDirection
)

const buyOrderDirectionStr = "Buy"
const sellOrderDirectionStr = "Sell"

var orderDirectionStrs = map[OrderDirection]string{
	BuyOrderDirection:  buyOrderDirectionStr,
	SellOrderDirection: sellOrderDirectionStr,
}

func (o OrderDirection) String() string {
	str, _ := orderDirectionStrs[o]
	return str
}
