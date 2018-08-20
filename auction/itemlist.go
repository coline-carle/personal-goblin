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

// Item struct representing a World of Warcraft item
type Item struct {
	ID       int64
	Name     string
	Level    int
	Class    int
	Subclass int
}

// LoadList load a single list
func LoadList(items map[int64]Item, data []byte) {
	jsonparser.ArrayEach(data, func(line []byte, dataType jsonparser.ValueType, offset int, err error) {
		item, err := parseItem(line)
		if err == nil {
			items[item.ID] = item
		} else {
			log.Println(err)
		}
	})
}

func parseItem(data []byte) (Item, error) {
	var err error
	item := Item{}
	item.ID, err = getInt64(data, "id")
	if err != nil {
		return item, errors.Wrap(err, "error reading item id")
	}

	item.Name, err = jsonparser.GetString(data, "name")
	if err != nil {
		return item, errors.Wrap(err, "error reading item name")
	}
	var value int64
	value, err = jsonparser.GetInt(data, "level")
	if err != nil {
		return item, errors.Wrap(err, "error reading item level")
	}
	item.Level = int(value)
	// value, err = jsonparser.GetInt(data, "class")
	// if err != nil {
	// 	return item, errors.Wrap(err, "error reading item class")
	// }
	// item.Class = int(value)
	value, err = jsonparser.GetInt(data, "subclass")
	if err != nil {
		return item, errors.Wrap(err, "error reading item subclass")
	}
	item.Subclass = int(value)
	return item, nil
}

// LoadLists load the lists of object id to track
func LoadLists(filenames []string) (map[int64]Item, error) {

	items := make(map[int64]Item)
	for _, filename := range filenames {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, errors.Wrap(err, errJSONData)
		}
		LoadList(items, data)

	}
	return items, nil
}
