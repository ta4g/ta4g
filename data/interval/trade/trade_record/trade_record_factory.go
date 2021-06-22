package trade_record

import (
	"github.com/ta4g/ta4g/data/interval/trade/constants/equity_type"
	"github.com/ta4g/ta4g/data/interval/trade/constants/errors"
	"github.com/ta4g/ta4g/data/interval/trade/constants/trade_direction"
	"github.com/ta4g/ta4g/data/interval/trade/transaction_fee"
)

type TradeRecordFactory struct {
	ExchangeFee  float64
	OrderFee     float64
	PerUnitFee   float64
	Symbol       string
	TradeRecords []*TradeRecord
	transaction_fee.TransactionFeeConfig
}

type CurrencyParams struct {
	Amount float64
}
type StockParams struct {
	Amount, Price float64
}
type OptionParams struct {
	StrikePrice, Amount, Price float64
}

func (t *TradeRecordFactory) AddFunds(params CurrencyParams) error {
	if params.Amount < 0.0 {
		return errors.InvalidAmount
	}
	cash, err := NewCashOrderItem(
		t.TransactionFeeConfig,
		trade_direction.Neutral,
		params.Amount,
	)
	if nil != err {
		return err
	}

	t.TradeRecords = append(t.TradeRecords, cash)
	return nil
}

func (t *TradeRecordFactory) RemoveFunds(params CurrencyParams) error {
	if params.Amount < 0.0 {
		return errors.InvalidAmount
	}
	if params.Amount >= t.CurrentFunds() {
		return errors.InsufficientFunds
	}
	cash, err := NewCashOrderItem(
		t.TransactionFeeConfig,
		trade_direction.Neutral,
		-1*params.Amount,
	)
	if nil != err {
		return err
	}
	t.TradeRecords = append(t.TradeRecords, cash)
	return nil
}

func (t *TradeRecordFactory) BuyStock(params StockParams) error {
	transactionRecord, err := NewStockOrderItem(
		t.TransactionFeeConfig, trade_direction.Buy, t.Symbol, params.Amount, params.Price,
	)
	if nil != err {
		return err
	}

	err = t.RemoveFunds(CurrencyParams{Amount: transactionRecord.NetPrice})
	if nil != err {
		return err
	}

	t.TradeRecords = append(t.TradeRecords, transactionRecord)
	return nil
}

func (t *TradeRecordFactory) SellStock(params StockParams) error {
	transactionRecord, err := NewStockOrderItem(
		t.TransactionFeeConfig, trade_direction.Sell, t.Symbol, params.Amount, params.Price,
	)
	if nil != err {
		return err
	}

	err = t.AddFunds(CurrencyParams{Amount: transactionRecord.NetPrice})
	if nil != err {
		return err
	}

	t.TradeRecords = append(t.TradeRecords, transactionRecord)
	return nil
}

func (t *TradeRecordFactory) BuyOption(expiration int64, params OptionParams) error {
	transactionRecord, err := NewOptionOrderItem(
		t.TransactionFeeConfig, trade_direction.Buy, t.Symbol, expiration, params.StrikePrice, params.Amount, params.Price,
	)
	if nil != err {
		return err
	}

	err = t.RemoveFunds(CurrencyParams{Amount: transactionRecord.NetPrice})
	if nil != err {
		return err
	}

	t.TradeRecords = append(t.TradeRecords, transactionRecord)
	return nil
}

func (t *TradeRecordFactory) SellOption(expiration int64, params OptionParams) error {
	transactionRecord, err := NewOptionOrderItem(
		t.TransactionFeeConfig, trade_direction.Sell, t.Symbol, expiration, params.StrikePrice, params.Amount, params.Price,
	)
	if nil != err {
		return err
	}

	err = t.AddFunds(CurrencyParams{Amount: transactionRecord.NetPrice})
	if nil != err {
		return err
	}

	t.TradeRecords = append(t.TradeRecords, transactionRecord)
	return nil
}

func (t *TradeRecordFactory) BuyOptionVertical(expiration int64, strikePrice, amount, price []float64) error {
	transactionRecord, err := NewOptionOrderItem(
		t.TransactionFeeConfig, trade_direction.Buy, t.Symbol, expiration, strikePrice, amount, price,
	)
	if nil != err {
		return err
	}

	err = t.RemoveFunds(transactionRecord.NetPrice)
	if nil != err {
		return err
	}

	t.TradeRecords = append(t.TradeRecords, transactionRecord)
	return nil
}

func (t *TradeRecordFactory) SellOptions(expiration int64, strikePrices, amounts, prices []float64) error {
	transactionRecord, err := NewOptionOrderItem(
		t.TransactionFeeConfig, trade_direction.Sell, t.Symbol, expiration, strikePrice, amount, price,
	)
	if nil != err {
		return err
	}

	err = t.AddFunds(transactionRecord.NetPrice)
	if nil != err {
		return err
	}

	t.TradeRecords = append(t.TradeRecords, transactionRecord)
	return nil
}

func (t *TradeRecordFactory) BuyCrypto(symbol string, amount, price float64) error {
	transactionRecord, err := NewCryptoOrderItem(
		t.TransactionFeeConfig, trade_direction.Buy, t.Symbol, amount, price,
	)
	if nil != err {
		return err
	}

	err = t.RemoveFunds(transactionRecord.NetPrice)
	if nil != err {
		return err
	}

	t.TradeRecords = append(t.TradeRecords, transactionRecord)
	return nil
}

func (t *TradeRecordFactory) SellCrypto(symbol string, amount, price float64) error {
	transactionRecord, err := NewCryptoOrderItem(
		t.TransactionFeeConfig, trade_direction.Sell, t.Symbol, amount, price,
	)
	if nil != err {
		return err
	}

	err = t.AddFunds(transactionRecord.NetPrice)
	if nil != err {
		return err
	}

	t.TradeRecords = append(t.TradeRecords, transactionRecord)
	return nil
}

func (t *TradeRecordFactory) CurrentFunds() float64 {
	var output float64
	for _, tradeRecord := range t.TradeRecords {
		if tradeRecord.EquityType == equity_type.Cash {
			output += tradeRecord.NetPrice
		}
	}
	return output
}
