package auction

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestGetAuctionHouseSnapshotURL(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if r.URL.EscapedPath() != "/wow/auction/data/server" {
			t.Errorf("Expected request to ‘/wow/auction/data/server’, got ‘%s’", r.URL.EscapedPath())
		}
		//Check if file exists and open
		file, err := os.Open("./fixtures/api_auction_call.json")
		defer file.Close()
		if err != nil {
			t.Errorf("unexpected error loading fixture '%s'", err)
		}
		w.Header().Set("Content-Type", "application/json")
		io.Copy(w, file)

	}))
	defer ts.Close()

	const expectedURL = "http://auction-api-eu.worldofwarcraft.com/auction-data/random-number/auctions.json"

	var expectedLastMod jsonTime
	parsedLastMod, err := time.Parse(time.RFC1123, "Fri, 10 Aug 2018 18:42:37 UTC")
	expectedLastMod = jsonTime(parsedLastMod)
	if err != nil {
		t.Errorf("Unexpected error parsing expectedLastMod date '%s'", err)
	}

	snapshot, err := GetSnapshotURL(ts.URL, "apikey", "server")
	if err != nil {
		t.Fatalf("unexpected error GetAuctionHouseSnapshotURL() %s", err)
	}

	if snapshot.URL != expectedURL {
		t.Errorf("Unexpected snapshot URL\n* expected: '%s'\n* got: %s", expectedURL, snapshot.URL)
	}

	if expectedLastMod != snapshot.LastModified {
		t.Errorf("unexpedted LastModified\n* expected: '%s'\n* got: %s", expectedLastMod, snapshot.LastModified)
	}
}
