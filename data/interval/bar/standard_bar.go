package bar

import (
	"math"
	"math/rand"
	"time"
)

// Compile time type assertion
var _ Bar = &StandardBar{}

type StandardBar struct {
	UnixTime     int64   `csv:"time" avro:"time"`
	Open         float64 `csv:"open" avro:"open"`
	High         float64 `csv:"high" avro:"high"`
	Low          float64 `csv:"low" avro:"low"`
	Close        float64 `csv:"close" avro:"close"`
	Volume       float64 `csv:"volume" avro:"volume"`
	OpenInterest int64   `csv:"open_interest" avro:"open_interest"`
}

func New(
	t time.Time,
	openValue,
	highValue,
	lowValue,
	closeValue,
	volumeValue float64,
	openInterestValue int64,
) Bar {
	return &StandardBar{
		UnixTime:     t.Unix(),
		Open:         openValue,
		High:         highValue,
		Low:          lowValue,
		Close:        closeValue,
		Volume:       volumeValue,
		OpenInterest: openInterestValue,
	}
}

func copyToStandardBar(input Bar) *StandardBar {
	t := input.GetTime()
	return &StandardBar{
		UnixTime:     t.Unix(),
		Open:         input.GetOpen(),
		High:         input.GetHigh(),
		Low:          input.GetLow(),
		Close:        input.GetClose(),
		Volume:       input.GetVolume(),
		OpenInterest: input.GetOpenInterest(),
	}
}

func NewFakeBar(t time.Time) Bar {
	seed := math.Abs(rand.Float64())
	volume := math.Abs(float64(rand.Int()))
	return &StandardBar{
		UnixTime:     t.Unix(),
		Open:         seed * 0.25,
		High:         seed * 0.9,
		Low:          seed * 0.1,
		Close:        seed * 0.75,
		Volume:       volume,
		OpenInterest: int64(volume * 0.25),
	}

}

func (i StandardBar) GetTime() time.Time {
	return time.Unix(i.UnixTime, 0)
}

func (i StandardBar) GetOpen() float64 {
	return i.Open
}

func (i StandardBar) GetHigh() float64 {
	return i.High
}

func (i StandardBar) GetLow() float64 {
	return i.Low
}

func (i StandardBar) GetClose() float64 {
	return i.Close
}

func (i StandardBar) GetVolume() float64 {
	return i.Volume
}

func (i StandardBar) GetOpenInterest() int64 {
	return i.OpenInterest
}

func (i StandardBar) Clone() (Bar, error) {
	return &StandardBar{
		UnixTime:     i.UnixTime,
		Open:         i.Open,
		High:         i.High,
		Low:          i.Low,
		Close:        i.Close,
		Volume:       i.Volume,
		OpenInterest: i.OpenInterest,
	}, nil
}
