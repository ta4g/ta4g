package bar

import (
	"math/big"
	"time"
)

// Bar is modeled after the standard OHLC (Open/High/Low/Close) bar used commonly in trading platforms.
//
// This is a simple wrapper around some basic price movement over a period of time,
// and allows us to understand what happened during in interval.
//
type Bar interface {
	// Time interval of this bar
	Time() time.Time

	// Open price for the current bar
	Open() *big.Float

	// High price for the current bar
	High() *big.Float

	// Low price for the current bar
	Low() *big.Float

	// Close price for the current bar
	Close() *big.Float

	// Volume of shares, options, coins, etc traded during this bar
	Volume() *big.Int

	// OpenInterest (Optional) amount of derivatives currently outstanding for this bar
	OpenInterest() *big.Int
}
