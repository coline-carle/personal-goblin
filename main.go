package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/coline-carle/personal-goblin/auction"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	snapshot, err := auction.GetSnapshotURL(auction.EURegion, os.Getenv("BATTENET_API_KEY"), "twisting-nether")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	resp, err := http.Get(snapshot.URL)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	auctions, err := auction.Parse(body)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
