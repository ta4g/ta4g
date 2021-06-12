package orders

import (
	"time"
)

// Order represents a collection of the items that are purchased or sold in a single batch
type Order struct {
	// UnixTime the order was placed, for back-testing we will assume all orders are filled immediately.
	UnixTime int64 `csv:"time" avro:"time" json:"time"`
	// OrderItems are all of the items that are purchased or sold
	OrderItems []*OrderItem `csv:"items" avro:"items" json:"items"`
}

func NewOrder(t time.Time, items ...*OrderItem) *Order {
	output := &Order{
		UnixTime:   t.Unix(),
		OrderItems: make([]*OrderItem, 0, len(items)),
	}
	for _, item := range items {
		output.OrderItems = append(output.OrderItems, item.Clone())
	}
	return output
}

func (s *Order) Append(item *OrderItem) {
	s.OrderItems = append(s.OrderItems, item.Clone())
}

func (s *Order) AddItem(index int, item *OrderItem) {
	items := make([]*OrderItem, 0, len(s.OrderItems)+1)
	items = append(items, s.OrderItems[0:index]...)
	items = append(items, item.Clone())
	if index < len(s.OrderItems) {
		items = append(items, s.OrderItems[index:]...)
	}
	s.OrderItems = items
}

func (s *Order) RemoveItem(index int) {
	items := make([]*OrderItem, 0, len(s.OrderItems)-1)
	items = append(items, s.OrderItems[0:index]...)
	if index < len(s.OrderItems) {
		items = append(items, s.OrderItems[index+1:]...)
	}
	s.OrderItems = items
}

func (s *Order) Clone() *Order {
	return NewOrder(
		time.Unix(s.UnixTime, 0),
		s.OrderItems...,
	)
}
