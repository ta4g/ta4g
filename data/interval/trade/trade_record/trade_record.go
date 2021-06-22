package trade_record

import (
	"github.com/ta4g/ta4g/data/interval/trade/constants/equity_type"
	"github.com/ta4g/ta4g/data/interval/trade/constants/errors"
	"github.com/ta4g/ta4g/data/interval/trade/constants/trade_direction"
	"github.com/ta4g/ta4g/data/interval/trade/transaction_fee"
)

type TradeRecord struct {
	// TransactionGroup - What is the unique key for this transaction grouping?
	// eg a CoveredCall or Option Vertical will have a key that groups the Stock + Option trade together.
	TransactionGroup string `csv:"transaction_group" avro:"transaction_group" json:"transaction_group"`

	// TradeDirection - are we buying or selling?
	trade_direction.TradeDirection `csv:"trade_direction" avro:"trade_direction" json:"trade_direction"`

	// EquityType - what type of item is this?
	equity_type.EquityType `csv:"item_type" avro:"item_type" json:"item_type"`

	// Symbol or ID of the item we are buying or selling
	// For options and other derivatives we also have two additional fields: StrikePrice and ExpirationDate
	Symbol string `csv:"symbol" avro:"symbol" json:"symbol"`

	// StrikePrice - What is the price at which this derivative can be exercised?
	StrikePrice float64 `csv:"strike_price" avro:"strike_price" json:"strike_price"`

	// ExpirationDate - What is the date when this derivative expires and either becomes exercised or expires worthless?
	ExpirationDate int64 `csv:"expiration_date" avro:"expiration_date" json:"expiration_date"`

	// Amount - How many of the items are you buying or selling?
	// Ex 100 shares of stock, or 1x option contract
	Amount float64 `csv:"amount" avro:"amount" json:"amount"`

	// QuantityPerAmount - how many items are included amount we buy or sell?
	// Stock == 1
	// Option == 10, 100, or some other customer number
	// Crypto == Variable, depending on the currency
	QuantityPerAmount float64 `csv:"quantity_per_amount" avro:"quantity_per_amount" json:"quantity_per_amount"`

	// Price - How much is it per item?
	// 1. A stock is $10 so it's $10 per share
	// 2. An option is $1 per 100 shares, so it's $100 for the contract
	Price float64 `csv:"price" avro:"price" json:"price"`

	// NetPrice - How much are we do we pay or are paid?
	NetPrice float64 `csv:"net_price" avro:"net_price" json:"net_price"`

	// MarginMaintenance - How much margin is this position consuming?
	MarginMaintenance float64 `csv:"margin_maintenance" avro:"margin_maintenance" json:"margin_maintenance"`
}

func NewCashOrderItem(
	transactionFeeConfig transaction_fee.TransactionFeeConfig,
	direction trade_direction.TradeDirection,
	amount float64,
) (*TradeRecord, error) {
	transactionRecord := &TradeRecord{
		Amount:            amount,
		EquityType:        equity_type.Cash,
		Price:             1.0,
		QuantityPerAmount: 1.0,
		NetPrice:          amount,
		Symbol:            "USD",
		TradeDirection:    direction,
	}
	feeType, ok := transactionFeeConfig.FeesByEquityType()[transactionRecord.EquityType]
	if !ok {
		return nil, errors.InvalidArgument
	}
	transactionRecord.applyNetPrice(feeType)
	return transactionRecord, nil

}

func NewStockOrderItem(
	transactionFeeConfig transaction_fee.TransactionFeeConfig,
	direction trade_direction.TradeDirection,
	symbol string,
	amount, price float64,
) (*TradeRecord, error) {
	transactionRecord := &TradeRecord{
		Amount:            amount,
		EquityType:        equity_type.Stock,
		Price:             price,
		QuantityPerAmount: 1.0,
		Symbol:            symbol,
		TradeDirection:    direction,
	}
	feeType, ok := transactionFeeConfig.FeesByEquityType()[transactionRecord.EquityType]
	if !ok {
		return nil, errors.InvalidArgument
	}
	transactionRecord.applyNetPrice(feeType)
	return transactionRecord, nil
}

func NewOptionOrderItem(
	transactionFeeConfig transaction_fee.TransactionFeeConfig,
	direction trade_direction.TradeDirection,
	symbol string,
	expirationDate int64,
	strikePrice, amount, price float64,
) (*TradeRecord, error) {
	transactionRecord := &TradeRecord{
		Amount:            amount,
		EquityType:        equity_type.Option,
		ExpirationDate:    expirationDate,
		Price:             price,
		QuantityPerAmount: 100,
		StrikePrice:       strikePrice,
		Symbol:            symbol,
		TradeDirection:    direction,
	}
	feeType, ok := transactionFeeConfig.FeesByEquityType()[transactionRecord.EquityType]
	if !ok {
		return nil, errors.InvalidArgument
	}

	transactionRecord.applyNetPrice(feeType)
	return transactionRecord, nil

}

func NewCryptoOrderItem(
	transactionFeeConfig transaction_fee.TransactionFeeConfig,
	direction trade_direction.TradeDirection,
	symbol string,
	amount, price float64,
) (*TradeRecord, error) {
	transactionRecord := &TradeRecord{
		Amount:            amount,
		EquityType:        equity_type.Crypto,
		Price:             price,
		QuantityPerAmount: 1.0,
		Symbol:            symbol,
		TradeDirection:    direction,
	}
	feeType, ok := transactionFeeConfig.FeesByEquityType()[transactionRecord.EquityType]
	if !ok {
		return nil, errors.InvalidArgument
	}
	transactionRecord.applyNetPrice(feeType)
	return transactionRecord, nil
}

func (s *TradeRecord) Clone() *TradeRecord {
	output := *s
	return &output
}

func (s *TradeRecord) Opposite() *TradeRecord {
	output := *s
	output.TradeDirection = output.TradeDirection.Opposite()
	return &output
}

func (s *TradeRecord) applyNetPrice(fee *transaction_fee.TransactionFee) {
	// The amount we are paying or getting paid
	output := s.Amount * s.QuantityPerAmount * s.Price
	// Per-contract or per-share transaction_fee we must pay
	output += s.Amount * fee.Amount
	// Exchange transaction_fee per order-item
	output += fee.Exchange
	// Finally, per-order fee for the broker
	output += fee.Order

	// When we are buying we need to pay
	// When we are selling we are getting paid
	if s.TradeDirection == trade_direction.Buy {
		output *= -1.0
	}

	s.NetPrice = output
}

func (s *TradeRecord) CalculatePrice(exchangeFee, perOrderFee, perUnitFee float64) float64 {
	// The amount we are paying or getting paid
	output := s.Amount * s.QuantityPerAmount * s.Price
	// Per-contract or per-share transaction_fee we must pay
	output += s.Amount * perUnitFee
	// Exchange transaction_fee per order-item
	output += exchangeFee
	// Finally, per-order fee for the broker
	output += perOrderFee

	// When we are buying we need to pay
	// When we are selling we are getting paid
	if s.TradeDirection == trade_direction.Buy {
		output *= -1.0
	}

	return output
}

func (s *TradeRecord) MarginRequirement() float64 {
	switch s.EquityType {
	case equity_type.Cash:
		return 0.0
	case equity_type.Stock:
		fallthrough
	case equity_type.Option:
		fallthrough
	case equity_type.Crypto:
		output := s.Amount * s.QuantityPerAmount * s.Price
		// When we are buying we need to pay
		// When we are selling we are getting paid, but that also uses up margin
		if s.TradeDirection == trade_direction.Sell {
			output *= -1.0
		}
		return output
	}
	return 0.0
}
