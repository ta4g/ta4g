package time_series

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewInMemoryTimeSeries(t *testing.T) {
	t.Parallel()

	// December 1st, 2022
	now := time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC)

	t.Run("New", func(t *testing.T) {
		type args struct {
			intervalSize time.Duration
			values       []time.Time
			ok           bool
		}
		tests := map[string]args{
			"Invalid IntervalSize":   {0, nil, false},
			"Nil time values":        {time.Second, nil, false},
			"Empty time values":      {time.Second, make([]time.Time, 0), false},
			"Non-sorted time values": {time.Second, []time.Time{now, now.Add(-1 * time.Second)}, false},
			"OK":                     {time.Second, []time.Time{now.Add(-2 * time.Second), now.Add(-1 * time.Second), now}, true},
		}
		for key, arg := range tests {
			t.Run(key, func(t *testing.T) {
				output, err := NewInMemoryTimeSeries(arg.intervalSize, arg.values)
				if !arg.ok {
					require.Error(t, err)
					require.Nil(t, output)
				} else {
					require.NoError(t, err)
					require.NotNil(t, output)
					require.Equal(t, output.IntervalSize(), arg.intervalSize)

					// Cast it so we can inspect the internals
					inMemoryTimeSeries, ok := output.(*InMemoryTimeSeries)
					require.True(t, ok)
					require.NotNil(t, inMemoryTimeSeries)
					require.Equal(t, inMemoryTimeSeries.intervalSize, arg.intervalSize)
					require.Equal(t, inMemoryTimeSeries.values, arg.values)
				}
			})
		}
	})

	t.Run("ReadOnly Operations", func(t *testing.T) {
		values := []time.Time{
			now.Add(-5 * Day),
			now.Add(-4 * Day),
			now.Add(-2 * Day), // Weekend hop
			now.Add(-1 * Day),
			now,
			now.Add(1 * Day),
			now.Add(2 * Day),
		}

		series, err := NewInMemoryTimeSeries(Day, values)
		require.NoError(t, err)
		require.NotNil(t, series)

		err = series.Add(4) // Move to "now"
		require.NoError(t, err)

		t.Run("IntervalSize", func(t *testing.T) {
			require.Equal(t, series.IntervalSize(), Day)
		})

		t.Run("Min Value", func(t *testing.T) {
			require.Equal(t, series.MinValue().String(), now.Add(-5*Day).String())
		})

		t.Run("Max Value", func(t *testing.T) {
			require.Equal(t, series.MaxValue().String(), now.Add(2*Day).String())
		})

		t.Run("Current Value", func(t *testing.T) {
			require.Equal(t, series.CurrentValue().String(), now.String())
		})

		t.Run("Offset", func(t *testing.T) {
			output, err := series.Offset(0)
			require.NoError(t, err)
			require.Equal(t, output.String(), series.CurrentValue().String())

			// Walk forward, staying in-range
			output, err = series.Offset(2)
			require.NoError(t, err)
			require.Equal(t, output.String(), series.CurrentValue().Add(2*Day).String())

			// Walk forward, now go out of range
			output, err = series.Offset(10)
			require.Error(t, err)
			require.Equal(t, output.String(), TimeZero.String())

			// Walk backward, staying in-range
			output, err = series.Offset(-2)
			require.NoError(t, err)
			require.Equal(t, output.String(), series.CurrentValue().Add(-2*Day).String())

			// Walk backward, now go out of range
			output, err = series.Offset(-10)
			require.Error(t, err)
			require.Equal(t, output.String(), TimeZero.String())
		})

		t.Run("Range", func(t *testing.T) {
			// Just today
			output, err := series.Range(0, 0)
			require.NoError(t, err)
			require.Len(t, output, 1)
			require.Equal(t, output[0].String(), series.CurrentValue().String())

			// Walk forward, staying in-range
			output, err = series.Range(0, 1)
			require.NoError(t, err)
			require.Len(t, output, 2)
			require.Equal(t, output[0].String(), series.CurrentValue().String())
			require.Equal(t, output[1].String(), series.CurrentValue().Add(Day).String())

			// Walk forward, now go out of range
			output, err = series.Range(1, 10)
			require.Error(t, err)
			require.Nil(t, output)

			// Walk backward, staying in-range
			output, err = series.Range(-2, -1)
			require.NoError(t, err)
			require.Len(t, output, 2)
			require.Equal(t, output[0].String(), series.CurrentValue().Add(-2*Day).String())
			require.Equal(t, output[1].String(), series.CurrentValue().Add(-1*Day).String())

			// Finally get a range that wraps the current value
			output, err = series.Range(-1, 1)
			require.NoError(t, err)
			require.Len(t, output, 3)
			require.Equal(t, output[0].String(), series.CurrentValue().Add(-1*Day).String())
			require.Equal(t, output[1].String(), series.CurrentValue().String())
			require.Equal(t, output[2].String(), series.CurrentValue().Add(1*Day).String())
		})
	})

	t.Run("Write Operations", func(t *testing.T) {
		values := []time.Time{
			now.Add(-5 * Day),
			now.Add(-4 * Day),
			now.Add(-2 * Day), // Weekend hop
			now.Add(-1 * Day),
			now,
			now.Add(1 * Day),
			now.Add(2 * Day),
		}

		t.Run("Add", func(t *testing.T) {
			series, err := NewInMemoryTimeSeries(
				Day,
				values,
			)
			require.NoError(t, err)
			require.NotNil(t, series)

			err = series.Add(4) // Move to "now"
			require.NoError(t, err)

			// Move forward, staying in-series
			err = series.Add(1)
			require.NoError(t, err)
			require.Equal(t, series.CurrentValue().String(), now.Add(Day).String())

			// Move forward, out of range
			err = series.Add(10)
			require.Error(t, err)
			require.Equal(t, series.CurrentValue().String(), now.Add(Day).String())

			// Move backward, staying in-series
			err = series.Add(-2)
			require.NoError(t, err)
			require.Equal(t, series.CurrentValue().String(), now.Add(-1*Day).String())

			// Move backward, out of range
			err = series.Add(-10)
			require.Error(t, err)
			require.Equal(t, series.CurrentValue().String(), now.Add(-1*Day).String())

			// Move back to starting position
			err = series.Add(1)
			require.NoError(t, err)
			require.Equal(t, series.CurrentValue().String(), now.String())
		})

		t.Run("MoveTo", func(t *testing.T) {
			series, err := NewInMemoryTimeSeries(
				Day,
				values,
			)
			require.NoError(t, err)
			require.NotNil(t, series)

			err = series.Add(4) // Move to "now"
			require.NoError(t, err)

			// Move forward, staying in-series
			err = series.MoveTo(now.Add(2 * Day))
			require.NoError(t, err)
			require.Equal(t, series.CurrentValue().String(), now.Add(2*Day).String())

			// Move forward, out of range
			err = series.MoveTo(now.Add(10 * Day))
			require.Error(t, err)
			require.Equal(t, series.CurrentValue().String(), now.Add(2*Day).String())

			// Move backward, staying in-series
			err = series.MoveTo(now.Add(-2 * Day))
			require.NoError(t, err)
			require.Equal(t, series.CurrentValue().String(), now.Add(-2*Day).String())

			// Move backward, out of range
			err = series.MoveTo(now.Add(-10 * Day))
			require.Error(t, err)
			require.Equal(t, series.CurrentValue().String(), now.Add(-2*Day).String())

			// Move back to starting position
			err = series.MoveTo(now)
			require.NoError(t, err)
			require.Equal(t, series.CurrentValue().String(), now.String())
		})
	})

	t.Run("Copy", func(t *testing.T) {
		values := []time.Time{
			now.Add(-5 * Day),
			now.Add(-4 * Day),
			now.Add(-2 * Day), // Weekend hop
			now.Add(-1 * Day),
			now,
			now.Add(1 * Day),
			now.Add(2 * Day),
		}

		series, err := NewInMemoryTimeSeries(Day, values)
		require.NoError(t, err)
		require.NotNil(t, series)

		err = series.Add(4) // Move to "now"
		require.NoError(t, err)

		// Clone the series
		output, err := series.Copy()
		require.NoError(t, err)
		require.NotNil(t, output)
		require.Equal(t, output.IntervalSize(), series.IntervalSize())
		require.Equal(t, output.MinValue().String(), series.MinValue().String())
		require.Equal(t, output.MaxValue().String(), series.MaxValue().String())
		require.Equal(t, output.CurrentValue().String(), series.CurrentValue().String())

		// Move the original series, the clone should not be affected
		err = series.Add(1)
		require.NoError(t, err)
		require.Equal(t, series.CurrentValue().String(), now.Add(Day).String())
		require.Equal(t, output.CurrentValue().String(), now.String())

		// Now move the clone and verify again, the original should not be affected
		err = output.Add(-2)
		require.NoError(t, err)
		require.Equal(t, series.CurrentValue().String(), now.Add(Day).String())
		require.Equal(t, output.CurrentValue().String(), now.Add(-2*Day).String())
	})
}
