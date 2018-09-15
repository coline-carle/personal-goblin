package storage

import "database/sql"

// Storage is the SQL database interface
type Storage interface {
	Connect() error
	ApplyMigrations() error
	Tx(fn func(*sql.Tx) error) error

	GetWatchAuctionsItems(tx *sql.Tx) ([]uint64, error)
}
