package trade_record

import (
	"github.com/ta4g/ta4g/data/interval/trade/constants/equity_type"
	"github.com/ta4g/ta4g/data/interval/trade/constants/trade_direction"
	"github.com/ta4g/ta4g/gen/proto/interval/trade_record"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type TradeRecord = trade_record.TradeRecord

func NewCashOrderItem(
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
	applyNetPrice(transactionRecord)
	return transactionRecord, nil

}

func NewStockOrderItem(
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
	applyNetPrice(transactionRecord)
	return transactionRecord, nil
}

func NewOptionOrderItem(
	direction trade_direction.TradeDirection,
	symbol string,
	expirationDate int64,
	strikePrice, amount, price float64,
) (*TradeRecord, error) {
	transactionRecord := &TradeRecord{
		Amount:            amount,
		EquityType:        equity_type.Option,
		ExpirationDate:    timestamppb.New(time.Unix(expirationDate, 0)),
		Price:             price,
		QuantityPerAmount: 100,
		StrikePrice:       strikePrice,
		Symbol:            symbol,
		TradeDirection:    direction,
	}
	applyNetPrice(transactionRecord)
	return transactionRecord, nil

}

func NewCryptoOrderItem(
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
	applyNetPrice(transactionRecord)
	return transactionRecord, nil
}

func Clone(s *TradeRecord) *TradeRecord {
	output := *s
	return &output
}

func Opposite(s *TradeRecord, ) *TradeRecord {
	output := &TradeRecord{
		TradeDirection:    trade_direction.Opposite(s.TradeDirection),
		EquityType:        s.EquityType,
		Symbol:            s.Symbol,
		StrikePrice:       s.StrikePrice,
		ExpirationDate:    s.ExpirationDate,
		Amount:            s.Amount,
		QuantityPerAmount: s.QuantityPerAmount,
		Price:             s.Price,
	}
	applyNetPrice(output)
	return output
}

func applyNetPrice(s *TradeRecord) {
	// The amount we are paying or getting paid
	output := s.Amount * s.QuantityPerAmount * s.Price

	// When we are buying we need to pay
	// When we are selling we are getting paid
	if s.TradeDirection == trade_direction.Buy {
		output *= -1.0
	}

	s.NetPrice = output
}

//func CalculatePrice(s *TradeRecord, exchangeFee, perOrderFee, perUnitFee float64) float64 {
//	// The amount we are paying or getting paid
//	output := s.Amount * s.QuantityPerAmount * s.Price
//	// Per-contract or per-share transaction_fee we must pay
//	output += s.Amount * perUnitFee
//	// Exchange transaction_fee per order-item
//	output += exchangeFee
//	// Finally, per-order fee for the broker
//	output += perOrderFee
//
//	// When we are buying we need to pay
//	// When we are selling we are getting paid
//	if s.TradeDirection == trade_direction.Buy {
//		output *= -1.0
//	}
//
//	return output
//}
