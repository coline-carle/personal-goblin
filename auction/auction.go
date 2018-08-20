package auction

import (
	"log"

	"github.com/buger/jsonparser"
	"github.com/pkg/errors"
)

const (
	// ErrMalformedAuctionLine error
	ErrMalformedAuctionLine = "malformed auction line"
)

// Auction representation of an auction based on json auction house dump
type Auction struct {
	// Auc internal blizzard auction id
	Auc int64
	// Item item id
	Item int64
	// Buyout is the price set as direct buyout
	Buyout int64
	// Quantity item stack size
	Quantity int64

	// calculated buyout / unit
	buyoutUnit float64
}

type parser struct {
	auctions map[int64]Auctions
	filtered bool
	data     []byte
}

// Parse return a list of auction as defined in the auction house snaphsot
func Parse(data []byte) (map[int64]Auctions, error) {
	p := &parser{
		auctions: make(map[int64]Auctions),
		data:     data,
		filtered: false,
	}
	return p.Parse()
}

func (a *Auction) CalcBuyoutUnit() {
	a.buyoutUnit = float64(a.Buyout) / float64(a.Quantity)
}

// ParseFilter return a list of auction as defined in the auction house snaphsot filtered using a whitelist
func ParseFilter(data []byte, items map[int64]Item) (map[int64]Auctions, error) {
	auctions := make(map[int64]Auctions, len(items))
	for id := range items {
		auctions[id] = []Auction{}
	}
	p := &parser{
		auctions: auctions,
		data:     data,
		filtered: true,
	}
	return p.Parse()
}

func (p *parser) Parse() (map[int64]Auctions, error) {
	jsonparser.ArrayEach(p.data, func(line []byte, dataType jsonparser.ValueType, offset int, err error) {
		p.parseLine(line)
		if err != nil {
			log.Println(err)
		}

	}, "auctions")
	return p.auctions, nil

}

func getInt64(value []byte, name string) (int64, error) {
	var err error
	integerStr, _, _, err := jsonparser.Get(value, name)
	if err != nil {
		return 0, errors.Wrap(err, ErrMalformedAuctionLine)
	}
	return jsonparser.ParseInt(integerStr)
}

func (p *parser) validID(id int64) {

}

func (p *parser) parseLine(line []byte) error {
	var err error
	auction := Auction{}
	auction.Item, err = getInt64(line, "item")
	if err != nil {
		return errors.Wrap(err, "malformed auction")
	}
	if p.filtered {
		if _, ok := p.auctions[auction.Item]; !ok {
			return nil
		}
	}

	auction.Auc, err = getInt64(line, "auc")
	if err != nil {
		return errors.Wrap(err, "malformed auction")
	}

	auction.Buyout, err = getInt64(line, "buyout")
	if err != nil {
		return errors.Wrap(err, "malformed auction")
	}

	auction.Quantity, err = getInt64(line, "quantity")
	if err != nil {
		return errors.Wrap(err, "malformed auction")
	}

	p.auctions[auction.Item] = append(p.auctions[auction.Item], auction)
	return nil
}

// FetchLatestPrices return more recently updated price for a realm
// baseURL battle.net API entrypoint for a given region
// apiKey battle.net API key
// realmSlug slug of the realm
// lastModified the fetched timestamp should be superior to it for the fuction to fetch new pices
// func FetchLatestPrices(baseURL string, apiKey string, realmSlug string, lastModified time.Time) error {
// 	snapshot, err := GetSnapshotURL(baseURL, apiKey, realmSlug)
// 	if err != nil {
// 		return errors.Wrap(err, "can't fetch auction snapshot url")
// 	}
// 	return nil
// }
