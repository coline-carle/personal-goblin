package auction

import (
	"io/ioutil"
	"testing"
)

func TestParse(t *testing.T) {
	var err error
	var data []byte
	data, err = ioutil.ReadFile("fixtures/medivh.json")
	if err != nil {
		t.Fatalf("unexpected eror opening fixutre: %s\n", err)
	}
	var auctions []Auction
	auctions, err = Parse(data)
	if err != nil {
		t.Fatalf("unexpected eror parsing auction file: %s\n", err)
	}
	if len(auctions) < 10000 {
		t.Errorf("expected more than 10000 auction got %d", len(auctions))
	}
}

func TestParseFilter(t *testing.T) {
	var err error
	var data []byte
	data, err = ioutil.ReadFile("fixtures/medivh.json")
	if err != nil {
		t.Fatalf("unexpected eror opening fixutre: %s\n", err)
	}

	whitelist, err := LoadWhitelist()
	if err != nil {
		t.Fatalf("unexpected eror loading whitelist: %s\n", err)
	}
	var auctions []Auction
	auctions, err = ParseFilter(data, whitelist)
	if err != nil {
		t.Fatalf("unexpected eror parsing auction file: %s\n", err)
	}
	if len(auctions) > 10000 || len(auctions) < 5000 {
		t.Errorf("expected more than 10000 auction got %d", len(auctions))
	}
}

func BenchmarkParse(b *testing.B) {
	var err error
	var data []byte
	data, err = ioutil.ReadFile("fixtures/medivh.json")
	if err != nil {
		b.Fatalf("unexpected eror opening fixutre: %s\n", err)
	}
	for i := 0; i < b.N; i++ {
		auctions, _ := Parse(data)
		for range auctions {
		}
	}
}
