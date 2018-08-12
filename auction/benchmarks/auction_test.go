package test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/coline-carle/personal-goblin/auction"
)

func Benchmark_Parse(b *testing.B) {
	var err error
	var data []byte
	data, err = ioutil.ReadFile("../fixtures/medivh.json")
	if err != nil {
		b.Fatalf("unexpected eror opening fixutre: %s\n", err)
	}
	for i := 0; i < b.N; i++ {
		auctions, _ := auction.Parse(data)
		for range auctions {
		}
	}
}

func Benchmark_ParseFilter(b *testing.B) {
	var err error
	var data []byte
	data, err = ioutil.ReadFile("../fixtures/medivh.json")
	if err != nil {
		b.Fatalf("unexpected eror opening fixutre: %s\n", err)
	}

	files, err := filepath.Glob("./data/*.json")
	if err != nil {
		b.Fatalf("unexpected eror loading filters: %s\n", err)
	}

	whitelist, err := auction.LoadLists(files)
	if err != nil {
		b.Fatalf("unexpected eror loading whitelist: %s\n", err)
	}

	for i := 0; i < b.N; i++ {
		auctions, _ := auction.ParseFilter(data, whitelist)
		for range auctions {
		}
	}
}
