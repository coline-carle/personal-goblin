package auction

import (
	"math"
	"testing"
)

var priceTests = []struct {
	input    Auctions
	expected float64
}{

	{
		[]Auction{
			Auction{Quantity: 1, Buyout: 5},
			Auction{Quantity: 1, Buyout: 13},
			Auction{Quantity: 1, Buyout: 13},
			Auction{Quantity: 1, Buyout: 15},
			Auction{Quantity: 1, Buyout: 15},
			Auction{Quantity: 1, Buyout: 15},
			Auction{Quantity: 1, Buyout: 16},
			Auction{Quantity: 1, Buyout: 17},
			Auction{Quantity: 1, Buyout: 17},
			Auction{Quantity: 1, Buyout: 19},
			Auction{Quantity: 1, Buyout: 20},
			Auction{Quantity: 1, Buyout: 20},
			Auction{Quantity: 1, Buyout: 20},
			Auction{Quantity: 1, Buyout: 20},
			Auction{Quantity: 1, Buyout: 20},
			Auction{Quantity: 1, Buyout: 20},
			Auction{Quantity: 1, Buyout: 21},
			Auction{Quantity: 1, Buyout: 21},
			Auction{Quantity: 1, Buyout: 29},
			Auction{Quantity: 1, Buyout: 45},
			Auction{Quantity: 1, Buyout: 45},
			Auction{Quantity: 1, Buyout: 46},
			Auction{Quantity: 1, Buyout: 47},
			Auction{Quantity: 1, Buyout: 100},
		},
		14.5,
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
		33602.47,
	},
}

func TestPrice(t *testing.T) {
	for _, tt := range priceTests {
		price, err := BuyoutAverage(tt.input)
		if err != nil {
			t.Fatalf("unexpected error from BuyoutAverage()\n* Error: %s\n", err)
		}
		if math.IsNaN(price) {
			t.Fatalf("unexpected NaN value in price: %f\n", tt.expected)
		}
		if math.Abs(price-tt.expected)/tt.expected > 0.01 {
			t.Fatalf("unexpected Buyout Average \n* expected: %1.2f \n* got: %1.2f\n", tt.expected, price)
		}

	}
}
