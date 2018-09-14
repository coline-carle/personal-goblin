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
	auctions, err := Parse(data)
	if err != nil {
		t.Fatalf("unexpected eror parsing auction file: %s\n", err)
	}
	const expectedLen = 5355
	if len(auctions) != expectedLen {
		t.Errorf("expected %d auction got %d", expectedLen, len(auctions))
	}
}

func TestParseFilter(t *testing.T) {
	var err error
	var data []byte
	data, err = ioutil.ReadFile("fixtures/medivh.json")
	if err != nil {
		t.Fatalf("unexpected eror opening fixutre: %s\n", err)
	}

	whitelist := []uint64{23784, 161136}
	auctions, err := ParseFilter(data, whitelist)
	if err != nil {
		t.Fatalf("unexpected eror parsing auction file: %s\n", err)
	}
	if len(auctions) == 3 {
		t.Errorf("expected more than %d auctiosn got %d", 3, len(auctions))
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
