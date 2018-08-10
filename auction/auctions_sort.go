package auction

import "sort"

// Auctions auction type for sorting
type Auctions []Auction

func (a Auctions) Len() int {
	return len(a)
}

func (a Auctions) Less(i, j int) bool {
	return a[i].Buyout < a[j].Buyout
}

func (a Auctions) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Auctions) sort() {
	sort.Sort(a)
}
