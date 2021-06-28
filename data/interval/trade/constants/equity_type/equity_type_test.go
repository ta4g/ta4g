package equity_type

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
			require.NotEmpty(t, index.String())
		})
	}
}
