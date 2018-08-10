package auction

import "testing"

func TestLoadWhitelist(t *testing.T) {
	items, err := LoadWhitelist()
	if err != nil {
		t.Fatalf("unexpected error loading whitelist: %s\n", err)
	}
	if _, ok := items[23784]; !ok {
		t.Error("expected to found item with 23784 id\n")
	}
}
