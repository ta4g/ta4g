package cost_model

import (
	"github.com/ta4g/ta4g/data/interval/trade/orders"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Fees struct {
	// Exchange is the fee per Order that is charged
	Exchange float64 `csv:"exchange" avro:"exchange" json:"exchange"`
	// Order is the fee per Order that is charged
	Order float64 `csv:"order" avro:"order" json:"order"`
	// Amount is the fee per Amount that is charged
	Amount float64 `csv:"amount" avro:"amount" json:"amount"`
}

// StandardCostModel is used when there are the usual options/maintenance/broker fees.
// This is useful for more real back test with more real performance
type StandardCostModel struct {
	USD    *Fees `csv:"usd" avro:"usd" json:"usd"`
	Stock  *Fees `csv:"stock" avro:"stock" json:"stock"`
	Option *Fees `csv:"option" avro:"option" json:"option"`
	Crypto *Fees `csv:"crypto" avro:"crypto" json:"crypto"`
}

// NewStandardCostModel creates a new CostModel instance using the given fees
func NewStandardCostModel(usd, stock, option, crypto *Fees) CostModel {
	return &StandardCostModel{
		USD:    usd,
		Stock:  stock,
		Option: option,
		Crypto: crypto,
	}
}

// DefaultStandardCostModel is the pre-canned cost model using the fees currently posted
// to TD Ameritrade and Coinbase.
func DefaultStandardCostModel() CostModel {
	return NewStandardCostModel(
		// USD, free to hold and exchange
		&Fees{},
		// Stocks are free to buy and sell, but there is an exchange fee
		&Fees{Exchange: 0.75},
		// Options hav an exchange and per-contract fee
		&Fees{Exchange: 0.75, Amount: 0.65},
		// Crypto is a flat-fee (estimation for simplicity)
		&Fees{Order: 0.99},
	)
}

func (s StandardCostModel) BalanceChangeOnOpen(order *orders.Order) (float64, float64, error) {
	orderCost := float64(0)
	marginRequirement := float64(0)

	for _, item := range order.OrderItems {
		var fee *Fees
		switch item.ItemType {
		case orders.USD:
			fee = s.USD
		case orders.Stock:
			fee = s.Stock
		case orders.Option:
			fee = s.Option
		case orders.Crypto:
			fee = s.Crypto
		default:
			return 0, 0, status.Error(codes.OutOfRange, "unknown fee type")
		}

		orderCost += item.CalculatePrice(fee.Exchange, fee.Order, fee.Amount)
		marginRequirement += item.MarginRequirement()
	}
	return orderCost, marginRequirement, nil
}

func (s StandardCostModel) BalanceChangeOnClose(order *orders.Order) (float64, float64, error) {
	// TODO: We don't have maintenance cost yet, but when we do this function will change
	orderCost, marginRequirement, err := s.BalanceChangeOnOpen(order)
	if nil != err {
		return 0, 0, err
	}
	return -1 * orderCost, -1 * marginRequirement, nil
}
