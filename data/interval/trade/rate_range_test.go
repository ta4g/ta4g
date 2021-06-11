package trade

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetMarginInterestRate(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  float64
	}{
		{"Negative price", -1.0, 9.50},
		{"Low price", 1.0, 9.50},
		{"Medium price", 50001.0, 8.00},
		{"High price", 1250000.0, 7.50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := GetMarginInterestRate(tt.value)
			require.Equal(t, output, tt.want)
		})
	}
}
