package bar

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
	pb "github.com/ta4g/ta4g/gen/interval/bar"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"io/ioutil"
)

// Loader reads and writes the Bar data to the desired format.
// There are several loaders to choose from, each of which are self-contained with their own schemas:
// 1. CSV
// 2. Avro
// 3. Proto
type Loader interface {
	Read(ctx context.Context, input io.Reader) ([]Bar, error)
	Write(ctx context.Context, output io.Writer, input []Bar) error
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

func (c csvLoader) Read(ctx context.Context, input io.Reader) ([]Bar, error) {
	logger := ctxzap.Extract(ctx)

	// Pull in the CSV
	data, err := ioutil.ReadAll(input)
	if nil != err {
		logger.Error("Failed to read all rows", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Read the rows
	var bars []StandardBar
	err = csvutil.Unmarshal(data, &bars)
	if nil != err {
		logger.Error("Failed to unmarshal rows", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Type conversion
	output := make([]Bar, 0, len(bars))
	for _, b := range bars {
		newBar, err := b.Clone()
		if nil != err {
			logger.Error("Failed to clone bar", zap.Error(err))
			return nil, err
		}
		output = append(output, newBar)
	}
	return output, nil
}

func (c csvLoader) Write(ctx context.Context, output io.Writer, input []Bar) error {
	logger := ctxzap.Extract(ctx)

	// Type conversion
	bars := make([]StandardBar, 0, len(input))
	for _, b := range input {
		value, ok := b.(*StandardBar)
		if !ok {
			value = copyToStandardBar(b)
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

func (j jsonNewLineLoader) Read(ctx context.Context, input io.Reader) ([]Bar, error) {
	logger := ctxzap.Extract(ctx)

	// Pull in the CSV
	reader := bufio.NewReader(input)
	output := make([]Bar, 0)
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
		bar := &StandardBar{}
		err = json.Unmarshal(data, bar)
		if nil != err {
			logger.Error("Failed to unmarshal row", zap.Error(err))
			return nil, status.Error(codes.Internal, err.Error())
		}
		output = append(output, bar)
	}
	return output, nil
}

func (j jsonNewLineLoader) Write(ctx context.Context, writer io.Writer, bars []Bar) error {
	logger := ctxzap.Extract(ctx)

	for _, bar := range bars {
		// Serialize as json
		stdBar := copyToStandardBar(bar)
		data, err := json.Marshal(stdBar)
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

func (a avroLoader) Read(ctx context.Context, input io.Reader) ([]Bar, error) {
	logger := ctxzap.Extract(ctx)

	decoder := avro.NewDecoderForSchema(avroSchema, input)

	output := make([]Bar, 0)
	for {
		stdBar := &StandardBar{}
		err := decoder.Decode(stdBar)
		if nil != err && err == io.EOF {
			break
		}
		if nil != err {
			logger.Error("Failed to unmarshal row", zap.Error(err))
			return nil, status.Error(codes.Internal, err.Error())
		}
		output = append(output, stdBar)
	}
	return output, nil
}

func (a avroLoader) Write(ctx context.Context, output io.Writer, input []Bar) error {
	logger := ctxzap.Extract(ctx)

	encoder := avro.NewEncoderForSchema(avroSchema, output)
	for _, bar := range input {
		stdBar, ok := bar.(*StandardBar)
		if !ok {
			stdBar = copyToStandardBar(bar)
		}
		err := encoder.Encode(stdBar)
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

func (a protoLoader) Read(ctx context.Context, input io.Reader) ([]Bar, error) {
	logger := ctxzap.Extract(ctx)

	data, err := ioutil.ReadAll(input)
	if nil != err {
		logger.Error("Failed to read all rows", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	messages := &pb.StandardBars{}
	err = proto.Unmarshal(data, messages)
	if nil != err {
		logger.Error("Failed to unmarshal rows", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Convert the bars
	output := make([]Bar, 0)
	for _, bar := range messages.Bars {
		row := New(
			bar.GetTime().AsTime(),
			bar.GetOpen(),
			bar.GetHigh(),
			bar.GetLow(),
			bar.GetClose(),
			bar.GetVolume(),
			bar.GetOpenInterest(),
		)
		output = append(output, row)
	}
	return output, nil
}

func (a protoLoader) Write(ctx context.Context, output io.Writer, input []Bar) error {
	logger := ctxzap.Extract(ctx)

	// Type conversion
	bars := make([]*pb.StandardBar, 0, len(input))
	for _, b := range input {
		value := &pb.StandardBar{
			Time:         timestamppb.New(b.GetTime()),
			Open:         b.GetOpen(),
			High:         b.GetHigh(),
			Low:          b.GetLow(),
			Close:        b.GetClose(),
			Volume:       b.GetVolume(),
			OpenInterest: b.GetOpenInterest(),
		}
		bars = append(bars, value)
	}

	data, err := proto.Marshal(&pb.StandardBars{Bars: bars})
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
