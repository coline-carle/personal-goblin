package auction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var sortTests = []struct {
	input    Auctions
	expected Auctions
}{
	{
		Auctions{Auction{Buyout: 1, Quantity: 1}},
		Auctions{Auction{Buyout: 1, Quantity: 1}},
	},
	{
		Auctions{
			Auction{Buyout: 1, Quantity: 1},
			Auction{Buyout: 3, Quantity: 1},
			Auction{Buyout: 2, Quantity: 1},
			Auction{Buyout: 0, Quantity: 1},
		},
		Auctions{
			Auction{Buyout: 0, Quantity: 1},
			Auction{Buyout: 1, Quantity: 1},
			Auction{Buyout: 2, Quantity: 1},
			Auction{Buyout: 3, Quantity: 1},
		},
	},
}

func TestSort(t *testing.T) {
	for _, tt := range sortTests {
		tt.input.sort()
		assert.Equal(t, tt.expected, tt.input)
	}
}
