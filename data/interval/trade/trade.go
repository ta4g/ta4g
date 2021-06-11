package trade

type Trade interface {
	GetEntry() Order
	GetAdjustments() Order
	GetExit() Order
	GetOrderCostModel() CostModel
	GetHoldingCostModel() CostModel
}

