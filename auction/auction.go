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
	Auc uint64
	// Item item id
	Item uint64
	// Buyout is the price set as direct buyout
	Buyout uint64
	// Quantity item stack size
	Quantity uint64
}

type parser struct {
	auctions map[uint64]Auctions
	filtered bool
	data     []byte
}

// Parse return a list of auction as defined in the auction house snaphsot
func Parse(data []byte) (map[uint64]Auctions, error) {
	p := &parser{
		auctions: make(map[uint64]Auctions),
		data:     data,
		filtered: false,
	}
	return p.Parse()
}

// BuyoutUnit is the price for a single item
func (a *Auction) BuyoutUnit() uint64 {
	return a.Buyout / a.Quantity
}

// ParseFilter return a list of auction as defined in the auction house snaphsot filtered using a whitelist
func ParseFilter(data []byte, items []uint64) (map[uint64]Auctions, error) {
	auctions := make(map[uint64]Auctions, len(items))
	for _, id := range items {
		auctions[id] = []Auction{}
	}
	p := &parser{
		auctions: auctions,
		data:     data,
		filtered: true,
	}
	return p.Parse()
}

func (p *parser) Parse() (map[uint64]Auctions, error) {
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
	var val int64
	auction := Auction{}
	val, err = getInt64(line, "item")
	if err != nil {
		return errors.Wrap(err, "malformed auction")
	}
	auction.Item = uint64(val)

	if p.filtered {
		if _, ok := p.auctions[auction.Item]; !ok {
			return nil
		}
	}

	val, err = getInt64(line, "auc")
	if err != nil {
		return errors.Wrap(err, "malformed auction")
	}
	auction.Auc = uint64(val)

	val, err = getInt64(line, "buyout")
	if err != nil {
		return errors.Wrap(err, "malformed auction")
	}
	auction.Buyout = uint64(val)

	val, err = getInt64(line, "quantity")
	if err != nil {
		return errors.Wrap(err, "malformed auction")
	}
	auction.Quantity = uint64(val)

	p.auctions[auction.Item] = append(p.auctions[auction.Item], auction)
	return nil
}
