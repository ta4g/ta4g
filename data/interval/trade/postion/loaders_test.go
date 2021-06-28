package postion

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/require"
	"github.com/ta4g/ta4g/data/interval/trade/constants/trade_direction"
	"github.com/ta4g/ta4g/data/interval/trade/trade_record"
	"github.com/ta4g/ta4g/data/time/time_series"
	"strings"
	"testing"
	"time"
)

func TestCSVLoader(t *testing.T) {
	// December 1st, 2022
	now := time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC)

	buyCoveredCallOrder := NewTrade(
		now,
		trade_record.NewStockOrderItem(trade_direction.Buy, "ABC", 100, 10.01),
		trade_record.NewOptionOrderItem(trade_direction.Sell, "ABC CALL @ 10.0", 1, 1.01*100),
	)
	sellCoveredCallOrder := NewTrade(
		now.Add(10*time_series.Day),
		trade_record.NewStockOrderItem(trade_direction.Sell, "ABC", 100, 10.01),
		trade_record.NewOptionOrderItem(trade_direction.Buy, "ABC CALL @ 10.0", 1, 1.01*100),
	)
	orders := []*Trade{buyCoveredCallOrder, sellCoveredCallOrder}

	ctx := context.Background()
	loader := NewCSVLoader()

	buff := bytes.NewBuffer([]byte{})
	err := loader.Write(ctx, buff, orders)
	require.Error(t, err)

	// NOTE: CSV serialization of nested structures is not supported yet, so this is marked as a TODO
	//
	//lines := strings.Split(buff.String(), "\n")
	//require.Len(t, lines, len(postion)+2)
	//require.Empty(t, lines[len(postion)+1]) // Last line is blank
	//
	//reader := bytes.NewReader(buff.Bytes())
	//output, err := loader.Read(ctx, reader)
	//require.NoError(t, err)
	//require.Len(t, output, len(postion))
	//for index, row := range output {
	//	b := postion[index]
	//	require.Equal(t, row.UnixTime, b.UnixTime)
	//}
}

func TestJsonNewLineLoader(t *testing.T) {
	// December 1st, 2022
	now := time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC)

	buyCoveredCallOrder := NewTrade(
		now,
		trade_record.NewStockOrderItem(trade_direction.Buy, "ABC", 100, 10.01),
		trade_record.NewOptionOrderItem(trade_direction.Sell, "ABC CALL @ 10.0", 1, 1.01*100),
	)
	sellCoveredCallOrder := NewTrade(
		now.Add(10*time_series.Day),
		trade_record.NewStockOrderItem(trade_direction.Sell, "ABC", 100, 10.01),
		trade_record.NewOptionOrderItem(trade_direction.Buy, "ABC CALL @ 10.0", 1, 1.01*100),
	)
	orders := []*Trade{buyCoveredCallOrder, sellCoveredCallOrder}

	ctx := context.Background()
	loader := NewJsonNewLineLoader()

	buff := bytes.NewBuffer([]byte{})
	err := loader.Write(ctx, buff, orders)
	require.NoError(t, err)

	lines := strings.Split(buff.String(), "\n")
	require.Len(t, lines, len(orders)+1)
	require.Empty(t, lines[len(orders)]) // Last line is blank

	reader := bytes.NewReader(buff.Bytes())
	output, err := loader.Read(ctx, reader)
	require.NoError(t, err)
	require.Len(t, output, len(orders))
	for index, row := range output {
		b := orders[index]
		require.Equal(t, row.UnixTime, b.UnixTime)
		require.Len(t, row.OrderItems, len(b.OrderItems))
		for itemIndex, orderItem := range row.OrderItems {
			bOrderItem := b.OrderItems[itemIndex]
			require.Equal(t, orderItem, bOrderItem)
		}
	}
}

func TestAvroLoader(t *testing.T) {
	// December 1st, 2022
	now := time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC)

	buyCoveredCallOrder := NewTrade(
		now,
		trade_record.NewStockOrderItem(trade_direction.Buy, "ABC", 100, 10.01),
		trade_record.NewOptionOrderItem(trade_direction.Sell, "ABC CALL @ 10.0", 1, 1.01*100),
	)
	sellCoveredCallOrder := NewTrade(
		now.Add(10*time_series.Day),
		trade_record.NewStockOrderItem(trade_direction.Sell, "ABC", 100, 10.01),
		trade_record.NewOptionOrderItem(trade_direction.Buy, "ABC CALL @ 10.0", 1, 1.01*100),
	)
	orders := []*Trade{buyCoveredCallOrder, sellCoveredCallOrder}

	ctx := context.Background()
	loader := NewAvroLoader()

	buff := bytes.NewBuffer([]byte{})
	err := loader.Write(ctx, buff, orders)
	require.NoError(t, err)

	reader := bytes.NewReader(buff.Bytes())
	output, err := loader.Read(ctx, reader)
	require.NoError(t, err)
	require.Len(t, output, len(orders))
	for index, row := range output {
		b := orders[index]
		require.Equal(t, row.UnixTime, b.UnixTime)
		require.Len(t, row.OrderItems, len(b.OrderItems))
		for itemIndex, orderItem := range row.OrderItems {
			bOrderItem := b.OrderItems[itemIndex]
			require.Equal(t, orderItem, bOrderItem)
		}
	}
}

func TestProtoLoader(t *testing.T) {
	// December 1st, 2022
	now := time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC)

	buyCoveredCallOrder := NewTrade(
		now,
		trade_record.NewStockOrderItem(trade_direction.Buy, "ABC", 100, 10.01),
		trade_record.NewOptionOrderItem(trade_direction.Sell, "ABC CALL @ 10.0", 1, 1.01*100),
	)
	sellCoveredCallOrder := NewTrade(
		now.Add(10*time_series.Day),
		trade_record.NewStockOrderItem(trade_direction.Sell, "ABC", 100, 10.01),
		trade_record.NewOptionOrderItem(trade_direction.Buy, "ABC CALL @ 10.0", 1, 1.01*100),
	)
	orders := []*Trade{buyCoveredCallOrder, sellCoveredCallOrder}

	ctx := context.Background()
	loader := NewProtoLoader()

	buff := bytes.NewBuffer([]byte{})
	err := loader.Write(ctx, buff, orders)
	require.NoError(t, err)

	reader := bytes.NewReader(buff.Bytes())
	output, err := loader.Read(ctx, reader)
	require.NoError(t, err)
	require.Len(t, output, len(orders))
	for index, row := range output {
		b := orders[index]
		require.Equal(t, row.UnixTime, b.UnixTime)
		require.Len(t, row.OrderItems, len(b.OrderItems))
		for itemIndex, orderItem := range row.OrderItems {
			bOrderItem := b.OrderItems[itemIndex]
			require.Equal(t, orderItem, bOrderItem)
		}
	}
}
