package bar

import (
	"time"
)

// Bar is modeled after the standard OHLC (Open/High/Low/Close) bar used commonly in trading platforms.
//
// This is a simple wrapper around some basic price movement over a period of time,
// and allows us to understand what happened during in interval.
//
type Bar interface {
	// GetTime interval of this bar
	GetTime() time.Time

	// GetOpen price for the current bar
	GetOpen() float64

	// GetHigh price for the current bar
	GetHigh() float64

	// GetLow price for the current bar
	GetLow() float64

	// GetClose price for the current bar
	GetClose() float64

	// GetVolume of shares, options, coins, etc traded during this bar
	GetVolume() float64

	// GetOpenInterest (Optional) amount of derivatives currently outstanding for this bar
	// If there is no open interest then this will be -1, to indicate no data
	GetOpenInterest() int64

	// Clone makes a copy of the underlying bar
	Clone() (Bar, error)
}
