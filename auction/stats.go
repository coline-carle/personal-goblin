package auction

import "math"

func sums(auctions Auctions) (int64, int64) {
	total := int64(0)
	totalQuantity := int64(0)
	for _, auction := range auctions {
		total += auction.Buyout
		totalQuantity += auction.Quantity
	}
	return total, totalQuantity
}

func mean(auctions Auctions) float64 {
	total, totalQuantity := sums(auctions)
	return float64(total) / float64(totalQuantity)
}

func stdDev(auctions Auctions, mean float64) float64 {
	total := 0.0
	for _, auction := range auctions {
		total += math.Pow(float64(auction.Buyout*auction.Quantity)-mean, 2)
	}
	variance := total / float64(len(auctions)-1)
	return math.Sqrt(variance)
}
