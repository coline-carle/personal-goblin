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
}

func TestPrice(t *testing.T) {
	for _, tt := range priceTests {
		price, err := BuyoutAverage(tt.input)
		if err != nil {
			t.Fatalf("unexpected error from BuyoutAverage()\n* Error: %s", err)
		}
		if math.Abs(price-tt.expected)/tt.expected > 0.01 {
			t.Fatalf("unexpected Buyout Average \n* expected: %1.2f \n* got: %1.2f\n", tt.expected, price)
		}

	}
}
