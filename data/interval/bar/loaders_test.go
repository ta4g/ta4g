package bar

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/require"
	"github.com/ta4g/ta4g/data/time/time_series"
	"strings"
	"testing"
	"time"
)

func TestCSVLoader(t *testing.T) {
	// December 1st, 2022
	now := time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC)

	bars := []Bar{
		NewFakeBar(now),
		NewFakeBar(now.Add(time_series.Day)),
		NewFakeBar(now.Add(2*time_series.Day)),
		NewFakeBar(now.Add(3*time_series.Day)),
		NewFakeBar(now.Add(4*time_series.Day)),
	}

	ctx := context.Background()
	loader := NewCSVLoader()

	buff := bytes.NewBuffer([]byte{})
	err := loader.Write(ctx, buff, bars)
	require.NoError(t, err)

	lines := strings.Split(buff.String(), "\n")
	require.Len(t, lines, len(bars)+2)
	require.Empty(t, lines[len(bars)+1]) // Last line is blank

	reader := bytes.NewReader(buff.Bytes())
	output, err := loader.Read(ctx, reader)
	require.NoError(t, err)
	require.Len(t, output, len(bars))
	for index, row := range output {
		b := bars[index]
		require.Equal(t, row.GetTime().String(), b.GetTime().String())
		require.Equal(t, row.GetOpen(), b.GetOpen())
		require.Equal(t, row.GetHigh(), b.GetHigh())
		require.Equal(t, row.GetLow(), b.GetLow())
		require.Equal(t, row.GetClose(), b.GetClose())
		require.Equal(t, row.GetVolume(), b.GetVolume())
		require.Equal(t, row.GetOpenInterest(), b.GetOpenInterest())
	}
}

func TestJsonNewLineLoader(t *testing.T) {
	// December 1st, 2022
	now := time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC)

	bars := []Bar{
		NewFakeBar(now),
		NewFakeBar(now.Add(time_series.Day)),
		NewFakeBar(now.Add(2*time_series.Day)),
		NewFakeBar(now.Add(3*time_series.Day)),
		NewFakeBar(now.Add(4*time_series.Day)),
	}

	ctx := context.Background()
	loader := NewJsonNewLineLoader()

	buff := bytes.NewBuffer([]byte{})
	err := loader.Write(ctx, buff, bars)
	require.NoError(t, err)

	lines := strings.Split(buff.String(), "\n")
	require.Len(t, lines, len(bars)+1)
	require.Empty(t, lines[len(bars)]) // Last line is blank

	reader := bytes.NewReader(buff.Bytes())
	output, err := loader.Read(ctx, reader)
	require.NoError(t, err)
	require.Len(t, output, len(bars))
	for index, row := range output {
		b := bars[index]
		require.Equal(t, row.GetTime().String(), b.GetTime().String())
		require.Equal(t, row.GetOpen(), b.GetOpen())
		require.Equal(t, row.GetHigh(), b.GetHigh())
		require.Equal(t, row.GetLow(), b.GetLow())
		require.Equal(t, row.GetClose(), b.GetClose())
		require.Equal(t, row.GetVolume(), b.GetVolume())
		require.Equal(t, row.GetOpenInterest(), b.GetOpenInterest())
	}
}

func TestAvroLoader(t *testing.T) {
	// December 1st, 2022
	now := time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC)

	bars := []Bar{
		NewFakeBar(now),
		NewFakeBar(now.Add(time_series.Day)),
		NewFakeBar(now.Add(2*time_series.Day)),
		NewFakeBar(now.Add(3*time_series.Day)),
		NewFakeBar(now.Add(4*time_series.Day)),
	}

	ctx := context.Background()
	loader := NewAvroLoader()

	buff := bytes.NewBuffer([]byte{})
	err := loader.Write(ctx, buff, bars)
	require.NoError(t, err)

	reader := bytes.NewReader(buff.Bytes())
	output, err := loader.Read(ctx, reader)
	require.NoError(t, err)
	require.Len(t, output, len(bars))
	for index, row := range output {
		b := bars[index]
		require.Equal(t, row.GetTime().String(), b.GetTime().String())
		require.Equal(t, row.GetOpen(), b.GetOpen())
		require.Equal(t, row.GetHigh(), b.GetHigh())
		require.Equal(t, row.GetLow(), b.GetLow())
		require.Equal(t, row.GetClose(), b.GetClose())
		require.Equal(t, row.GetVolume(), b.GetVolume())
		require.Equal(t, row.GetOpenInterest(), b.GetOpenInterest())
	}
}

func TestProtoLoader(t *testing.T) {
	// December 1st, 2022
	now := time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC)

	bars := []Bar{
		NewFakeBar(now),
		NewFakeBar(now.Add(time_series.Day)),
		NewFakeBar(now.Add(2*time_series.Day)),
		NewFakeBar(now.Add(3*time_series.Day)),
		NewFakeBar(now.Add(4*time_series.Day)),
	}

	ctx := context.Background()
	loader := NewProtoLoader()

	buff := bytes.NewBuffer([]byte{})
	err := loader.Write(ctx, buff, bars)
	require.NoError(t, err)

	reader := bytes.NewReader(buff.Bytes())
	output, err := loader.Read(ctx, reader)
	require.NoError(t, err)
	require.Len(t, output, len(bars))
	for index, row := range output {
		b := bars[index]
		require.Equal(t, row.GetTime().String(), b.GetTime().String())
		require.Equal(t, row.GetOpen(), b.GetOpen())
		require.Equal(t, row.GetHigh(), b.GetHigh())
		require.Equal(t, row.GetLow(), b.GetLow())
		require.Equal(t, row.GetClose(), b.GetClose())
		require.Equal(t, row.GetVolume(), b.GetVolume())
		require.Equal(t, row.GetOpenInterest(), b.GetOpenInterest())
	}
}
