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
	expected := map[int64]Item{
		23784:  Item{ID: 23784, Name: "Adamantite Frame", Level: 33, Class: 0, Subclass: 1},
		161136: Item{ID: 161136, Name: "Azerite Forged Protection Plating", Level: 33, Class: 0, Subclass: 1},
	}
	assert.Equal(t, expected, items)
}
