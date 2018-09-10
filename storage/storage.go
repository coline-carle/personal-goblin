package storage

import "database/sql"

type Storage interface {
	Connect() error
	ApplyMigrations() error
	Tx(fn func(*sql.Tx) error) error
}
