package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/coline-carle/personal-goblin/auction"
)

func main() {
	data, err := ioutil.ReadFile("fixtures/medivh.json")
	auction.Parse(data)
	auction, err := auction.GetSnapshotURL(auction.EURegion, "***REMOVED***", "archimonde")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Printf("%+v", auction)
}
