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
}

type parser struct {
	whitelist map[int64]struct{}
	auctions  []Auction
	data      []byte
}

// Parse return a list of auction as defined in the auction house snaphsot
func Parse(data []byte) ([]Auction, error) {
	p := &parser{
		auctions: []Auction{},
		data:     data,
	}
	return p.Parse()
}

// ParseFilter return a list of auction as defined in the auction house snaphsot filtered using a whitelist
func ParseFilter(data []byte, filter map[int64]struct{}) ([]Auction, error) {
	p := &parser{
		auctions:  []Auction{},
		data:      data,
		whitelist: filter,
	}
	return p.Parse()
}

func (p *parser) Parse() ([]Auction, error) {
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
	if len(p.whitelist) > 0 {
		if _, ok := p.whitelist[auction.Item]; !ok {
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

	p.auctions = append(p.auctions, auction)
	return nil
}
