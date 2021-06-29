package order_type

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test(t *testing.T) {
	require.Empty(t, min.String(), fmt.Sprintf("%v", int(min)))
	require.Empty(t, max.String(), fmt.Sprintf("%v", int(max)))

	for index := min + 1; index < max; index++ {
		t.Run(fmt.Sprintf("%v: String", int(index)), func(t *testing.T) {
			require.NotEmpty(t, index.String(), fmt.Sprintf("%v", int(index)))
		})

		t.Run(fmt.Sprintf("%v: To/FromProto", int(index)), func(t *testing.T) {
			protoValue := ToProto(index)
			orderTypeValue := FromProto(protoValue)
			require.Equal(t, orderTypeValue, index)
		})
	}

	t.Run("Opposite", func(t *testing.T) {
		require.Equal(t, Opposite(min), min)                               // Same
		require.Equal(t, Opposite(PortfolioOpen), PortfolioClose)          // Switch
		require.Equal(t, Opposite(PortfolioClose), PortfolioOpen)          // Switch
		require.Equal(t, Opposite(PositionAdjustment), PositionAdjustment) // Same
		require.Equal(t, Opposite(PositionOpen), PositionOpen)             // Switch
		require.Equal(t, Opposite(PositionOpen), PositionOpen)             // Switch
		require.Equal(t, Opposite(max), max)                               // Same
	})
}
