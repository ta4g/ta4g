package constants

type OrderType int

const (
	_ OrderType = iota

	// AddFundsOrderType
	// Add USD funds to an order, this is a "seed" step to allocate capital
	// or add capital when necessary (eg margin call).
	AddFundsOrderType

	// RemoveFundsOrderType
	// Remove USD funds from an order, this is when we are closing out a position
	// and removing funds form the active pool.
	RemoveFundsOrderType

	// EnterOrderType
	// Enter into a new order for a position, this is the starting point for any trade.
	EnterOrderType

	// ExitOrderType
	// Exit from an order's for a position, this is the ending point for any trade.
	ExitOrderType

	// AdjustmentOrderType
	// Adjust the order, this could mean:
	// - Moving an option contract up or down
	// - Accounting for a split where we go from 100 -> 200 shares with no change in price,
	// - Adjusting for a dividend where we are paid (or need to pay with a short).
	// - ...
	AdjustmentOrderType
)
const (
	addFundsOrderTypeStr    = "add_funds"
	removeFundsOrderTypeStr = "remove_funds"
	enterOrderTypeStr       = "enter"
	exitOrderTypeStr        = "exit"
	adjustmentOrderTypeStr  = "adjustment"
)

var orderTypes = map[OrderType]string{
	AddFundsOrderType:    addFundsOrderTypeStr,
	RemoveFundsOrderType: removeFundsOrderTypeStr,
	EnterOrderType:       enterOrderTypeStr,
	ExitOrderType:        exitOrderTypeStr,
	AdjustmentOrderType:  adjustmentOrderTypeStr,
}

func (o OrderType) String() string {
	return orderTypes[o]
}
