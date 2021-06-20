package constants

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestItemType(t *testing.T) {
	for index := minEquityType + 1 ; index < maxEquityType; index++ {
		t.Run(fmt.Sprintf("%v: String", int(index)), func(t *testing.T) {
			require.NotEmpty(t, index.String())
		})
	}
}
