package auction

import (
	"io/ioutil"
	"log"

	"github.com/buger/jsonparser"
	"github.com/pkg/errors"
)

const errJSONData = "error loading json data whitelist"
const errMalformedItem = "malformed item"
const errMalformedData = "malformed data"

func loadData(data []byte) []int64 {
	items := []int64{}
	jsonparser.ArrayEach(data, func(line []byte, dataType jsonparser.ValueType, offset int, err error) {
		item, err := parseItem(line)
		if err == nil {
			items = append(items, item)
		} else {
			log.Println(err)
		}
	})
	return items
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

// LoadLists load the lists of object id to track
func LoadLists(filenames []string) ([]int64, error) {
	items := []int64{}
	for _, filename := range filenames {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, errors.Wrap(err, errJSONData)
		}
		filter := loadData(data)
		items = append(items, filter...)

	}
	return items, nil
}
