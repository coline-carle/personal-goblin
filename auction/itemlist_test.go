package auction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadLists(t *testing.T) {
	items, err := LoadLists([]string{"./fixtures/filter.json"})
	if err != nil {
		t.Fatalf("unexpected error loading whitelist: %s\n", err)
	}
	expected := map[int64]Item{23784: Item{ID: 23784, Name: "Adamantite Frame", Level: 33, Class: 7, Subclass: 1}}
	assert.Equal(t, expected, items)
}
