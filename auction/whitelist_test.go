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
	expected := []int64{23784, 161136, 98717, 161131, 161137}
	assert.Equal(t, expected, items)
}
