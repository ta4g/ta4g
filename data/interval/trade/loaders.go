package trade

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
)

// Loader reads and writes the Order data to the desired format.
// There are several loaders to choose from, each of which are self-contained with their own schemas:
// 1. CSV
// 2. Avro
// 3. Proto
type Loader interface {
	Read(ctx context.Context, input io.Reader) ([]Order, error)
	Write(ctx context.Context, output io.Writer, input []Order) error
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

func (c csvLoader) Read(ctx context.Context, input io.Reader) ([]Order, error) {
	logger := ctxzap.Extract(ctx)

	// Pull in the CSV
	data, err := ioutil.ReadAll(input)
	if nil != err {
		logger.Error("Failed to read all rows", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Read the rows
	var bars []StandardOrder
	err = csvutil.Unmarshal(data, &bars)
	if nil != err {
		logger.Error("Failed to unmarshal rows", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Type conversion
	output := make([]Order, 0, len(bars))
	for _, b := range bars {
		newOrder, err := b.Clone()
		if nil != err {
			logger.Error("Failed to clone bar", zap.Error(err))
			return nil, err
		}
		output = append(output, newOrder)
	}
	return output, nil
}

func (c csvLoader) Write(ctx context.Context, output io.Writer, input []Order) error {
	logger := ctxzap.Extract(ctx)

	// Type conversion
	bars := make([]StandardOrder, 0, len(input))
	for _, b := range input {
		value, ok := b.(*StandardOrder)
		if !ok {
			value = copyToStandardOrder(b)
		}
		bars = append(bars, *value)
	}

	data, err := csvutil.Marshal(bars)
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

func (j jsonNewLineLoader) Read(ctx context.Context, input io.Reader) ([]Order, error) {
	logger := ctxzap.Extract(ctx)

	// Pull in the CSV
	reader := bufio.NewReader(input)
	output := make([]Order, 0)
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
		bar := &StandardOrder{}
		err = json.Unmarshal(data, bar)
		if nil != err {
			logger.Error("Failed to unmarshal row", zap.Error(err))
			return nil, status.Error(codes.Internal, err.Error())
		}
		output = append(output, bar)
	}
	return output, nil
}

func (j jsonNewLineLoader) Write(ctx context.Context, writer io.Writer, bars []Order) error {
	logger := ctxzap.Extract(ctx)

	for _, bar := range bars {
		// Serialize as json
		stdOrder := copyToStandardOrder(bar)
		data, err := json.Marshal(stdOrder)
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

func (a avroLoader) Read(ctx context.Context, input io.Reader) ([]Order, error) {
	logger := ctxzap.Extract(ctx)

	decoder := avro.NewDecoderForSchema(avroSchema, input)

	output := make([]Order, 0)
	for {
		stdOrder := &StandardOrder{}
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

func (a avroLoader) Write(ctx context.Context, output io.Writer, input []Order) error {
	logger := ctxzap.Extract(ctx)

	encoder := avro.NewEncoderForSchema(avroSchema, output)
	for _, bar := range input {
		stdOrder, ok := bar.(*StandardOrder)
		if !ok {
			stdOrder = copyToStandardOrder(bar)
		}
		err := encoder.Encode(stdOrder)
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

func (a protoLoader) Read(ctx context.Context, input io.Reader) ([]Order, error) {
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

	// Convert the bars
	orders := make([]Order, 0)
	for _, pbOrder := range messages.Orders {
		orderItems := make([]OrderItem, 0, len(pbOrder.Items))
		for _, pbOrderitem := range pbOrder.Items {
			orderItems = append(orderItems, NewOrderItem(
				OrderDirection(pbOrderitem.OrderDirection),
				pbOrderitem.Symbol,
				pbOrderitem.IsOption,
				pbOrderitem.UnitQuantity,
				pbOrderitem.PricePerUnit,
			))
		}

		row := NewStandardOrder(
			pbOrder.GetTime().AsTime(),
			orderItems...,
		)
		orders = append(orders, row)
	}
	return orders, nil
}

func (a protoLoader) Write(ctx context.Context, output io.Writer, input []Order) error {
	logger := ctxzap.Extract(ctx)

	// Type conversion
	pbOrders := make([]*pb.Order, 0, len(input))
	for _, orderItem := range input {
		items := orderItem.GetItems()
		orderItems := make([]*pb.OrderItem, 0, len(items))
		for _, item := range items {
			orderItems = append(orderItems, &pb.OrderItem{
				OrderDirection: int64(item.GetOrderDirection()),
				Symbol:         item.GetSymbol(),
				IsOption:       item.GetIsOption(),
				UnitQuantity:   item.GetUnitQuantity(),
				PricePerUnit:   item.GetPricePerUnit(),
			})
		}

		value := &pb.Order{
			Time:  timestamppb.New(orderItem.GetTime()),
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
