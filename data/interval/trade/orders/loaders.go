package orders

import (
	"bufio"
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/hamba/avro"
	"github.com/jszwec/csvutil"
	pb "github.com/ta4g/ta4g/gen/interval/trade"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"io/ioutil"
	"time"
)

// Loader reads and writes the Order data to the desired format.
// There are several loaders to choose from, each of which are self-contained with their own schemas:
// 1. CSV
// 2. Avro
// 3. Proto
type Loader interface {
	Read(ctx context.Context, input io.Reader) ([]*Order, error)
	Write(ctx context.Context, output io.Writer, input []*Order) error
}

// Compile time type assertions
var _ Loader = &csvLoader{}
var _ Loader = &jsonNewLineLoader{}
var _ Loader = &avroLoader{}
var _ Loader = &protoLoader{}

type csvLoader struct{}
type jsonNewLineLoader struct{}
type avroLoader struct{}
type protoLoader struct{}

//go:embed schema.avro
var schemaStr string
var avroSchema avro.Schema

func init() {
	schema, err := avro.Parse(schemaStr)
	if nil != err {
		panic(err)
	} else {
		avroSchema = schema
	}
}

//
// CSV Loader
//

func NewCSVLoader() Loader {
	return &csvLoader{}
}

func (c csvLoader) Read(ctx context.Context, input io.Reader) ([]*Order, error) {
	logger := ctxzap.Extract(ctx)

	// Pull in the CSV
	data, err := ioutil.ReadAll(input)
	if nil != err {
		logger.Error("Failed to read all rows", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Read the rows
	var bars []Order
	err = csvutil.Unmarshal(data, &bars)
	if nil != err {
		logger.Error("Failed to unmarshal rows", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Type conversion
	output := make([]*Order, 0, len(bars))
	for _, item := range bars {
		output = append(output, item.Clone())
	}
	return output, nil
}

func (c csvLoader) Write(ctx context.Context, output io.Writer, input []*Order) error {
	logger := ctxzap.Extract(ctx)

	// Type conversion
	items := make([]Order, 0, len(input))
	for _, value := range input {
		items = append(items, *value)
	}

	data, err := csvutil.Marshal(items)
	if nil != err {
		logger.Error("Failed to marshal rows", zap.Error(err))
		return status.Error(codes.Internal, err.Error())
	}

	_, err = io.Copy(output, bytes.NewReader(data))
	if nil != err {
		logger.Error("Failed to write all rows", zap.Error(err))
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

//
// JSON New Line Loader
//

func NewJsonNewLineLoader() Loader {
	return &jsonNewLineLoader{}
}

func (j jsonNewLineLoader) Read(ctx context.Context, input io.Reader) ([]*Order, error) {
	logger := ctxzap.Extract(ctx)

	// Pull in the CSV
	reader := bufio.NewReader(input)
	output := make([]*Order, 0)
	for {
		// Read the rows line by line
		data, err := reader.ReadBytes('\n')
		if nil != err && err == io.EOF {
			break
		}
		if nil != err {
			logger.Error("Failed to read line", zap.Error(err))
			return nil, status.Error(codes.Internal, err.Error())
		}
		if len(data) == 0 {
			break
		}

		// Now parse the JSON and add it to the output
		item := &Order{}
		err = json.Unmarshal(data, item)
		if nil != err {
			logger.Error("Failed to unmarshal row", zap.Error(err))
			return nil, status.Error(codes.Internal, err.Error())
		}
		output = append(output, item)
	}

	return output, nil
}

func (j jsonNewLineLoader) Write(ctx context.Context, writer io.Writer, input []*Order) error {
	logger := ctxzap.Extract(ctx)

	for _, item := range input {
		// Serialize as json
		data, err := json.Marshal(item)
		if nil != err {
			logger.Error("Failed to marshal row", zap.Error(err))
			return status.Error(codes.Internal, err.Error())
		}

		// Write the bar
		_, err = writer.Write(data)
		if nil != err {
			logger.Error("Failed to write line", zap.Error(err))
			return status.Error(codes.Internal, err.Error())
		}

		// Now write the delimiter
		_, err = writer.Write([]byte("\n"))
		if nil != err {
			logger.Error("Failed to write line", zap.Error(err))
			return status.Error(codes.Internal, err.Error())
		}
	}
	return nil
}

//
// Avro Loader
//

func NewAvroLoader() Loader {
	return &avroLoader{}
}

func (a avroLoader) Read(ctx context.Context, input io.Reader) ([]*Order, error) {
	logger := ctxzap.Extract(ctx)

	decoder := avro.NewDecoderForSchema(avroSchema, input)

	output := make([]*Order, 0)
	for {
		stdOrder := &Order{}
		err := decoder.Decode(stdOrder)
		if nil != err && err == io.EOF {
			break
		}
		if nil != err {
			logger.Error("Failed to unmarshal row", zap.Error(err))
			return nil, status.Error(codes.Internal, err.Error())
		}
		output = append(output, stdOrder)
	}
	return output, nil
}

func (a avroLoader) Write(ctx context.Context, output io.Writer, input []*Order) error {
	logger := ctxzap.Extract(ctx)

	encoder := avro.NewEncoderForSchema(avroSchema, output)
	for _, item := range input {
		err := encoder.Encode(item)
		if nil != err {
			logger.Error("Failed to marshal row", zap.Error(err))
			return status.Error(codes.Internal, err.Error())
		}
	}
	return nil
}

//
// Proto Loader
//

func NewProtoLoader() Loader {
	return &protoLoader{}
}

func (a protoLoader) Read(ctx context.Context, input io.Reader) ([]*Order, error) {
	logger := ctxzap.Extract(ctx)

	data, err := ioutil.ReadAll(input)
	if nil != err {
		logger.Error("Failed to read all rows", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	messages := &pb.Orders{}
	err = proto.Unmarshal(data, messages)
	if nil != err {
		logger.Error("Failed to unmarshal rows", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Convert the rows
	output := make([]*Order, 0)
	for _, pbOrder := range messages.Orders {
		// Convert them to OrderItems
		orderItems := make([]*OrderItem, 0, len(pbOrder.Items))
		for _, item := range pbOrder.Items {
			orderItems = append(
				orderItems,
				&OrderItem{
					Direction:         Direction(item.Direction),
					ItemType:          ItemType(item.ItemType),
					Symbol:            item.Symbol,
					Amount:            item.Amount,
					QuantityPerAmount: item.QuantityPerAmount,
					Price:             item.Price,
				},
			)
		}
		// Now make an order
		row := NewOrder(
			pbOrder.GetTime().AsTime(),
			orderItems...,
		)
		output = append(output, row)
	}
	return output, nil
}

func (a protoLoader) Write(ctx context.Context, output io.Writer, input []*Order) error {
	logger := ctxzap.Extract(ctx)

	// Type conversion
	pbOrders := make([]*pb.Order, 0, len(input))
	for _, orderItem := range input {
		items := orderItem.OrderItems
		orderItems := make([]*pb.OrderItem, 0, len(items))
		for _, item := range items {
			orderItems = append(orderItems, &pb.OrderItem{
				Direction:         int64(item.Direction),
				ItemType:          int64(item.ItemType),
				Symbol:            item.Symbol,
				Amount:            item.Amount,
				QuantityPerAmount: item.QuantityPerAmount,
				Price:             item.Price,
			})
		}

		value := &pb.Order{
			Time:  timestamppb.New(time.Unix(orderItem.UnixTime, 0)),
			Items: orderItems,
		}
		pbOrders = append(pbOrders, value)
	}

	data, err := proto.Marshal(&pb.Orders{Orders: pbOrders})
	if nil != err {
		logger.Error("Failed to marshal rows", zap.Error(err))
		return status.Error(codes.Internal, err.Error())
	}

	_, err = io.Copy(output, bytes.NewReader(data))
	if nil != err {
		logger.Error("Failed to write all rows", zap.Error(err))
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
