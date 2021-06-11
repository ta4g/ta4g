package trade

type CostModel interface {
	// BalanceChangeOnOpen returns the trading cost of a single order, this is the cost (or profit) of an opening the position
	BalanceChangeOnOpen(Order) (float64, error)
	// BalanceChangeOnClose returns the trading cost of a single order, this is the cost (or profit) of a closing the position
	BalanceChangeOnClose(Order) (float64, error)
}

// Compile type type enforcement
var _ CostModel = &NoCostModel{}
var _ CostModel = &StandardCostModel{}
var _ CostModel = &EvilGeniusCostModel{}

//
// NoCostModel
//

// NoCostModel is used when there are no trading fees, no maintenance fees, and no brokerage fees
// This is useful for a theoretical back test on the 100% most optimal performance,
// and should only be used as a benchmark for testing
type NoCostModel struct{}

func NewNoCostModel() CostModel {
	return &NoCostModel{}
}

func (n NoCostModel) BalanceChangeOnOpen(order Order) (float64, error) {
	total := 0.0
	for _, item := range order.GetItems() {
		total += item.GetNetCost(0.0, 0.0, 0.0)
	}
	return total, nil
}

func (n NoCostModel) BalanceChangeOnClose(order Order) (float64, error) {
	total, err := n.BalanceChangeOnOpen(order)
	if nil != err {
		return 0.0, err
	}
	return -1 * total, nil
}

//
// StandardCostModel
//

// StandardCostModel is used when there are the usual options/maintenance/broker fees.
// This is useful for more real back test with more real performance
type StandardCostModel struct {
	ExchangeStockFee float64
	FixStockFee      float64
	LinearStockFee   float64

	ExchangeOptionFee float64
	FixOptionFee      float64
	LinearOptionFee   float64

	MarginRates []RateRange
}

func NewStandardCostModel(
	exchangeStockFee,
	fixStockFee,
	linearStockFee,
	exchangeOptionFee,
	fixOptionFee,
	linearOptionFee float64,
	marginRates []RateRange,
) CostModel {
	return &StandardCostModel{
		ExchangeStockFee:  exchangeStockFee,
		FixStockFee:       fixStockFee,
		LinearStockFee:    linearStockFee,
		ExchangeOptionFee: exchangeOptionFee,
		FixOptionFee:      fixOptionFee,
		LinearOptionFee:   linearOptionFee,
		MarginRates:       marginRates,
	}
}

func DefaultStandardCostModel() CostModel {
	return &StandardCostModel{
		ExchangeStockFee:  0.75, // Minimal fee from the exchange
		FixStockFee:       0.0,
		LinearStockFee:    0.0,
		ExchangeOptionFee: 0.75, // Minimal fee from the exchange
		FixOptionFee:      0.65, // Minimal fee per contract
		LinearOptionFee:   0.0,
		MarginRates:       StandardRateRanges,
	}
}

func (s StandardCostModel) BalanceChangeOnOpen(order Order) (float64, error) {
	stockTotal := float64(0)
	marginTotal := float64(0)

	for _, item := range order.GetItems() {
		if item.GetIsOption() {
			marginTotal += item.GetNetCost(s.ExchangeStockFee, s.FixStockFee, s.LinearStockFee)
		} else {
			stockTotal += item.GetNetCost(s.ExchangeOptionFee, s.FixOptionFee, s.LinearOptionFee)
		}
	}

	output := stockTotal + marginTotal
	return output, nil
}

func (s StandardCostModel) BalanceChangeOnClose(order Order) (float64, error) {
	// TODO: We don't have maintenance cost yet, but when we do this function will change
	total, err := s.BalanceChangeOnOpen(order)
	if nil != err {
		return 0.0, err
	}
	return -1 * total, nil
}

//
// EvilGeniusCostModel
//

// EvilGeniusCostModel is a cost model that is useful for back testing volatile markets
// This cost model will increase a fixed percentage each time you call any of the functions
type EvilGeniusCostModel struct {
	IncreaseAmount float64
	*StandardCostModel
}

func NewEvilGeniusCostModel() CostModel {
	return &EvilGeniusCostModel{}
}

func (n *EvilGeniusCostModel) BalanceChangeOnOpen(order Order) (float64, error) {
	n.increase()
	return n.StandardCostModel.BalanceChangeOnOpen(order)
}

func (n *EvilGeniusCostModel) BalanceChangeOnClose(order Order) (float64, error) {
	n.increase()
	return n.StandardCostModel.BalanceChangeOnClose(order)
}

func (n *EvilGeniusCostModel) increase() {
	n.StandardCostModel.ExchangeStockFee *= n.IncreaseAmount
	n.StandardCostModel.FixStockFee *= n.IncreaseAmount
	n.StandardCostModel.LinearStockFee *= n.IncreaseAmount

	n.StandardCostModel.ExchangeOptionFee *= n.IncreaseAmount
	n.StandardCostModel.FixOptionFee *= n.IncreaseAmount
	n.StandardCostModel.LinearOptionFee *= n.IncreaseAmount
}
