package equity_type

import "github.com/ta4g/ta4g/gen/interval/trade"

type EquityType trade.EquityType

const (
	min EquityType = EquityType(trade.EquityType_UNKNOWN_EQUITY_TYPE)
	Cash = EquityType(trade.EquityType_CASH_EQUITY_TYPE)
	Stock = EquityType(trade.EquityType_STOCK_EQUITY_TYPE)
	Option = EquityType(trade.EquityType_OPTION_EQUITY_TYPE)
	Crypto = EquityType(trade.EquityType_CRYPTO_EQUITY_TYPE)
	max = EquityType(trade.EquityType_CRYPTO_EQUITY_TYPE+1)
)

type TransactionBalance struct {
	AvailableCash     float64 // The amount of cash currently available to use.
	AvailableMargin   float64 // The amount of margin currently available to use.
	MaintenanceMargin float64 // How much margin we need to maintain to keep the position open
	CostBasisValue    float64 // The cost basis for all open positions.
	CurrentValue      float64 // The value for all open positions.
}

var itemTypes = map[EquityType]string{
	Cash:   "cash",
	Stock:  "stock",
	Option: "option",
	Crypto: "crypto",
}

func (i EquityType) String() string {
	return itemTypes[i]
}
