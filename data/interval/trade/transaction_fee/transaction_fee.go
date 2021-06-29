package transaction_fee

import (
	"github.com/ta4g/ta4g/data/interval/trade/constants/equity_type"
	"github.com/ta4g/ta4g/gen/proto/interval/transaction_fee"
)

type FeeAmount transaction_fee.FeeAmount
type TransactionFee transaction_fee.TransactionFee

type TransactionFeeConfig interface {
	// FeesByEquityType - map of EquityType to TransactionFee structure
	FeesByEquityType() map[equity_type.EquityType]*TransactionFee
}
