package cost_model

import (
	"github.com/ta4g/ta4g/data/interval/trade/constants/equity_type"
	"github.com/ta4g/ta4g/data/interval/trade/postion"
	"github.com/ta4g/ta4g/data/interval/trade/transaction_fee"
)

// NoCostModel is used when there are no trading transaction_fee, no maintenance transaction_fee, and no brokerage transaction_fee
// This is useful for a theoretical back test on the 100% most optimal performance,
// and should only be used as a benchmark for testing
type NoCostModel struct{}

// Compile type type enforcement
var _ CostModel = &NoCostModel{}

func NewNoCostModel() CostModel {
	return &NoCostModel{}
}

func (n NoCostModel) Fees() map[equity_type.EquityType]*transaction_fee.TransactionFee {
	noFees := &transaction_fee.TransactionFee{}

	return map[equity_type.EquityType]*transaction_fee.TransactionFee{
		equity_type.Crypto: noFees,
		equity_type.Cash:   noFees,
		equity_type.Option: noFees,
		equity_type.Stock:  noFees,
	}
}

func (n NoCostModel) BalanceChangeOnOpen(order *postion.Order) (float64, float64, error) {
	total := 0.0
	for _, item := range order.OrderItems {
		total += item.CalculatePrice(0.0, 0.0, 0.0)
	}
	return total, 0.0, nil
}

func (n NoCostModel) BalanceChangeOnClose(order *postion.Order) (float64, float64, error) {
	total, _, err := n.BalanceChangeOnOpen(order)
	if nil != err {
		return 0, 0, err
	}
	return -1 * total, 0.0, nil
}
