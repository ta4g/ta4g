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

type StockParam struct {
	Amount, Price float64
}

type OptionParams struct {
	Expiration                 int64
	StrikePrice, Amount, Price float64
}

type CrptoParam struct {
	Amount, Price float64
}

//
// Funding methods
//

func (t *TradeRecordFactory) AddFunds(params CurrencyParams) error {
	if params.Amount < 0.0 {
		return errors.InvalidAmount
	}
	cash, err := NewCashOrderItem(
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
		trade_direction.Neutral,
		-1*params.Amount,
	)
	if nil != err {
		return err
	}
	t.TradeRecords = append(t.TradeRecords, cash)
	return nil
}

//
// Stock methods
//

func (t *TradeRecordFactory) BuyStock(param StockParam) error {
	transactionRecord, err := NewStockOrderItem(
		trade_direction.Buy, t.Symbol, param.Amount, param.Price,
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

func (t *TradeRecordFactory) SellStock(param StockParam) error {
	transactionRecord, err := NewStockOrderItem(
		trade_direction.Sell, t.Symbol, param.Amount, param.Price,
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

//
// Option methods
//

func (t *TradeRecordFactory) BuyOption(param OptionParams) error {
	transactionRecord, err := NewOptionOrderItem(
		trade_direction.Buy, t.Symbol, param.Expiration, param.StrikePrice, param.Amount, param.Price,
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

func (t *TradeRecordFactory) SellOption(param OptionParams) error {
	transactionRecord, err := NewOptionOrderItem(trade_direction.Sell, t.Symbol, param.Expiration, param.StrikePrice, param.Amount, param.Price)
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

func (t *TradeRecordFactory) BuyCrypto(params CrptoParam) error {
	transactionRecord, err := NewCryptoOrderItem(trade_direction.Buy, t.Symbol, params.Amount, params.Price)
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

func (t *TradeRecordFactory) SellCrypto(params CrptoParam) error {
	transactionRecord, err := NewCryptoOrderItem(trade_direction.Sell, t.Symbol, params.Amount, params.Price)
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

func (t *TradeRecordFactory) CurrentFunds() float64 {
	var output float64
	for _, tradeRecord := range t.TradeRecords {
		if tradeRecord.EquityType == equity_type.Cash {
			output += tradeRecord.NetPrice
		}
	}
	return output
}
