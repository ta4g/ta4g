package trade_direction

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test(t *testing.T) {
	require.Equal(t, min.String(), "UNKNOWN")
	require.Equal(t, max.String(), fmt.Sprintf("%v", int(max)))

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
		require.Equal(t, Opposite(min), min)         // Same
		require.Equal(t, Opposite(Buy), Sell)        // Switch
		require.Equal(t, Opposite(Neutral), Neutral) // Same
		require.Equal(t, Opposite(Sell), Buy)        // Switch
		require.Equal(t, Opposite(max), max)         // Same
	})
}
