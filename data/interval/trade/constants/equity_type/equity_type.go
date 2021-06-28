package equity_type

import "github.com/ta4g/ta4g/gen/proto/interval/equity_type"

type EquityType equity_type.EquityType

const (
	min    = EquityType(equity_type.EquityType_UNKNOWN)
	Cash   = EquityType(equity_type.EquityType_CASH)
	Stock  = EquityType(equity_type.EquityType_STOCK)
	Option = EquityType(equity_type.EquityType_OPTION)
	Crypto = EquityType(equity_type.EquityType_CRYPTO)
	max    = EquityType(equity_type.EquityType_CRYPTO + 1)
)

func (e EquityType) ToProto() equity_type.EquityType {
	return equity_type.EquityType(e)
}

func (e EquityType) String() string {
	return e.ToProto().String()
}
