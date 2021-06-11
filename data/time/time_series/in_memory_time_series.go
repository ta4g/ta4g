package time_series

import (
	"time"
)

// Compile time type assertion
var _ TimeSeries = &InMemoryTimeSeries{}

// InMemoryTimeSeries is an in-memory time series iterator
//
// NOTE: This is not intended to be a thread-safe iterator.
//       If you want to perform multiple operations in parallel,
//       then use the `Copy` method to create a new TimeSeries for your go-routine.
//
type InMemoryTimeSeries struct {
	intervalSize    time.Duration
	values          []time.Time
	currentPosition int
}

// NewInMemoryTimeSeries creates a new TimeSeries instance.
// The input is validated to make sure all of these conditions are true:
//
// 1. The intervalSize is greater than zero.
// 2. The time values array is not empty.
// 3. The time values array is sorted in ascending order.
//
func NewInMemoryTimeSeries(intervalSize time.Duration, values []time.Time) (TimeSeries, error) {
	if intervalSize <= time.Duration(0) {
		return nil, InvalidArgument
	}

	// The input array must not be empty
	if len(values) == 0 {
		return nil, InvalidArgument
	}

	// The inputs must be in ascending order
	for index, value := range values {
		if index == 0 {
			continue
		}
		if values[index-1].After(value) {
			return nil, InvalidArgument
		}
	}

	output := &InMemoryTimeSeries{
		intervalSize:    intervalSize,
		values:          values,
		currentPosition: 0,
	}
	return output, nil
}

func (i *InMemoryTimeSeries) IntervalSize() time.Duration {
	return i.intervalSize
}

func (i *InMemoryTimeSeries) MinValue() time.Time {
	return i.values[0]
}

func (i *InMemoryTimeSeries) MaxValue() time.Time {
	return i.values[len(i.values)-1]
}

func (i *InMemoryTimeSeries) CurrentValue() time.Time {
	return i.values[i.currentPosition]
}

func (i *InMemoryTimeSeries) inRange(offset int) (int, bool) {
	index := i.currentPosition + offset
	return index, index >= 0 && index < len(i.values)
}

func (i *InMemoryTimeSeries) Offset(units int) (time.Time, error) {
	index, ok := i.inRange(units)
	if !ok {
		return TimeZero, OutOfRange
	}
	return i.values[index], nil
}

func (i *InMemoryTimeSeries) Range(start, end int) ([]time.Time, error) {
	if start > end {
		return nil, InvalidArgument
	}
	lowerIndex, ok := i.inRange(start)
	if !ok {
		return nil, OutOfRange
	}
	upperIndex, ok := i.inRange(end + 1)
	if !ok {
		return nil, OutOfRange
	}
	return i.values[lowerIndex:upperIndex], nil
}

func (i *InMemoryTimeSeries) Add(units int) error {
	index, ok := i.inRange(units)
	if !ok {
		return OutOfRange
	}
	i.currentPosition = index
	return nil
}

func (i *InMemoryTimeSeries) MoveTo(input time.Time) error {
	for index, value := range i.values {
		if value.Equal(input) {
			i.currentPosition = index
			return nil
		}
	}
	return InvalidArgument
}

func (i *InMemoryTimeSeries) Copy() (TimeSeries, error) {
	return &InMemoryTimeSeries{
		intervalSize:    i.intervalSize,
		values:          i.values,
		currentPosition: i.currentPosition,
	}, nil
}
