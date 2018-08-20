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
		Auctions{Auction{buyoutUnit: 1}},
		Auctions{Auction{buyoutUnit: 1}},
	},
	{
		Auctions{Auction{buyoutUnit: 1}, Auction{buyoutUnit: 3}, Auction{buyoutUnit: 2}, Auction{buyoutUnit: 0}},
		Auctions{Auction{buyoutUnit: 0}, Auction{buyoutUnit: 1}, Auction{buyoutUnit: 2}, Auction{buyoutUnit: 3}},
	},
}

func TestSort(t *testing.T) {
	for _, tt := range sortTests {
		tt.input.sort()
		assert.Equal(t, tt.expected, tt.input)
	}
}
