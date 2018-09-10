package common

import (
	"github.com/coline-carle/personal-goblin/storage"
	"github.com/coline-carle/personal-goblin/storage/postgres"
)

// NewStorage Create a new storage based on the URL
// FIXME: only work with postgres database
func NewStorage(url string) (storage.Storage, error) {
	return postgres.New(url)
}
