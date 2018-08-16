package auction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type jsonTime time.Time

// Snapshot is the descriptiàon of the snapshot returned by the battle.net
// API
type Snapshot struct {
	URL          string   `url:"url"`
	LastModified jsonTime `json:"lastModified"`
}

// APIResponse is the root of the response for the battle net api fetch
type APIResponse struct {
	Files []Snapshot `json:"files"`
}

// Region of the battle.net api Query (EuRegion, UsRegion, TwRegion, KrRegion)
type Region int

const apiPath = "wow/auction/data"

const (
	// EURegion Europe Region
	EURegion = "https://eu.api.battle.net"
	// USRegion U.S.A. Region
	USRegion = "https://us.api.battle.net"
	// KRRegion Korean Region
	KRRegion = "https://kr.api.battle.net"
	// TWRegion Taiwan Region
	TWRegion = "https://tw.api.battle.net"
)

func (t *jsonTime) UnmarshalJSON(s []byte) (err error) {
	r := strings.Replace(string(s), `"`, ``, -1)

	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(q/1000, 0).UTC()
	return
}

func (t jsonTime) String() string { return time.Time(t).String() }

// GetSnapshotURL fetch the auction snapshot url from battle net API
// apiKey you battle net API Key
// BaseURL the base url of the  Region
// realmSlug slug of the realm
func GetSnapshotURL(baseURL string, apiKey string, realmSlug string) (Snapshot, error) {
	snapshotURL, err := url.Parse(baseURL)
	if err != nil {
		return Snapshot{}, errors.Wrap(err, "can't parse wow API url")
	}
	snapshotURL.Path = path.Join(apiPath, realmSlug)
	q := snapshotURL.Query()
	q.Set("apikey", apiKey)
	snapshotURL.RawQuery = q.Encode()
	resp, err := http.Get(snapshotURL.String())
	if err != nil {
		return Snapshot{}, errors.Wrap(err, "error fetching api url")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Snapshot{}, errors.Wrap(err, "error decoding response body")
	}
	result := APIResponse{}
	json.Unmarshal(body, &result)

	// the response is supposed to return only one result
	if len(result.Files) != 1 {
		return Snapshot{}, fmt.Errorf("invalid api response. Unexpected len(result.Files) %d", len(result.Files))
	}
	return result.Files[0], nil
}
