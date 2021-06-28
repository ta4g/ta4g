package transaction_fee

import (
	"github.com/ta4g/ta4g/data/interval/trade/constants/equity_type"
	"github.com/ta4g/ta4g/gen/interval/trade"
)

type TransactionFee trade.TransactionFee

type TransactionFeeConfig interface {
	// FeesByEquityType - map of EquityType to TransactionFee structure
	FeesByEquityType() map[equity_type.EquityType]*TransactionFee
}
