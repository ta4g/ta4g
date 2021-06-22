package transaction_fee

import "github.com/ta4g/ta4g/data/interval/trade/constants/equity_type"

type TransactionFee struct {
	// Exchange is the fee per Order that is charged
	Exchange float64 `csv:"exchange" avro:"exchange" json:"exchange"`
	// Order is the fee per Order that is charged
	Order float64 `csv:"order" avro:"order" json:"order"`
	// Amount is the fee per Amount that is charged
	Amount float64 `csv:"amount" avro:"amount" json:"amount"`
}

type TransactionFeeConfig interface {
	// FeesByEquityType - map of EquityType to TransactionFee structure
	FeesByEquityType() map[equity_type.EquityType]*TransactionFee
}