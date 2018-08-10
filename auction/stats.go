package auction

import "math"

func sum(auctions Auctions) int64 {
	total := int64(0)
	for _, auction := range auctions {
		total += auction.Buyout * auction.Quantity
	}
	return total
}

func mean(auctions Auctions) float64 {
	return float64(sum(auctions)) / float64(len(auctions))
}

func stdDev(auctions Auctions, mean float64) float64 {
	total := 0.0
	for _, auction := range auctions {
		total += math.Pow(float64(auction.Buyout*auction.Quantity)-mean, 2)
	}
	variance := total / float64(len(auctions)-1)
	return math.Sqrt(variance)
}
