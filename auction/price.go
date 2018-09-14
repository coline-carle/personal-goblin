package auction

import (
	"errors"
)

var (
	// ErrNoAuctions error returned when the Auctions slice is empty
	ErrNoAuctions = errors.New("error: no auctions in slice")
)

const (
	throwFactor = 1.5
)

// Market represent the state of the market of a given item
type Market struct {
	AvgBuyout uint64
	MinBuyout uint64
	Size      uint64
}

func firstCut(auctions Auctions) Auctions {
	maxIndex := int(float64(len(auctions)) * 0.30)
	minIndex := int(float64(len(auctions)) * 0.15)
	prevPrice := float64(auctions[0].BuyoutUnit())
	var i int
	for i = 0; i < maxIndex; i++ {
		price := float64(auctions[i].BuyoutUnit())
		if i > minIndex && price > 1.2*prevPrice {
			break
		}
		prevPrice = price
	}
	// set is too small 1, 2 elements
	if i == 0 {
		i = 1
	}

	return auctions[0:i]
}

func secondCut(auctions Auctions) Auctions {
	auctionsMean := auctions.mean()
	stdDev := auctions.stdDev(auctionsMean)

	minPrice := auctionsMean - uint64(stdDev*throwFactor)
	maxPrice := auctionsMean + uint64(stdDev*throwFactor)

	minIndex := 0
	maxIndex := len(auctions)

	for i := 0; i < maxIndex; i++ {
		price := auctions[i].BuyoutUnit()
		if price < minPrice {
			minIndex = i + 1
		}
		if price > maxPrice {
			maxIndex = i
			break
		}
	}
	return auctions[minIndex:maxIndex]
}

// BuyoutAverage for a given item
func BuyoutAverage(auctions Auctions) (Market, error) {
	var market Market
	if len(auctions) == 0 {
		return market, ErrNoAuctions
	}
	market.Size = auctions.totalQuantity()

	auctions.sort()
	market.MinBuyout = auctions[0].Buyout
	auctions = firstCut(auctions)
	auctions = secondCut(auctions)

	market.AvgBuyout = auctions.mean()

	return market, nil
}
