package constants

// EquityType defines each of the types of items we can add to a transaction.
//
// This is used to create a standardized "order" and "portfolio" model around simple transaction types:
// 1. Open Portfolio == Add `Cash` item to the portfolio
// 2. Purchase Stock == Add `Stock` item to the portfolio + Subtract `Cache` from the portfolio
// 3. SellDirection Stock = Remove `Stock` and add `Cache`
//
// This allows us to make the portfolio a running balance of all the transactions within the portfolio.
//
// This makes the current balance is simply a projection of the transactions against
// the current value of the Stock/Option/Crypto items that are currently open.
type EquityType int

const (
	minEquityType EquityType = iota

	// Cash for tracking the hard currency within an account,
	// this is the outstanding balance of fiat currency you can spend.
	//
	// When a "Portfolio" is created we allocate some amount of cash up-front.
	//
	// This is to give the Portfolio a concrete starting point and to make both
	// `CurrentPortfolio` and `ReturnOnCapital` much simpler to calculate later on.
	//
	// The amount of Cash you have also impacts the margin balance in your account.
	//
	// Regulation T Margin is the amount of `Cash` you need to set aside to guarantee your portfolio is stable.
	// > https://tickertape.tdameritrade.com/trading/portfolio-margin-vs-reg-t-15526
	//
	// TL;DR: This is how much borrowing power the cash grants you when shorting or using option products.
	//
	// Similar to Cash this is added to the Portfolio when it is initially funded and is 1:1 with cash added to the account.
	// However unlike Cash this is not used in `CurrentPortfolio` and `ReturnOnCapital`,
	// it is merely used as a governor to limit how risky any Portfolio can be.
	//
	// Regulation T Margin Rules:
	// 1. Margin equity = Stock + (+/- cash balance)
	//
	// 2. Maintenance margin = 50% initial
	//
	// 3. 25% SRO requirements;
	//    Long equities = 25% requirement;
	//    Short equities = 30% requirement.
	//    *SRO - Self- Regulatory Organization -all securities and commodity exchanges in the United States
	//
	// 4. TD Ameritrade uses 30% minimum house maintenance requirement on long and short equities
	//
	// 5. Option requirements computed in real-time using FINRA rules and fixed percentages; please review Margin Handbook for details
	//
	// 6. Long options are not marginable and have 100% requirement
	//
	// For example:
	// 1. Fund Portfolio:
	//    - Add 30k Cash
	//    - Add 30k Margin
	// 2. BuyDirection 100x shares of stock at $10.00 per share, this costs `$1000.0` = `100 * $10.0`
	//    - Remove 1k Cash
	//    - Remove 0.3k Margin (rule 4, this is the max of rules 2-4)
	//    - Add 100x shares at cost-basis of $10.00
	// 3. SellDirection 100x shares of stock at $12.00 per share, this pays `$1200.0` = `100 * $12.0`
	//    - Remove 100x shares cost-basis of $10.00
	//    - Add 1.2k Cash
	//    - Add 0.5k Margin (0.3k from initial + 0.2k from profit)
	// 4. Net portfolio balance is:
	//    - 30.2k Cash
	//    - 30.2k Margin
	Cash

	// Stock is the standard share of stock you can purchase on an exchange.
	// These represent shares in a company, ETF, or other product.
	//
	// NOTE:
	//   The purchase or sale of an asset will have an impact on the amount of Cash and Margin
	//   you have in the Portfolio at any given time.
	Stock

	// Option is a type of derivative that allows you to "control"
	// more than 1x share of stock for a limited period of time.
	//
	// This is typically a `Call` or `Put` contract in units of 100x or 10x shares.
	// When a stock is split that will change to follow the number of shares that result.
	//
	// NOTE:
	//   The purchase or sale of an asset will have an impact on the amount of Cash and Margin
	//   you have in the Portfolio at any given time.
	Option

	// Crypto is any cryptocurrency that is available.
	// This could be any type cryptocurrency, the coin "name" is the symbol here.
	//
	// NOTE:
	//   The purchase or sale of an asset will have an impact on the amount of Cash and Margin
	//   you have in the Portfolio at any given time.
	Crypto

	maxEquityType
)

type TransactionBalance struct {
	AvailableCash     float64 // The amount of cash currently available to use.
	AvailableMargin   float64 // The amount of margin currently available to use.
	MaintenanceMargin float64 // How much margin we need to maintain to keep the position open
	CostBasisValue    float64 // The cost basis for all open positions.
	CurrentValue      float64 // The value for all open positions.
}

var itemTypes = map[EquityType]string{
	Cash:   "cash",
	Stock:  "stock",
	Option: "option",
	Crypto: "crypto",
}

func (i EquityType) String() string {
	return itemTypes[i]
}
