package trade

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

	buyCoveredCallOrder := NewStandardOrder(
		now,
		NewOrderItem(BuyOrderDirection, "ABC", false, 100, 10.01),
		NewOrderItem(SellOrderDirection, "ABC CALL @ 10.0", true, 1, 1.01*100),
	)
	sellCoveredCallOrder := NewStandardOrder(
		now.Add(10*time_series.Day),
		NewOrderItem(SellOrderDirection, "ABC", false, 100, 10.01),
		NewOrderItem(BuyOrderDirection, "ABC CALL @ 10.0", true, 1, 1.01*100),
	)
	orders := []Order{buyCoveredCallOrder, sellCoveredCallOrder}

	ctx := context.Background()
	loader := NewCSVLoader()

	buff := bytes.NewBuffer([]byte{})
	err := loader.Write(ctx, buff, orders)
	require.Error(t, err)

	// NOTE: CSV serialization of nested structures is not supported yet, so this is marked as a TODO
	//
	//lines := strings.Split(buff.String(), "\n")
	//require.Len(t, lines, len(orders)+2)
	//require.Empty(t, lines[len(orders)+1]) // Last line is blank
	//
	//reader := bytes.NewReader(buff.Bytes())
	//output, err := loader.Read(ctx, reader)
	//require.NoError(t, err)
	//require.Len(t, output, len(orders))
	//for index, row := range output {
	//	b := orders[index]
	//	require.Equal(t, row.GetTime().String(), b.GetTime().String())
	//}
}

func TestJsonNewLineLoader(t *testing.T) {
	// December 1st, 2022
	now := time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC)

	buyCoveredCallOrder := NewStandardOrder(
		now,
		NewOrderItem(BuyOrderDirection, "ABC", false, 100, 10.01),
		NewOrderItem(SellOrderDirection, "ABC CALL @ 10.0", true, 1, 1.01*100),
	)
	sellCoveredCallOrder := NewStandardOrder(
		now.Add(10*time_series.Day),
		NewOrderItem(SellOrderDirection, "ABC", false, 100, 10.01),
		NewOrderItem(BuyOrderDirection, "ABC CALL @ 10.0", true, 1, 1.01*100),
	)
	orders := []Order{buyCoveredCallOrder, sellCoveredCallOrder}


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
		require.Equal(t, row.GetTime().String(), b.GetTime().String())
		require.Len(t, row.GetItems(), len(b.GetItems()))
		for itemIndex, orderItem := range row.GetItems() {
			bOrderItem := b.GetItems()[itemIndex]
			require.Equal(t, orderItem.GetOrderDirection(), bOrderItem.GetOrderDirection())
			require.Equal(t, orderItem.GetSymbol(), bOrderItem.GetSymbol())
			require.Equal(t, orderItem.GetIsOption(), bOrderItem.GetIsOption())
			require.Equal(t, orderItem.GetUnitQuantity(), bOrderItem.GetUnitQuantity())
			require.Equal(t, orderItem.GetPricePerUnit(), bOrderItem.GetPricePerUnit())
		}
	}
}

func TestAvroLoader(t *testing.T) {
	// December 1st, 2022
	now := time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC)

	buyCoveredCallOrder := NewStandardOrder(
		now,
		NewOrderItem(BuyOrderDirection, "ABC", false, 100, 10.01),
		NewOrderItem(SellOrderDirection, "ABC CALL @ 10.0", true, 1, 1.01*100),
	)
	sellCoveredCallOrder := NewStandardOrder(
		now.Add(10*time_series.Day),
		NewOrderItem(SellOrderDirection, "ABC", false, 100, 10.01),
		NewOrderItem(BuyOrderDirection, "ABC CALL @ 10.0", true, 1, 1.01*100),
	)
	orders := []Order{buyCoveredCallOrder, sellCoveredCallOrder}


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
		require.Equal(t, row.GetTime().String(), b.GetTime().String())
		require.Len(t, row.GetItems(), len(b.GetItems()))
		for itemIndex, orderItem := range row.GetItems() {
			bOrderItem := b.GetItems()[itemIndex]
			require.Equal(t, orderItem.GetOrderDirection(), bOrderItem.GetOrderDirection())
			require.Equal(t, orderItem.GetSymbol(), bOrderItem.GetSymbol())
			require.Equal(t, orderItem.GetIsOption(), bOrderItem.GetIsOption())
			require.Equal(t, orderItem.GetUnitQuantity(), bOrderItem.GetUnitQuantity())
			require.Equal(t, orderItem.GetPricePerUnit(), bOrderItem.GetPricePerUnit())
		}
	}
}

func TestProtoLoader(t *testing.T) {
	// December 1st, 2022
	now := time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC)

	buyCoveredCallOrder := NewStandardOrder(
		now,
		NewOrderItem(BuyOrderDirection, "ABC", false, 100, 10.01),
		NewOrderItem(SellOrderDirection, "ABC CALL @ 10.0", true, 1, 1.01*100),
	)
	sellCoveredCallOrder := NewStandardOrder(
		now.Add(10*time_series.Day),
		NewOrderItem(SellOrderDirection, "ABC", false, 100, 10.01),
		NewOrderItem(BuyOrderDirection, "ABC CALL @ 10.0", true, 1, 1.01*100),
	)
	orders := []Order{buyCoveredCallOrder, sellCoveredCallOrder}


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
		require.Equal(t, row.GetTime().String(), b.GetTime().String())
		require.Len(t, row.GetItems(), len(b.GetItems()))
		for itemIndex, orderItem := range row.GetItems() {
			bOrderItem := b.GetItems()[itemIndex]
			require.Equal(t, orderItem.GetOrderDirection(), bOrderItem.GetOrderDirection())
			require.Equal(t, orderItem.GetSymbol(), bOrderItem.GetSymbol())
			require.Equal(t, orderItem.GetIsOption(), bOrderItem.GetIsOption())
			require.Equal(t, orderItem.GetUnitQuantity(), bOrderItem.GetUnitQuantity())
			require.Equal(t, orderItem.GetPricePerUnit(), bOrderItem.GetPricePerUnit())
		}
	}
}
