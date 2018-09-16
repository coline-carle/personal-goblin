package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/coline-carle/personal-goblin/auction"
	"github.com/coline-carle/personal-goblin/storage/common"
	"github.com/joho/godotenv"
)

type itemLine struct {
	ID       int64
	Name     string
	Level    int
	Subclass int
	Buyout   int64
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func main() {
	godotenv.Load()
	storage, err := common.NewStorage("postgres://postgres:postgres@localhost/postgres?sslmode=disable")
	checkErr(err)

	fmt.Println("connected")

	err = storage.Connect()
	checkErr(err)

	err = storage.ApplyMigrations()
	checkErr(err)

	fmt.Println("connected and migrated")

	var ids []uint64
	err = storage.Tx(func(tx *sql.Tx) error {
		var err2 error
		ids, err2 = storage.GetWatchAuctionsItems(tx)
		return err2
	})
	checkErr(err)

	snapshot, err := auction.GetSnapshotURL(auction.EURegion, os.Getenv("BATTLENET_API_KEY"), "hyjal")
	checkErr(err)
	fmt.Printf("got snapshot url: %s\n", snapshot.URL)
	resp, err := http.Get(snapshot.URL)
	fmt.Printf("snaphsot status: %s\n", resp.Status)
	checkErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("download done")
	checkErr(err)
	auctions, err := auction.ParseFilter(body, ids)
	checkErr(err)
	for id, auc := range auctions {
		buyout, err := auction.BuyoutAverage(auc)
		if err != nil {
			log.Println(err)
		}

		fmt.Printf("id: %d - buyout: %.2f\n", id, float64(buyout.AvgBuyout)/100/100)
	}
}
