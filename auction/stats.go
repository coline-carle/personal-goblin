package auction

import "math"

func (auctions Auctions) sum() uint64 {
	total := uint64(0)
	totalQuantity := uint64(0)
	for _, auction := range auctions {
		total += auction.Buyout
		totalQuantity += auction.Quantity
	}
	return total
}

func (auctions Auctions) totalQuantity() uint64 {
	totalQuantity := uint64(0)
	for _, auction := range auctions {
		totalQuantity += auction.Quantity
	}
	return totalQuantity
}

func (auctions Auctions) mean() uint64 {
	total := auctions.sum()
	totalQuantity := auctions.totalQuantity()
	return uint64(float64(total) / float64(totalQuantity))
}

func (auctions Auctions) stdDev(mean uint64) float64 {
	if len(auctions) == 1 {
		return 0.0
	}
	total := uint64(0)
	for _, auction := range auctions {
		total += (auction.BuyoutUnit() - mean) * (auction.BuyoutUnit() - mean) * auction.Quantity
	}
	variance := float64(total) / float64(len(auctions)-1)
	return math.Sqrt(variance)
}
