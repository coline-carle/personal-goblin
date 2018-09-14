package auction

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var statsTests = []struct {
	input  Auctions
	mean   uint64
	stdDev float64
}{
	{
		Auctions{Auction{Buyout: 100, Quantity: 1}, Auction{Buyout: 300, Quantity: 1}, Auction{Buyout: 200, Quantity: 1}, Auction{Buyout: 000, Quantity: 1}},
		150,
		129.10,
	},
	{
		Auctions{Auction{Buyout: 100, Quantity: 1}},
		100,
		0.0,
	},
}

func TestStats(t *testing.T) {
	for _, tt := range statsTests {
		mean := tt.input.mean()
		assert.Equal(t, tt.mean, mean)
		std := tt.input.stdDev(mean)
		if math.IsNaN(std) {
			t.Fatalf("std: NaN")
		}
		if math.Abs(std-tt.stdDev)/tt.stdDev > 0.01 {
			t.Fatalf("unexpected stdDev \n* expected: %1.2f \n* got: %1.2f\n", tt.stdDev, std)
		}
	}
}
