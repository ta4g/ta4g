package equity_type

import "github.com/ta4g/ta4g/gen/proto/interval/equity_type"

type EquityType = equity_type.EquityType

const (
	min    = equity_type.EquityType_UNKNOWN
	Cash   = equity_type.EquityType_CASH
	Stock  = equity_type.EquityType_STOCK
	Option = equity_type.EquityType_OPTION
	Crypto = equity_type.EquityType_CRYPTO
	max    = equity_type.EquityType_CRYPTO + 1
)

func FromProto(e equity_type.EquityType) EquityType {
	return e
}

func ToProto(e EquityType) equity_type.EquityType {
	return e
}
