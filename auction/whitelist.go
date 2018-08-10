package auction

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/buger/jsonparser"
	"github.com/pkg/errors"
)

const errJSONData = "error loading json data whitelist"
const errMalformedItem = "malformed item"
const errMalformedData = "malformed data"

func loadData(items map[int64]struct{}, data []byte) {

	jsonparser.ArrayEach(data, func(line []byte, dataType jsonparser.ValueType, offset int, err error) {
		id, err := parseItem(line)
		if err == nil {
			items[id] = struct{}{}
		} else {
			log.Println(err)
		}
	})
}

func parseItem(item []byte) (int64, error) {
	var err error
	integerStr, _, _, err := jsonparser.Get(item, "id")
	if err != nil {
		return 0, errors.Wrap(err, errMalformedItem)
	}
	integer, err := jsonparser.ParseInt(integerStr)
	if err != nil {
		return 0, errors.Wrap(err, errMalformedItem)
	}
	return integer, nil
}

// LoadWhitelist load the lists of the object we want to track !
func LoadWhitelist() (map[int64]struct{}, error) {
	files, err := filepath.Glob("./data/*.json")
	if err != nil {
		return nil, errors.Wrap(err, errJSONData)
	}

	items := make(map[int64]struct{})
	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, errors.Wrap(err, errJSONData)
		}
		loadData(items, data)

	}
	return items, nil
}
