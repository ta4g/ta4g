package constants

type OrderType int

const (
	_ OrderType = iota
	EnterOrderType
	ExitOrderType
	AdjustmentOrderType
)
const (
	enterOrderTypeStr      = "enter"
	exitOrderTypeStr       = "exit"
	adjustmentOrderTypeStr = "adjustment"
)

var orderTypes = map[OrderType]string{
	EnterOrderType:      enterOrderTypeStr,
	ExitOrderType:       exitOrderTypeStr,
	AdjustmentOrderType: adjustmentOrderTypeStr,
}

func (o OrderType) String() string {
	return orderTypes[o]
}
