package constants

type ItemType int

const (
	_      ItemType = iota
	USD             // USD is a cash item used for tracking the state of an account
	Stock           // Stock is the standard share of stock you can purchase on an exchange
	Option          // Option is a type of derivative that allows you to "control" more than 1x share of stock for a limited period of time
	Crypto          // Crypto is any cryptocurrency that is available
)

const (
	usdItemTypeStr    = "USD"
	stockItemTypeStr  = "Stock"
	optionItemTypeStr = "Option"
	cryptoItemTypeStr = "Crypto"
)

var itemTypes = map[ItemType]string{
	USD:    usdItemTypeStr,
	Stock:  stockItemTypeStr,
	Option: optionItemTypeStr,
	Crypto: cryptoItemTypeStr,
}

func (i ItemType) String() string {
	return itemTypes[i]
}
