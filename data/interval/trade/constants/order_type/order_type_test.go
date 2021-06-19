package order_type

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOrderType(t *testing.T) {
	require.Empty(t, minOrderType.String(), fmt.Sprintf("%v", 0))
	require.Empty(t, maxOrderType.String(), fmt.Sprintf("%v", int(maxOrderType)))

	for index := minOrderType + 1; index < maxOrderType; index++ {
		t.Run(fmt.Sprintf("%v: String", int(index)), func(t *testing.T) {
			require.NotEmpty(t, index.String(), fmt.Sprintf("%v", int(index)))
		})
	}
}
