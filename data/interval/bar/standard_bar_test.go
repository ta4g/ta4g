package bar

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestInMemoryBar(t *testing.T) {
	now := time.Now()
	ptrOpen := 100.456
	ptrHigh := 200.456
	ptrLow := 50.456
	ptrClose := 150.456
	ptrVolume := 123.0
	ptrOpenInterest := int64(500)

	bar := New(
		now,
		ptrOpen,
		ptrHigh,
		ptrLow,
		ptrClose,
		ptrVolume,
		ptrOpenInterest,
	)
	require.Equal(t, bar.GetTime().Unix(), now.Unix())
	require.Equal(t, bar.GetOpen(), ptrOpen)
	require.Equal(t, bar.GetHigh(), ptrHigh)
	require.Equal(t, bar.GetLow(), ptrLow)
	require.Equal(t, bar.GetClose(), ptrClose)
	require.Equal(t, bar.GetVolume(), ptrVolume)
	require.Equal(t, bar.GetOpenInterest(), ptrOpenInterest)

	bar = New(
		now,
		ptrOpen,
		ptrHigh,
		ptrLow,
		ptrClose,
		ptrVolume,
		-1,
	)
	require.Equal(t, bar.GetOpenInterest(), int64(-1))
}
