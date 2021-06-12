package orders

type OrderItem struct {
	// Direction - are we buying or selling?
	Direction `csv:"direction" avro:"direction" json:"direction"`

	// ItemType - what type of item is this?
	ItemType `csv:"item_type" avro:"item_type" json:"item_type"`

	// Symbol or ID of the item we are buying or selling
	Symbol string `csv:"symbol" avro:"symbol" json:"symbol"`

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
}

func NewUSDOrderItem(direction Direction, symbol string, amount, price float64) *OrderItem {
	return &OrderItem{
		Direction:         direction,
		ItemType:          USD,
		Symbol:            symbol,
		Amount:            amount,
		QuantityPerAmount: 1.0,
		Price:             price,
	}
}

func NewStockOrderItem(direction Direction, symbol string, amount, price float64) *OrderItem {
	return &OrderItem{
		Direction:         direction,
		ItemType:          Stock,
		Symbol:            symbol,
		Amount:            amount,
		QuantityPerAmount: 1.0,
		Price:             price,
	}
}

func NewOptionOrderItem(direction Direction, symbol string, amount, price float64) *OrderItem {
	return &OrderItem{
		Direction:         direction,
		ItemType:          Option,
		Symbol:            symbol,
		Amount:            amount,
		QuantityPerAmount: 100,
		Price:             price,
	}
}

func NewCryptoOrderItem(direction Direction, symbol string, amount, price float64) *OrderItem {
	return &OrderItem{
		Direction:         direction,
		ItemType:          Crypto,
		Symbol:            symbol,
		Amount:            amount,
		QuantityPerAmount: 1.0,
		Price:             price,
	}
}

func (s *OrderItem) Clone() *OrderItem {
	output := *s
	return &output
}

func (s *OrderItem) CalculatePrice(exchangeFee, perOrderFee, perUnitFee float64) float64 {
	// The amount we are paying or getting paid
	output := s.Amount * s.QuantityPerAmount * s.Price
	// Per-contract or per-share fees we must pay
	output += s.Amount * perUnitFee
	// Exchange fees per order-item
	output += exchangeFee
	// Finally, per-order fee for the broker
	output += perOrderFee

	// When we are buying we need to pay
	// When we are selling we are getting paid
	if s.Direction == Buy {
		output *= -1.0
	}

	return output
}

func (s *OrderItem) MarginRequirement() float64 {
	switch s.ItemType {
	case USD:
		return 0.0
	case Stock:
		fallthrough
	case Option:
		fallthrough
	case Crypto:
		output := s.Amount * s.QuantityPerAmount * s.Price
		// When we are buying we need to pay
		// When we are selling we are getting paid, but that also uses up margin
		if s.Direction == Sell {
			output *= -1.0
		}
		return output
	}
	return 0.0
}
