package bar

import (
	"math/big"
	"time"
)

// Compile time type assertion
var _ Bar = &InMemoryBar{}

type InMemoryBar struct {
	barTime         time.Time
	barOpen         *big.Float
	barHigh         *big.Float
	barLow          *big.Float
	barClose        *big.Float
	barVolume       *big.Int
	barOpenInterest *big.Int
}

func New(
	barTime time.Time,
	barOpen,
	barHigh,
	barLow,
	barClose big.Float,
	barVolume big.Int,
	barOpenInterest *big.Int,
) Bar {
	return &InMemoryBar{
		barTime:         barTime,
		barOpen:         &barOpen,
		barHigh:         &barHigh,
		barLow:          &barLow,
		barClose:        &barClose,
		barVolume:       &barVolume,
		barOpenInterest: copyPreciseInt(barOpenInterest),
	}
}

func (i InMemoryBar) Time() time.Time {
	return i.barTime
}

func (i InMemoryBar) Open() *big.Float {
	return copyPreciseFloat(i.barOpen)
}

func (i InMemoryBar) High() *big.Float {
	return copyPreciseFloat(i.barHigh)
}

func (i InMemoryBar) Low() *big.Float {
	return copyPreciseFloat(i.barLow)
}

func (i InMemoryBar) Close() *big.Float {
	return copyPreciseFloat(i.barClose)
}

func (i InMemoryBar) Volume() *big.Int {
	return copyPreciseInt(i.barVolume)
}

func (i InMemoryBar) OpenInterest() *big.Int {
	return copyPreciseInt(i.barOpenInterest)
}

func copyPreciseFloat(input *big.Float) *big.Float {
	// Nothing for input, nothing for output
	if nil == input {
		return nil
	}
	// Return a copy
	output := big.NewFloat(0)
	output.Set(input)
	return output
}

func copyPreciseInt(input *big.Int) *big.Int {
	// Nothing for input, nothing for output
	if nil == input {
		return nil
	}
	// Return a copy
	output := big.NewInt(0)
	output.Set(input)
	return output
}
