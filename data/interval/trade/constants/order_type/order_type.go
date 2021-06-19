package order_type

// OrderType - what is the nature of an order?
// - Portfolio Open - start a new portfolio
// - Portfolio Close - close a portfolio
// -
type OrderType int

const (
	minOrderType OrderType = iota

	// PortfolioOpen - add funds to a Portfolio, this is a "seed" step to allocate capital
	// to a portfolio on creation.
	PortfolioOpen

	// PortfolioClose - remove all the funds from a Portfolio, this is the "closeout" step to remove all capital
	// from a portfolio on completion.
	PortfolioClose

	// PositionOpen - enter into a trade for a given stock/option/etc.
	//
	// This is the starting point for any trade, based on the nature of the trade we expect to:
	// 1. Buy/Sell a Stock
	// 2. Buy/Sell an Option(s)
	// 3. Buy a crypto asset
	//
	PositionOpen

	// PositionClose - close out all or part-of an order for a stock/option/etc.
	// This is the opposite of a PositionOpen where if we purchased stock we then expect to sell a stock. etc.
	PositionClose

	// AdjustmentOrderType
	// Adjust the order, this could mean:
	// - Moving an option contract up or down
	// - Accounting for a split where we go from 100 -> 200 shares with no change in price,
	// - Adjusting for a dividend where we are paid (or need to pay with a short).
	// - ...
	AdjustmentOrderType

	maxOrderType
)

var orderTypes = map[OrderType]string{
	PortfolioOpen:       "portfolio open",
	PortfolioClose:      "portfolio close",
	PositionOpen:        "enter",
	PositionClose:       "exit",
	AdjustmentOrderType: "adjustment",
}

func (o OrderType) String() string {
	return orderTypes[o]
}
