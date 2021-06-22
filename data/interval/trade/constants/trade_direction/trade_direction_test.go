package trade_direction

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
	}
}
