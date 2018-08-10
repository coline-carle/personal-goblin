package auction

import (
	"errors"
	"fmt"
)

var (
	// ErrNoAuctions error returned when the Auctions slice is empty
	ErrNoAuctions = errors.New("error: no auctions in slice")
)

const (
	throwFactor = 1.5
)

// BuyoutAverage for a given item
func BuyoutAverage(auctions Auctions) (float64, error) {
	if len(auctions) == 0 {
		return 0, ErrNoAuctions
	}
	auctions.sort()
	maxIndex := int(float64(len(auctions)) * 0.30)
	minIndex := int(float64(len(auctions)) * 0.15)
	prevPrice := float64(auctions[0].Buyout)
	var i int
	for i = 0; i < maxIndex; i++ {
		price := float64(auctions[i].Buyout)
		if i > minIndex && price > 1.2*prevPrice {
			fmt.Println(i)
			break
		}
		prevPrice = price
	}

	auctions = auctions[0:i]

	auctionsMean := mean(auctions)
	stdDev := stdDev(auctions, auctionsMean)

	minPrice := auctionsMean - stdDev*throwFactor
	maxPrice := auctionsMean + stdDev*throwFactor

	minIndex = 0
	maxIndex = len(auctions)

	for i = 0; i < maxIndex; i++ {
		price := float64(auctions[i].Buyout)
		if price < minPrice {
			minIndex = i + 1
		}
		if price > maxPrice {
			maxIndex = i
			break
		}
	}

	auctions = auctions[minIndex:maxIndex]

	return mean(auctions), nil
}
