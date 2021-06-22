package cost_model

import (
	"github.com/ta4g/ta4g/data/interval/trade/constants/equity_type"
	"github.com/ta4g/ta4g/data/interval/trade/postion"
	"github.com/ta4g/ta4g/data/interval/trade/transaction_fee"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// StandardCostModel is used when there are the usual options/maintenance/broker transaction_fee.
// This is useful for more real back test with more real performance
type StandardCostModel struct {
	Cash   *transaction_fee.TransactionFee `csv:"cash" avro:"cash" json:"cash"`
	Stock  *transaction_fee.TransactionFee `csv:"stock" avro:"stock" json:"stock"`
	Option *transaction_fee.TransactionFee `csv:"option" avro:"option" json:"option"`
	Crypto *transaction_fee.TransactionFee `csv:"crypto" avro:"crypto" json:"crypto"`
}

// Compile type type enforcement
var _ CostModel = &StandardCostModel{}

// NewStandardCostModel creates a new CostModel instance using the given transaction_fee
func NewStandardCostModel(cash, stock, option, crypto *transaction_fee.TransactionFee) CostModel {
	return &StandardCostModel{
		Cash:   cash,
		Stock:  stock,
		Option: option,
		Crypto: crypto,
	}
}

// DefaultStandardCostModel is the pre-canned cost model using the transaction_fee currently posted
// to TD Ameritrade and Coinbase.
func DefaultStandardCostModel() CostModel {
	return NewStandardCostModel(
		// Cash, free to hold and exchange
		&transaction_fee.TransactionFee{},
		// Stocks are free to buy and sell, but there is an exchange fee
		&transaction_fee.TransactionFee{Exchange: 0.75},
		// Options hav an exchange and per-contract fee
		&transaction_fee.TransactionFee{Exchange: 0.75, Amount: 0.65},
		// Crypto is a flat-fee (estimation for simplicity)
		&transaction_fee.TransactionFee{Order: 0.99},
	)
}

func (s StandardCostModel) Fees() map[equity_type.EquityType]*transaction_fee.TransactionFee {
	return map[equity_type.EquityType]*transaction_fee.TransactionFee{
		equity_type.Crypto: s.Crypto,
		equity_type.Cash:   s.Cash,
		equity_type.Option: s.Option,
		equity_type.Stock:  s.Stock,
	}
}

func (s StandardCostModel) BalanceChangeOnOpen(order *postion.Order) (float64, float64, error) {
	orderCost := float64(0)
	marginRequirement := float64(0)

	for _, item := range order.OrderItems {
		var fee *transaction_fee.TransactionFee
		switch item.EquityType {
		case equity_type.Cash:
			fee = s.Cash
		case equity_type.Stock:
			fee = s.Stock
		case equity_type.Option:
			fee = s.Option
		case equity_type.Crypto:
			fee = s.Crypto
		default:
			return 0, 0, status.Error(codes.OutOfRange, "unknown fee type")
		}

		orderCost += item.CalculatePrice(fee.Exchange, fee.Order, fee.Amount)
		marginRequirement += item.MarginRequirement()
	}
	return orderCost, marginRequirement, nil
}

func (s StandardCostModel) BalanceChangeOnClose(order *postion.Order) (float64, float64, error) {
	// TODO: We don't have maintenance cost yet, but when we do this function will change
	orderCost, marginRequirement, err := s.BalanceChangeOnOpen(order)
	if nil != err {
		return 0, 0, err
	}
	return -1 * orderCost, -1 * marginRequirement, nil
}
