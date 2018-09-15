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

// BenchmarkParse bench without filtering
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

// BenchmarkParseFilter bench with filtering
func BenchmarkParseFilter(b *testing.B) {
	var err error
	var data []byte
	data, err = ioutil.ReadFile("fixtures/medivh.json")
	if err != nil {
		b.Fatalf("unexpected eror opening fixutre: %s\n", err)
	}

	whitelist := []uint64{
		152505, 152506, 152507, 152508, 152509, 152510, 152511,
		152632, 152637, 152634, 152636, 152494, 152495, 163082,
		152638, 152639, 152640, 152641, 162519,
		163222, 163223, 163224, 163225, 152559, 152560, 152561, 152557,
		162113, 152496, 152497, 152503, 152550,
		152578, 160325, 152580, 152581, 152582, 160322,
	}
	if err != nil {
		b.Fatalf("unexpected eror loading whitelist: %s\n", err)
	}

	for i := 0; i < b.N; i++ {
		auctions, _ := ParseFilter(data, whitelist)
		for range auctions {
		}
	}
}
