package rates

type RateRange struct {
	Min, Max float64
	Rate     float64
}

// StandardRateRanges these are the standard rates as of 2021-06 as posted on the TD Ameritrade price sheet
// > https://www.tdameritrade.com/pricing/margin-and-interest-rates.html
var StandardRateRanges = []RateRange{
	{0, 10000, 9.50},
	{10000.00, 24999.99, 9.25},
	{25000.00, 49999.99, 9.00},
	{50000.00, 99999.99, 8.00},
	{100000.00, 249999.99, 7.75},
	{250000.00, 499999.99, 7.50},
}

func GetMarginInterestRate(value float64) float64 {
	if value <= 0 {
		value = 0.0
	}
	for _, rateRange := range StandardRateRanges {
		if rateRange.Min <= value && value <= rateRange.Max {
			return rateRange.Rate
		}
	}
	return 7.50
}
