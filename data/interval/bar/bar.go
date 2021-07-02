package bar

import (
	"github.com/ta4g/ta4g/gen/proto/interval/bar"
)

// Bar is modeled after the standard OHLC (Open/High/Low/Close) bar used commonly in trading platforms.
//
// This is a simple wrapper around some basic price movement over a period of time,
// and allows us to understand what happened during in interval.
//
type Bar = bar.OHLCBar
