package time_series

import "time"

// TimeSeries is a wrapper around a ... series of time intervals!
// It represents a current point in a stream of time intervals, all with the same uniform size (eg Day, Hour, Minute, etc).
//
type TimeSeries interface {
	// IntervalSize is the size of each interval, typically Days
	IntervalSize() time.Duration

	// MinValue is the minimum point in the time series
	MinValue() time.Time

	// MaxValue is the maximum point in the time series
	MaxValue() time.Time

	// CurrentValue the current position in the ... time series
	CurrentValue() time.Time

	// Offset will return the `time.Time` relative to the current time.
	//
	// This does not change the current time and is useful for looking backwards in time,
	// or peeking into the future for back tests.
	//
	// Errors:
	// - If you ask for an offset that is out of bounds of the range, an error with GRPC status OutOfRange will be returned
	//
	// Three examples:
	// 1. Offset(1) returns the next time.Time in the series w/o moving CurrentTime
	// 2. Offset(-1) returns the previous time.Time in the series w/o moving CurrentTime
	// 3. Offset(0) returns the current time, this is not especially useful since we have CurrentTime but it is good to know
	Offset(units int) (time.Time, error)

	// Range will return a range of time relative to the current time.
	// This is an INCLUSIVE range, so the starting index and ending index are both part of the output.
	//
	// This does not change the current time and is useful for looking backwards in time,
	// or peeking into the future for back tests.
	//
	// Errors:
	// - If you ask for an offset that is out of bounds of the range, an error with GRPC status OutOfRange will be returned
	//
	// Three examples:
	// 1. Range(1, 2) returns the next 2x time.Time's in the series w/o moving CurrentTime
	// 2. Range(-2, -1) returns the previous 2x time.Time's in the series w/o moving CurrentTime
	// 3. Range(-1, 1) returns the previous, current, and next time.Time's in the series
	//
	Range(start, end int) ([]time.Time, error)

	// Add - Move forward or backwards in time by N units.
	//
	// This assumes the underlying calendar system is moving linearly along in time,
	// and also takes into account any holidays or market closures by skipping them when moving forwards or backwards.
	//
	// Two examples:
	// 1. If we are set to Days, then -15 would move 15 days into the past
	// 1. If we are set to Days, then 15 would move 15 days into the future
	//
	Add(units int) error

	// MoveTo - move forward or backward to a specific time
	// If this is out of range an error will be returned
	//
	MoveTo(value time.Time) error

	// Copy creates a copy of this time series instance at it's current point in time
	Copy() (TimeSeries, error)
}
