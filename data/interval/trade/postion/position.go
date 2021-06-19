package postion

import "sort"

type Position struct {
	Key string
	Orders
}

//func (p *Position) OpeningCost() map[string]float64 {
//	p.Orders.GetEntries()[0].
//}

func (p *Position) Symbols() []string {
	symbols := make(map[string]bool, len(p.Orders))
	for _, o := range p.Orders {
		for _, v := range o.OrderItems {
			symbols[v.Symbol] = true
		}
	}

	output := make([]string, 0, len(symbols))
	for symbol, _ := range symbols {
		output = append(output, symbol)
	}

	sort.Strings(output)

	return output
}
