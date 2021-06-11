package bar

import (
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
	"time"
)

func TestInMemoryBar(t *testing.T) {
	now := time.Now()
	ptrOpen := big.NewFloat(100.456)
	ptrHigh := big.NewFloat(200.456)
	ptrLow := big.NewFloat(50.456)
	ptrClose := big.NewFloat(150.456)
	ptrVolume := big.NewInt(123)
	ptrOpenInterest := big.NewInt(500)

	bar := New(
		now,
		*ptrOpen,
		*ptrHigh,
		*ptrLow,
		*ptrClose,
		*ptrVolume,
		ptrOpenInterest,
	)
	require.Equal(t, bar.Time().String(), now.String())

	require.NotNil(t, bar.Open())
	require.Equal(t, bar.Open().String(), ptrOpen.String())

	require.NotNil(t, bar.High())
	require.Equal(t, bar.High().String(), ptrHigh.String())

	require.NotNil(t, bar.Low())
	require.Equal(t, bar.Low().String(), ptrLow.String())

	require.NotNil(t, bar.Close())
	require.Equal(t, bar.Close().String(), ptrClose.String())

	require.NotNil(t, bar.Volume())
	require.Equal(t, bar.Volume().String(), ptrVolume.String())

	require.NotNil(t, bar.OpenInterest())
	require.Equal(t, bar.OpenInterest().String(), ptrOpenInterest.String())

	bar = New(
		now,
		*ptrOpen,
		*ptrHigh,
		*ptrLow,
		*ptrClose,
		*ptrVolume,
		nil,
	)
	require.Nil(t, bar.OpenInterest())
}
