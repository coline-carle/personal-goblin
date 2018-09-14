package auction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var priceTests = []struct {
	input    Auctions
	expected uint64
}{

	{
		[]Auction{
			Auction{Quantity: 1, Buyout: 50},
			Auction{Quantity: 1, Buyout: 130},
			Auction{Quantity: 1, Buyout: 130},
			Auction{Quantity: 1, Buyout: 150},
			Auction{Quantity: 1, Buyout: 150},
			Auction{Quantity: 1, Buyout: 150},
			Auction{Quantity: 1, Buyout: 160},
			Auction{Quantity: 1, Buyout: 170},
			Auction{Quantity: 1, Buyout: 170},
			Auction{Quantity: 1, Buyout: 190},
			Auction{Quantity: 1, Buyout: 200},
			Auction{Quantity: 1, Buyout: 200},
			Auction{Quantity: 1, Buyout: 200},
			Auction{Quantity: 1, Buyout: 200},
			Auction{Quantity: 1, Buyout: 200},
			Auction{Quantity: 1, Buyout: 200},
			Auction{Quantity: 1, Buyout: 210},
			Auction{Quantity: 1, Buyout: 210},
			Auction{Quantity: 1, Buyout: 290},
			Auction{Quantity: 1, Buyout: 450},
			Auction{Quantity: 1, Buyout: 450},
			Auction{Quantity: 1, Buyout: 460},
			Auction{Quantity: 1, Buyout: 470},
			Auction{Quantity: 1, Buyout: 1000},
		},
		145,
	},
	{
		[]Auction{
			Auction{Quantity: 1, Buyout: 5},
		},
		5,
	},
	{
		[]Auction{
			Auction{Quantity: 1, Buyout: 300000000},
		},
		300000000,
	},
	{
		[]Auction{
			Auction{Quantity: 2, Buyout: 67600},
			Auction{Quantity: 20, Buyout: 672000},
			Auction{Quantity: 20, Buyout: 672000},
			Auction{Quantity: 20, Buyout: 672000},
			Auction{Quantity: 20, Buyout: 672000},
			Auction{Quantity: 20, Buyout: 672000},
			Auction{Quantity: 20, Buyout: 672000},
			Auction{Quantity: 20, Buyout: 672000},
			Auction{Quantity: 20, Buyout: 672000},
			Auction{Quantity: 20, Buyout: 750000},
			Auction{Quantity: 20, Buyout: 750000},
			Auction{Quantity: 20, Buyout: 750000},
			Auction{Quantity: 31, Buyout: 1054000},
			Auction{Quantity: 16, Buyout: 1160224},
			Auction{Quantity: 27, Buyout: 1328508},
			Auction{Quantity: 29, Buyout: 1371621},
			Auction{Quantity: 20, Buyout: 1450280},
			Auction{Quantity: 101, Buyout: 4888400},
			Auction{Quantity: 153, Buyout: 7650000},
			Auction{Quantity: 200, Buyout: 9840800},
			Auction{Quantity: 200, Buyout: 9840800},
			Auction{Quantity: 200, Buyout: 78000000},
			Auction{Quantity: 200, Buyout: 78000000},
			Auction{Quantity: 200, Buyout: 78000000},
			Auction{Quantity: 200, Buyout: 78000000},
			Auction{Quantity: 200, Buyout: 78000000},
			Auction{Quantity: 200, Buyout: 78000000},
			Auction{Quantity: 200, Buyout: 78000000},
			Auction{Quantity: 200, Buyout: 78000000},
			Auction{Quantity: 200, Buyout: 78000000},
		},
		33600,
	},
}

func TestPrice(t *testing.T) {
	for _, tt := range priceTests {
		market, err := BuyoutAverage(tt.input)
		if err != nil {
			t.Fatalf("unexpected error from BuyoutAverage()\n* Error: %s\n", err)
		}
		assert.EqualValuesf(t, tt.expected, market.AvgBuyout, "expected %d, got %d\n", tt.expected, market.AvgBuyout)
	}
}
