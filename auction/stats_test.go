package auction

import (
	"math"
	"testing"
)

var statsTests = []struct {
	input  Auctions
	mean   float64
	stdDev float64
}{
	{
		Auctions{Auction{Buyout: 1, Quantity: 1}, Auction{Buyout: 3, Quantity: 1}, Auction{Buyout: 2, Quantity: 1}, Auction{Buyout: 0, Quantity: 1}},
		1.5,
		1.290,
	},
}

func TestStats(t *testing.T) {
	for _, tt := range statsTests {
		mean := mean(tt.input)
		if math.Abs(mean-tt.mean)/tt.mean > 0.01 {
			t.Fatalf("unexpected mean \n* expected: %1.2f \n* got: %1.2f\n", tt.mean, mean)
		}
		std := stdDev(tt.input, mean)
		if math.Abs(std-tt.stdDev)/tt.stdDev > 0.01 {
			t.Fatalf("unexpected stdDev \n* expected: %1.2f \n* got: %1.2f\n", tt.stdDev, std)
		}
	}
}
