package trade

import "time"

type Order interface {
	GetTime() time.Time
	GetItems() []OrderItem
	AddItem(index int, orderItem OrderItem)
	RemoveItem(index int)
	Clone() (Order, error)
}

type OrderItem interface {
	GetOrderDirection() OrderDirection
	GetSymbol() string
	GetIsOption() bool
	GetUnitQuantity() float64
	GetPricePerUnit() float64
	GetNetCost(exchangeFee, perOrderFee, perUnitFee float64) float64
}

// Compile time assertions
var _ Order = &StandardOrder{}
var _ OrderItem = &StandardOrderItem{}

//
// StandardOrder
//

type StandardOrder struct {
	Time  int64                `csv:"time" avro:"time" json:"time"`
	Items []*StandardOrderItem `csv:"items" avro:"items" json:"items"`
}

func NewStandardOrder(t time.Time, items ...OrderItem) Order {
	output := &StandardOrder{
		Time:  t.Unix(),
		Items: make([]*StandardOrderItem, 0, len(items)),
	}
	for _, item := range items {
		output.Items = append(output.Items, copyToOrderItem(item))
	}
	return output
}

func copyToStandardOrder(order Order) *StandardOrder {
	output := NewStandardOrder(order.GetTime(), order.GetItems()...)
	stdOrder, _ := output.(*StandardOrder)
	return stdOrder
}

func (s *StandardOrder) GetTime() time.Time {
	return time.Unix(s.Time, 0)
}

func (s *StandardOrder) GetItems() []OrderItem {
	output := make([]OrderItem, 0, len(s.Items))
	for _, item := range s.Items {
		output = append(output, item)
	}
	return output
}

func (s *StandardOrder) AddItem(index int, orderItem OrderItem) {
	stdItem := copyToOrderItem(orderItem)

	items := make([]*StandardOrderItem, 0, len(s.Items)+1)
	items = append(items, s.Items[0:index]...)
	items = append(items, stdItem)
	items = append(items, s.Items[index:]...)
	s.Items = items
}

func (s *StandardOrder) RemoveItem(index int) {
	items := make([]*StandardOrderItem, 0, len(s.Items)-1)
	items = append(items, s.Items[0:index]...)
	if index < len(s.Items) {
		items = append(items, s.Items[index+1:]...)
	}
	s.Items = items
}

func (s *StandardOrder) Clone() (Order, error) {
	return NewStandardOrder(s.GetTime(), s.GetItems()...), nil
}

//
// StandardOrderItem implements the OrderItem interface
//

type StandardOrderItem struct {
	OrderDirection `csv:"order_direction" avro:"order_direction" json:"order_direction"`
	Symbol         string  `csv:"symbol" avro:"symbol" json:"symbol"`
	IsOption       bool    `csv:"is_option" avro:"is_option" json:"is_option"`
	UnitQuantity   float64 `csv:"unit_quantity" avro:"unit_quantity" json:"unit_quantity"`
	PricePerUnit   float64 `csv:"price_per_unit" avro:"price_per_unit" json:"price_per_unit"`
}

func NewOrderItem(orderDirection OrderDirection, symbol string, IsOption bool, unitQuantity, pricePerUnit float64) OrderItem {
	return &StandardOrderItem{
		OrderDirection: orderDirection,
		Symbol:         symbol,
		IsOption:       IsOption,
		UnitQuantity:   unitQuantity,
		PricePerUnit:   pricePerUnit,
	}
}

func copyToOrderItem(orderItem OrderItem) *StandardOrderItem {
	if v, ok := orderItem.(*StandardOrderItem); ok {
		output := *v
		return &output
	}

	return &StandardOrderItem{
		OrderDirection: orderItem.GetOrderDirection(),
		Symbol:         orderItem.GetSymbol(),
		IsOption:       orderItem.GetIsOption(),
		UnitQuantity:   orderItem.GetUnitQuantity(),
		PricePerUnit:   orderItem.GetPricePerUnit(),
	}
}

func (s *StandardOrderItem) GetOrderDirection() OrderDirection {
	return s.OrderDirection
}

func (s *StandardOrderItem) GetSymbol() string {
	return s.Symbol
}

func (s *StandardOrderItem) GetIsOption() bool {
	return s.IsOption
}

func (s *StandardOrderItem) GetUnitQuantity() float64 {
	return s.UnitQuantity
}

func (s *StandardOrderItem) GetPricePerUnit() float64 {
	return s.PricePerUnit
}

func (s *StandardOrderItem) GetNetCost(exchangeFee, perOrderFee, perUnitFee float64) float64 {
	output := s.PricePerUnit * s.UnitQuantity
	output += s.UnitQuantity * perUnitFee
	output += exchangeFee
	output += perOrderFee

	// When we are buying we need to pay
	// When we are selling we are getting paid
	if s.GetOrderDirection() == BuyOrderDirection {
		output *= -1.0
	}

	return output
}
